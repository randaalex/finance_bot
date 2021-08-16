package emailprocessor

import (
	"context"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/emailprocessor/alfaparser"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
	"github.com/randaalex/finance_bot/pkg/telegram"
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
	Storage         Storage
	FireflyClient   *firefly.APIClient
	TelegramBot     *telegram.Bot
	EmailParser     *alfaparser.Parser
	logger          *logrus.Logger
	accounts        *[]entities.Account
	categories      *[]entities.Category
	settings        *entities.Settings
	//imapConnection2 *imapclient.Client
}

type parsedImapMessage struct {
	subject string
	body    string
}

func NewProcessor(
	storage Storage,
	fireflyClient *firefly.APIClient,
	bot *telegram.Bot,
	emailParser *alfaparser.Parser,
	logger *logrus.Logger,
	accounts *[]entities.Account,
	categories *[]entities.Category,
	settings *entities.Settings,
) *Processor {
	return &Processor{
		Storage:       storage,
		FireflyClient: fireflyClient,
		TelegramBot:   bot,
		EmailParser:   emailParser,
		logger:        logger,
		accounts:      accounts,
		categories:    categories,
		settings:      settings,
	}
}

func (p *Processor) ProcessEmail(ctx context.Context, email *parsedImapMessage) error {
	if isValidEmail, _ := regexp.Match(emailSubjectFormat, []byte(email.subject)); !isValidEmail {
		p.logger.Infof("Skip email with invalid subject: %s", email.subject)
		return nil
	}

	transactionReq, err := p.EmailParser.Parse(email.body)
	if err != nil {
		p.logger.Infof("parse error: %+v\n", err)

		sentry.ConfigureScope(func(scope *sentry.Scope) {
			scope.SetExtras(map[string]interface{}{
				"mail":   err.(alfaparser.ParseMailError).Mail,
				"line":   err.(alfaparser.ParseMailError).Line,
				"action": err.(alfaparser.ParseMailError).Action,
			})
			sentry.CaptureException(err)
		})

		return nil
	}

	return p.processTransaction(ctx, transactionReq)
}

func (p *Processor) processTransaction(ctx context.Context, transactionReq *entities.Transaction) error {
	recentTransaction, e := p.Storage.GetTransactionsLogByDescription(ctx, transactionReq.Description)
	if e != nil {
		p.logger.Infof("mapping not found: %v", transactionReq.Description)
	} else {
		p.logger.Infof("mapping found: %v\n", recentTransaction)
		transactionReq.CategoryId = int(recentTransaction.CategoryID)
	}

	fireflyTransactionStoreReq := entities.ConvertTransactionToFireflyTransactionStore(transactionReq)

	// TODO: add request for update VirtualBalance of account in firefly
	fireflyTransaction, r, err :=
		p.FireflyClient.TransactionsApi.StoreTransaction(context.TODO()).TransactionStore(*fireflyTransactionStoreReq).Execute()

	if err != nil {
		var rr []byte
		if r != nil {
			rr, _ = ioutil.ReadAll(r.Body)
		}

		p.logger.WithFields(logrus.Fields{
			"err": err, "httpResp": string(rr),
		}).Info("Error when calling `TransactionsApi.StoreTransaction`")

		if fireflyError, ok := err.(*firefly.GenericOpenAPIError); ok {
			if strings.Contains(string(fireflyError.Body()), "Duplicate of transaction") {
				p.logger.Infof("Transaction duplicated in firefly: %v\n", fireflyTransaction)

				return nil
			}
		}

		panic(err)
	}

	transaction := entities.ConvertFireflyTransactionToTransaction(&fireflyTransaction)

	p.TelegramBot.RenderTransaction(transaction)

	return nil
}
