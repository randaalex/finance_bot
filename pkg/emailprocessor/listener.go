package emailprocessor

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"github.com/getsentry/sentry-go"
)

const ImapUpdatesInterval = time.Minute

func (p *Processor) Start(ctx context.Context) {
	defer sentry.Recover()

	for {
		p.logger.Info("Connecting...")
		imapConnection, errConn := p.connect(ctx)
		defer imapConnection.Logout()

		if errConn != nil {
			p.logger.WithFields(logrus.Fields{"err": errConn}).Error("mail connection error")
			sentry.CaptureException(errConn)

			p.logger.Info("Sleeping...")
			time.Sleep(ImapUpdatesInterval)
			break
		}

		p.logger.Info("Processing new mails...")
		errProcess := p.process(ctx, imapConnection)
		if errProcess != nil {
			p.logger.WithFields(logrus.Fields{"err": errProcess}).Error("mail processing error")
			sentry.CaptureException(errProcess)
		}

		p.logger.Info("Sleeping...")
		time.Sleep(ImapUpdatesInterval)
	}
}

func (p *Processor) connect(ctx context.Context) (*client.Client, error) {
	imapConnection, err := client.DialTLS(p.settings.MailClientAddress, nil)
	if err != nil {
		return nil, ImapError{
			Err: fmt.Errorf("dial error: %w", err),
		}
	}

	if err = imapConnection.Login(p.settings.MailClientUsername, p.settings.MailClientPassword); err != nil {
		return nil, ImapError{
			Err: fmt.Errorf("login error: %w", err),
		}
	}

	_, err = imapConnection.Select(mailboxName, false)
	if err != nil {
		return nil, ImapError{
			Err: fmt.Errorf("select mailbox error: %w", err),
		}
	}

	return imapConnection, nil
}

func (p *Processor) process(ctx context.Context, imapConnection *client.Client) error {
	searchCriteria := imap.NewSearchCriteria()
	searchCriteria.WithoutFlags = []string{imap.SeenFlag}
	//searchCriteria.WithoutFlags = []string{}

	messageUids, err := imapConnection.Search(searchCriteria)
	if err != nil {
		return fmt.Errorf("imap search error: %w", err)
	}

	for _, uuid := range messageUids {
		seqSet := new(imap.SeqSet)
		seqSet.AddNum(uuid)

		var section imap.BodySectionName
		items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope, imap.FetchFlags}

		messages := make(chan *imap.Message, 0)
		done := make(chan error, 0)
		go func() {
			done <- imapConnection.Fetch(seqSet, items, messages)
		}()

		for msg := range messages {
			email, err := p.parseImapMessage(&section, msg)
			if err != nil {
				p.copyMailToErrorFolder(imapConnection, seqSet)

				p.logger.WithFields(logrus.Fields{"msg": msg, "err": err}).Error("parse imap message error")

				return fmt.Errorf("imap parse message error: %w", err)
			}

			err = p.ProcessEmail(ctx, email)
			if err != nil {
				p.copyMailToErrorFolder(imapConnection, seqSet)

				p.logger.WithFields(logrus.Fields{"msg": msg, "err": err}).Error("mail process error")

				return fmt.Errorf("mail process message error: %w", err)
			}
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func (p *Processor) copyMailToErrorFolder(imapConnection *client.Client, seqSet *imap.SeqSet) {
	err := imapConnection.Copy(seqSet, "ALFABANK_ERROR")
	if err != nil {
		panic(err) // TODO: fix
	}
}

func (p *Processor) parseImapMessage(section *imap.BodySectionName, msg *imap.Message) (*parsedImapMessage, error) {
	if msg == nil {
		return nil, errors.New("server didn't returned message body")
	}

	r := msg.GetBody(section)
	if r == nil {
		return nil, errors.New("server didn't returned message body")
	}

	mr, err := mail.CreateReader(r)
	if err != nil {
		return nil, fmt.Errorf("mail read error: %w", err)
	}

	for {
		mailPart, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("mail part read error: %w", err)
		}

		switch mailPart.Header.(type) {
		case *mail.InlineHeader:
			body, _ := ioutil.ReadAll(mailPart.Body)

			return &parsedImapMessage{msg.Envelope.Subject, string(body)}, nil
		}
	}

	return nil, errors.New("empty reader error")
}
