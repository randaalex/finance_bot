package bot

import (
	"bytes"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/entities"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func (b *Bot) TransactionShow(transaction *entities.FireflyTransaction) {
	keyboard := &telebot.ReplyMarkup{}

	keyboard.Inline(
		keyboard.Row(
			keyboard.Data("Change category", btnUpdateTransactionCategory, strconv.Itoa(transaction.Id)),
			keyboard.Data("Delete transaction", btnDeleteTransaction, strconv.Itoa(transaction.Id)),
		),
	)

	message := b.renderTransactionBody(transaction)

	_, err := b.Telebot.Send(b.User, message, keyboard)
	if err != nil {
		panic(err)
	}
}

func (b *Bot) renderTransactionBody(transaction *entities.FireflyTransaction) string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "Transaction #%d (%s)\n\n", transaction.Id, transaction.Type)
	fmt.Fprintf(buf, "%s\n", transaction.AccountName)
	fmt.Fprintf(buf, "%.2f %s\n", transaction.Amount, transaction.AccountCurrencyCode)
	fmt.Fprintf(buf, "%s\n", transaction.CategoryName)
	fmt.Fprintf(buf, "%s\n", transaction.Description)
	fmt.Fprintf(buf, "%s", transaction.Time.Format("02.01.06 15:04"))

	return buf.String()
}

func (b *Bot) TransactionCategoriesShow(transaction *entities.FireflyTransaction) {
	keyboard := &telebot.ReplyMarkup{}

	keyboard.Inline(
		keyboard.Row(
			keyboard.Data("Change category", btnUpdateTransactionCategory, strconv.Itoa(transaction.Id)),
			keyboard.Data("Delete transaction", btnDeleteTransaction, strconv.Itoa(transaction.Id)),
		),
	)

	message := b.renderTransactionBody(transaction)

	_, err := b.Telebot.Send(b.User, message, keyboard)
	if err != nil {
		panic(err)
	}
}
