package emailprocessor

import (
	"context"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
	"github.com/randaalex/finance_bot/pkg/parsers/alfamail"
	"github.com/randaalex/finance_bot/pkg/telegram"
	"log"
	"os"
	"regexp"
	"time"
)

const (
	mailboxName        = "ALFABANK"
	emailSubjectFormat = `^\d\.\d{4}`
)

type Storage interface {
	CreateTransactionsLog(ctx context.Context, arg db.CreateTransactionsLogParams) (db.TransactionsLog, error)
	GetTransactionsLogByDescription(ctx context.Context, description string) (db.TransactionsLog, error)
}

type Processor struct {
	Storage       Storage
	FireflyClient *firefly.APIClient
	TelegramBot   *telegram.Bot
	EmailParser   *alfamail.AlfaParser
	accounts      *[]entities.Account
	categories    *[]entities.Category
	settings      *Settings
	emails        []parsedImapMessage
}

type Settings struct {
	Address  string
	Username string
	Password string
}

type parsedImapMessage struct {
	subject string
	body    string
}

func NewProcessor(
	storage Storage,
	fireflyClient *firefly.APIClient,
	bot *telegram.Bot,
	emailParser *alfamail.AlfaParser,
	accounts *[]entities.Account,
	categories *[]entities.Category,
	settings *Settings,
) *Processor {
	return &Processor{
		Storage:       storage,
		FireflyClient: fireflyClient,
		TelegramBot:   bot,
		EmailParser:   emailParser,
		accounts:      accounts,
		categories:    categories,
		settings:      settings,
	}
}

func (p *Processor) Process(ctx context.Context) {
	log.Printf("process")

	for _, email := range p.emails {
		if isValidEmail, _ := regexp.Match(emailSubjectFormat, []byte(email.subject)); !isValidEmail {
			log.Printf("Skip email with invalid subject: %s", email.subject)
			continue
		}

		transactionReq := *p.EmailParser.Parse(email.body)

		p.processTransaction(ctx, transactionReq)

		time.Sleep(5 * time.Second)
	}
}

func (p *Processor) ProcessEmail(ctx context.Context, email parsedImapMessage) {
	log.Printf("process signle email")

	if isValidEmail, _ := regexp.Match(emailSubjectFormat, []byte(email.subject)); !isValidEmail {
		log.Printf("Skip email with invalid subject: %s", email.subject)
		return
	}

	transactionReq := *p.EmailParser.Parse(email.body)

	p.processTransaction(ctx, transactionReq)
}

func (p *Processor) processTransaction(ctx context.Context, transactionReq entities.Transaction) {
	recentTransaction, e := p.Storage.GetTransactionsLogByDescription(ctx, transactionReq.Description)
	if e != nil {
		log.Printf("mapping not found: %v\n", transactionReq.Description)
	} else {
		log.Printf("mapping found: %v\n", recentTransaction)
		transactionReq.CategoryId = recentTransaction.CategoryID
	}

	fireflyTransactionReq := entities.ConvertTransactionToFireflyTransaction(&transactionReq)

	// TODO: add request for update VirtualBalance of account in firefly
	fireflyTransaction, r, err :=
		p.FireflyClient.TransactionsApi.StoreTransaction(context.TODO()).Transaction(*fireflyTransactionReq).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.StoreTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)

		panic(err)
	}

	transaction := entities.ConvertFireflyTransactionToTransaction(&fireflyTransaction)

	p.TelegramBot.RenderTransaction(transaction)
}
