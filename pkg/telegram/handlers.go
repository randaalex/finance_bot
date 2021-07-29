package telegram

import (
	"context"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
	"gopkg.in/tucnak/telebot.v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func (b *Bot) updateTransactionCategoryHandler(c *telebot.Callback) {
	//ctx := context.TODO()

	fmt.Println(c.ID)
	fmt.Printf("C: %+v\n",c)

	res := strings.Split(c.Data, "|")
	fmt.Printf("TE%+v\n",res)

	fmt.Println("updateTransactionCategory btn pressed")


	transactionId, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}

	categoryId, err := strconv.Atoi(res[1])
	if err != nil {
		panic(err)
	}

	b.RenderTransactionCategoriesKeyboard(transactionId, categoryId)

	err = b.Telebot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		panic(err)
	}
}

func (b *Bot) setTransactionCategoryHandler(c *telebot.Callback) {
	ctx := context.TODO()

	fmt.Println(c.ID)
	fmt.Printf("C: %+v\n",c)

	res := strings.Split(c.Data, "|")
	fmt.Printf("TE%+v\n",res)

	fmt.Println("setTransactionCategory btn pressed")
	err := b.Telebot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		panic(err)
	}


	transactionId, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err)
	}

	categoryId, err := strconv.Atoi(res[1])
	if err != nil {
		panic(err)
	}

	rawFireflyTransaction1, r, err := b.FireflyClient.TransactionsApi.GetTransaction(ctx, int32(transactionId)).Execute()
	if err != nil {
		panic(err)
	}

	data1 := rawFireflyTransaction1.GetData()
	attrs1 := data1.GetAttributes()
	subTransactions1 := attrs1.GetTransactions()
	subTransaction1 := subTransactions1[0]
	subTransaction1.SetCategoryId(int32(categoryId))
	subTransaction1.SetBillIdNil()

	if subTransaction1.GetForeignCurrencyId() == 0 {
		subTransaction1.SetForeignCurrencyIdNil()
	}

	attrs1.SetTransactions([]firefly.TransactionSplit{subTransaction1})

	rawFireflyTransaction, r, err := b.FireflyClient.TransactionsApi.UpdateTransaction(context.TODO(), int32(transactionId)).Transaction(attrs1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.UpdateTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: Configuration
	fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.UpdateTransaction`: %v\n", rawFireflyTransaction)
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
	ttime, _ := time.Parse("2006-01-02", subTransaction.GetDate())

	fireflyTransaction := &entities.FireflyTransaction{
		Id:                  id,
		AccountId:           int(subTransaction.GetSourceId()),
		AccountName:         subTransaction.GetSourceName(),
		AccountCurrencyCode: subTransaction.GetCurrencyCode(),
		Time:                ttime,
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

	b.RenderTransaction(fireflyTransaction)
}

func (b *Bot) deleteTransactionHandler(c *telebot.Callback) {
	fmt.Println(c)
	fmt.Println("deleteTransaction btn pressed")
	err := b.Telebot.Respond(c, &telebot.CallbackResponse{Text: "not implemented yet"})
	if err != nil {
		panic(err)
	}
}

