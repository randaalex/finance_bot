package emailprocessor

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
)

func (p *Processor) DownloadEmails(ctx context.Context) {
	log.Println("Connecting to server...")

	imapClient, err := client.DialTLS(p.settings.Address, nil)
	if err != nil {
		panic(err)
	}

	defer imapClient.Logout()

	if err := imapClient.Login(p.settings.Username, p.settings.Password); err != nil {
		panic(err)
	}

	_, err = imapClient.Select(mailboxName, false)
	if err != nil {
		panic(err)
	}

	searchCriteria := imap.NewSearchCriteria()
	searchCriteria.WithoutFlags = []string{imap.SeenFlag}
	//searchCriteria.WithoutFlags = []string{}
	messageUids, err := imapClient.Search(searchCriteria)
	if err != nil {
		panic(err)
	}

	for _, uuid := range messageUids {
		seqSet := new(imap.SeqSet)
		seqSet.AddNum(uuid)

		var section imap.BodySectionName
		items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope, imap.FetchFlags}

		messages := make(chan *imap.Message, 0)
		done := make(chan error, 0)
		go func() {
			done <- imapClient.Fetch(seqSet, items, messages)
		}()

		for msg := range messages {
			if email := p.parseImapMessage(&section, msg); email != nil {
				//p.emails = append(p.emails, *email)
				p.ProcessEmail(ctx, *email)
			}
		}

		time.Sleep(5*time.Second)
	}


	//if len(messageUids) > 0 {
	//	seqSet := new(imap.SeqSet)
	//	seqSet.AddNum(messageUids...)
	//	//seqSet.AddNum(messageUids[0])
	//
	//	var section imap.BodySectionName
	//	items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope, imap.FetchFlags}
	//
	//	messages := make(chan *imap.Message, 0)
	//	done := make(chan error, 0)
	//	go func() {
	//		done <- imapClient.Fetch(seqSet, items, messages)
	//	}()
	//
	//	for msg := range messages {
	//		if email := p.parseImapMessage(&section, msg); email != nil {
	//			//p.emails = append(p.emails, *email)
	//			p.ProcessEmail(ctx, *email)
	//		}
	//	}
	//}
}

func (p *Processor) parseImapMessage(section *imap.BodySectionName, msg *imap.Message) *parsedImapMessage {
	if msg == nil {
		log.Println("Server didn't returned message")
	}

	r := msg.GetBody(section)
	if r == nil {
		log.Println("Server didn't returned message body")
	}

	mr, err := mail.CreateReader(r)
	if err != nil {
		panic(err)
	}

	for {
		mailPart, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		switch mailPart.Header.(type) {
		case *mail.InlineHeader:
			body, _ := ioutil.ReadAll(mailPart.Body)

			return &parsedImapMessage{msg.Envelope.Subject, string(body)}
		}
	}

	return nil
}
