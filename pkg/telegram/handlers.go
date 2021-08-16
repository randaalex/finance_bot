package telegram

import (
	"context"
	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
	"strings"
)

func (b *Bot) updateTransactionCategoryHandler(c *telebot.Callback) {
	b.logger.WithFields(logrus.Fields{
		"callback": c,
	}).Info("updateTransactionCategory btn pressed")

	ctx := context.TODO()

	res := strings.Split(c.Data, "|")

	transactionId, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err) // TODO: fix panic
	}

	categoryId, err := strconv.Atoi(res[1])
	if err != nil {
		panic(err) // TODO: fix panic
	}

	b.UpdateTransactionWithCategoriesKeyboard(ctx, c.Message, transactionId, categoryId)

	err = b.Telebot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		panic(err) // TODO: fix panic
	}
}

func (b *Bot) setTransactionCategoryHandler(c *telebot.Callback) {
	b.logger.WithFields(logrus.Fields{
		"callback": c,
	}).Info("setTransactionCategory btn pressed")

	ctx := context.TODO()

	res := strings.Split(c.Data, "|")

	transactionId, err := strconv.Atoi(res[0])
	if err != nil {
		panic(err) // TODO: fix panic
	}

	categoryId, err := strconv.Atoi(res[1])
	if err != nil {
		panic(err) // TODO: fix panic
	}

	rawFireflyTransaction1, r, err := b.FireflyClient.TransactionsApi.GetTransaction(ctx, int32(transactionId)).Execute()
	if err != nil {
		panic(err) // TODO: fix panic
	}

	transaction := entities.ConvertFireflyTransactionToTransaction(&rawFireflyTransaction1)
	transaction.CategoryId = categoryId

	_, err = b.Storage.CreateTransactionsLog(ctx, db.CreateTransactionsLogParams{
		Description: transaction.Description,
		CategoryID:  int32(categoryId),
	})
	if err != nil {
		panic(err) // TODO: fix panic
	}

	fireflyTransactionUpdate := entities.ConvertTransactionToFireflyTransactionUpdate(transaction)

	fireflyTransaction2, r, err :=
		b.FireflyClient.TransactionsApi.UpdateTransaction(context.TODO(), int32(transactionId)).TransactionUpdate(*fireflyTransactionUpdate).Execute()
	if err != nil {
		b.logger.WithFields(logrus.Fields{
			"err": err, "httpResp": r,
		}).Info("Error when calling `TransactionsApi.UpdateTransaction`")

		panic(err) // TODO: fix panic
	}

	transaction2 := entities.ConvertFireflyTransactionToTransaction(&fireflyTransaction2)

	b.UpdateTransaction(c.Message, transaction2)

	err = b.Telebot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		panic(err) // TODO: fix panic
	}
}

func (b *Bot) deleteTransactionHandler(c *telebot.Callback) {
	b.logger.WithFields(logrus.Fields{
		"callback": c,
	}).Info("deleteTransaction btn pressed")

	err := b.Telebot.Respond(c, &telebot.CallbackResponse{Text: "not implemented yet"})
	if err != nil {
		panic(err) // TODO: fix panic
	}
}
