package telegram

import (
	"context"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"gopkg.in/tucnak/telebot.v2"
	"os"
	"strconv"
	"strings"
)

func (b *Bot) updateTransactionCategoryHandler(c *telebot.Callback) {
	ctx := context.TODO()

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

	b.UpdateTransactionWithCategoriesKeyboard(ctx, c.Message, transactionId, categoryId)

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

	transaction := entities.ConvertFireflyTransactionToTransaction(&rawFireflyTransaction1)
	transaction.CategoryId = int32(categoryId)

	_, err = b.Storage.CreateTransactionsLog(ctx, db.CreateTransactionsLogParams{
		Description: transaction.Description,
		CategoryID:  int32(categoryId),
	})
	if err != nil {
		panic(err)
	}

	fireflyTransaction := entities.ConvertTransactionToFireflyTransaction(transaction)

	fireflyTransaction2, r, err :=
		b.FireflyClient.TransactionsApi.UpdateTransaction(context.TODO(), int32(transactionId)).Transaction(*fireflyTransaction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.UpdateTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)

		panic(err)
	}

	transaction2 := entities.ConvertFireflyTransactionToTransaction(&fireflyTransaction2)

	b.UpdateTransaction(c.Message, transaction2)
}

func (b *Bot) deleteTransactionHandler(c *telebot.Callback) {
	fmt.Println(c)
	fmt.Println("deleteTransaction btn pressed")
	err := b.Telebot.Respond(c, &telebot.CallbackResponse{Text: "not implemented yet"})
	if err != nil {
		panic(err)
	}
}

