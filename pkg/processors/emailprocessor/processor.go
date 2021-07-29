package emailprocessor

import (
	"context"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/entities"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/firefly"
	"github.com/randaalex/finance_bot/pkg/parsers/alfamail"
	"github.com/randaalex/finance_bot/pkg/telegram"
)

const (
	mailboxName        = "ALFABANK"
	emailSubjectFormat = `^\d\.\d{4}`
)

type Storage interface {
	CreateMapping(ctx context.Context, arg db.CreateMappingParams) (db.Mapping, error)
	GetMappingByTransactionDetails(ctx context.Context, transactionDetails string) (db.Mapping, error)
}

type Processor struct {
	Storage       Storage
	FireflyClient *firefly.APIClient
	TelegramBot   *telegram.Bot
	accounts      *map[int]string
	categories    *map[int]string
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
	accounts *map[int]string,
	categories *map[int]string,
	settings *Settings,
) *Processor {
	return &Processor{
		Storage:       storage,
		FireflyClient: fireflyClient,
		TelegramBot:   bot,
		accounts:      accounts,
		categories:    categories,
		settings:      settings,
	}
}

func (p *Processor) Process() {
	ctx := context.TODO()

	log.Printf("process")

	emailParser := alfamail.NewAlfaParser(nil)

	for _, email := range p.emails {
		if isTransactionEmail, _ := regexp.Match(emailSubjectFormat, []byte(email.subject)); !isTransactionEmail {
			log.Printf("Skip email with invalid subject: %s", email.subject)
			continue
		}

		parsedTransactions := emailParser.Parse(email.body)
		for _, transaction := range *parsedTransactions {
			selectedCategoryID := 0

			existedTransaction, e := p.Storage.GetMappingByTransactionDetails(ctx, transaction.Details)
			if e != nil {
				log.Printf("mapping not found: %v\n", transaction.Details)
			} else {
				log.Printf("mapping found: %v\n", existedTransaction)
				selectedCategoryID = int(existedTransaction.CategoryID)
			}

			sourceId := int32(7) // TODO: get sourceId from mapping
			categoryId := int32(selectedCategoryID)
			transactionHash := transaction.Hash()

			fmt.Printf("transaction: %v+\n", transaction)
			fmt.Printf("transaction Time: %v+\n", transaction.Time)

			split := firefly.TransactionSplit{
				Type:              "withdrawal",
				Date:              transaction.Time.Format("2006-01-02"),
				Amount:            fmt.Sprintf("%.2f", transaction.Amount),
				Description:       transaction.Details,
				SourceId:          *firefly.NewNullableInt32(&sourceId),
				Notes:             *firefly.NewNullableString(&transaction.Details),
				InternalReference: *firefly.NewNullableString(&transactionHash),
				Tags:              []string{},
				ImportHashV2:      *firefly.NewNullableString(&transactionHash),
			}

			if selectedCategoryID != 0 {
				split.CategoryId = *firefly.NewNullableInt32(&categoryId)
			}

			req := firefly.Transaction{
				Transactions: []firefly.TransactionSplit{
					split,
				},
			}

			fmt.Fprintf(os.Stderr, "Req: %v\n", req)

			// TODO: add request for update VirtualBalance of account in firefly
			rawFireflyTransaction, r, err := p.FireflyClient.TransactionsApi.StoreTransaction(context.TODO()).Transaction(req).Execute()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.StoreTransaction``: %v\n", err)
				fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
			}
			// response from `GetConfiguration`: Configuration
			fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.StoreTransaction`: %v\n", rawFireflyTransaction)
			if err != nil {
				fmt.Printf("r: %v\n", r)

				panic(err)
			}

			data := rawFireflyTransaction.GetData()
			attrs := data.GetAttributes()
			subTransactions := attrs.GetTransactions()
			subTransaction := subTransactions[0]

			id, _ := strconv.Atoi(data.Id)
			amount, _ := strconv.ParseFloat(subTransaction.GetAmount(), 32)
			foreignAmount, _ := strconv.ParseFloat(subTransaction.GetForeignAmount(), 32)

			fireflyTransaction := &entities.FireflyTransaction{
				Id:                  id,
				AccountId:           int(subTransaction.GetSourceId()),
				AccountName:         subTransaction.GetSourceName(),
				AccountCurrencyCode: subTransaction.GetCurrencyCode(),
				Time:                transaction.Time,
				Type:                subTransaction.GetType(),
				Amount:              float32(amount),
				ForeignAmount:       float32(foreignAmount),
				ForeignCurrencyCode: subTransaction.GetForeignCurrencyCode(),
				CategoryId:          int(subTransaction.GetCategoryId()),
				CategoryName:        subTransaction.GetCategoryName(),
				Hash:                subTransaction.GetInternalReference(),
				RawContent:          subTransaction.GetNotes(),
				Tags:                subTransaction.GetTags(),
				Description:         subTransaction.GetDescription(),
			}

			fmt.Printf("%v+\n", fireflyTransaction)

			p.TelegramBot.RenderTransaction(fireflyTransaction)
		}
	}
}
