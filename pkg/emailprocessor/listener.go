package emailprocessor

import (
	"context"
	"errors"
	"fmt"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"io"
	"io/ioutil"
	"time"
)

func (p *Processor) Connect(ctx context.Context) error {
	imapClient, err := client.DialTLS(p.settings.MailClientAddress, nil)
	if err != nil {
		return ImapError{
			Err: fmt.Errorf("dial error: %w", err),
		}
	}

	p.imapClient = imapClient

	defer imapClient.Logout()

	if err = p.imapClient.Login(p.settings.MailClientUsername, p.settings.MailClientPassword); err != nil {
		return ImapError{
			Err: fmt.Errorf("login error: %w", err),
		}
	}

	return nil
}

func (p *Processor) Start(ctx context.Context) error {
	imapClient, err := client.DialTLS(p.settings.MailClientAddress, nil)
	if err != nil {
		return ImapError{
			Err: fmt.Errorf("dial error: %w", err),
		}
	}

	p.imapClient = imapClient

	defer imapClient.Logout()

	if err = p.imapClient.Login(p.settings.MailClientUsername, p.settings.MailClientPassword); err != nil {
		return ImapError{
			Err: fmt.Errorf("login error: %w", err),
		}
	}

	_, err = p.imapClient.Select(mailboxName, false)
	if err != nil {
		return fmt.Errorf("select mailbox error: %w", err)
	}

	searchCriteria := imap.NewSearchCriteria()
	searchCriteria.WithoutFlags = []string{imap.SeenFlag}
	//searchCriteria.WithoutFlags = []string{}

	for {
		if err = p.fetchMessages(ctx, searchCriteria); err != nil {
			panic(err)
		}

		time.Sleep(10 * time.Second)
	}

	return nil
}

func (p *Processor) fetchMessages(ctx context.Context, searchCriteria *imap.SearchCriteria) error {
	messageUids, err := p.imapClient.Search(searchCriteria)
	if err != nil {
		return fmt.Errorf("search error: %w", err)
	}

	for _, uuid := range messageUids {
		seqSet := new(imap.SeqSet)
		seqSet.AddNum(uuid)

		var section imap.BodySectionName
		items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope, imap.FetchFlags}

		messages := make(chan *imap.Message, 0)
		done := make(chan error, 0)
		go func() {
			done <- p.imapClient.Fetch(seqSet, items, messages)
		}()

		for msg := range messages {
			email, err := p.parseImapMessage(&section, msg)
			if err != nil {
				p.copyMailToErrorFolder(seqSet)
				// TODO: logger, sentry
				return err
			}

			err = p.ProcessEmail(ctx, email)
			if err != nil {
				p.copyMailToErrorFolder(seqSet)
				// TODO: logger, sentry
				return err
			}
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}

func (p *Processor) copyMailToErrorFolder(seqSet *imap.SeqSet) {
	err := p.imapClient.Copy(seqSet, "ALFABANK_ERROR")
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
