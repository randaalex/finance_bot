package cli

import (
	"io"
	"io/ioutil"
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-message/mail"
	"github.com/spf13/cobra"
)

const (
	mailboxName = "ALFABANK"
)

func MailHandler(cmd *cobra.Command, args []string) {
	log.Println("Connecting to server...")

	var settings = getSettings()

	c, err := client.DialTLS(settings.MailClientAddress, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer c.Logout()

	if err := c.Login(settings.MailClientUsername, settings.MailClientPassword); err != nil {
		log.Fatal(err)
	}

	_, err = c.Select(mailboxName, false)
	if err != nil {
		log.Fatal(err)
	}

	criteria := imap.NewSearchCriteria()
	criteria.WithoutFlags = []string{imap.SeenFlag}
	ids, err := c.Search(criteria)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("IDs found:", ids)

	if len(ids) > 0 {
		seqSet := new(imap.SeqSet)
		seqSet.AddNum(ids...)

		var section imap.BodySectionName
		items := []imap.FetchItem{section.FetchItem(), imap.FetchEnvelope, imap.FetchFlags}

		messages := make(chan *imap.Message, 0)
		done := make(chan error, 0)
		go func() {
			done <- c.Fetch(seqSet, items, messages)
		}()

		for msg := range messages {
			if msg == nil {
				log.Fatal("Server didn't returned message")
			}

			r := msg.GetBody(&section)
			if r == nil {
				log.Fatal("Server didn't returned message body")
			}

			// Create a new mail reader
			mr, err := mail.CreateReader(r)
			if err != nil {
				log.Fatal(err)
			}

			for {
				mailPart, err := mr.NextPart()
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatal(err)
				}

				switch mailPart.Header.(type) {
				case *mail.InlineHeader:
					log.Printf("Subject: %v\n", msg.Envelope.Subject)

					body, _ := ioutil.ReadAll(mailPart.Body)
					log.Printf("Body: %v\n", string(body))

				}
			}
		}
	}

	log.Println("Done!")
}

