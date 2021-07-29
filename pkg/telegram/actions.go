package telegram

import (
	"bytes"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/entities"
	"gopkg.in/tucnak/telebot.v2"
	"strconv"
)

func (b *Bot) RenderTransaction(transaction *entities.FireflyTransaction) {
	keyboard := &telebot.ReplyMarkup{}

	keyboard.Inline(
		keyboard.Row(
			keyboard.Data("Change category", btnUpdateTransactionCategory, strconv.Itoa(transaction.Id), strconv.Itoa(transaction.CategoryId)),
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
	fmt.Fprintf(buf, "💳 %s\n", transaction.AccountName)
	fmt.Fprintf(buf, "💸 %.2f %s\n", transaction.Amount, transaction.AccountCurrencyCode)
	fmt.Fprintf(buf, "%s\n", transaction.GetCategoryName())
	fmt.Fprintf(buf, "📝 %s\n", transaction.Description)
	fmt.Fprintf(buf, "🕓 %s", transaction.Time.Format("02.01.06 15:04"))

	return buf.String()
}

func (b *Bot) RenderTransactionCategoriesKeyboard(transactionId, selectedCategoryId int) {
	keyboard := &telebot.ReplyMarkup{}

	i := 1
	var row []telebot.Btn
	var inline []telebot.Row

	for categoryId, categoryName := range *b.categories {
		categoryIdString := strconv.Itoa(categoryId)
		transactionIdString := strconv.Itoa(transactionId)

		row = append(row, keyboard.Data(categoryName, btnSetTransactionCategory, transactionIdString, categoryIdString))

		if i%2 == 0 { //even
			inline = append(inline, keyboard.Row(row...))
			row = []telebot.Btn{}
		}
		i++
	}

	keyboard.Inline(inline...)


	//message := b.renderTransactionBody(transaction)

	_, err := b.Telebot.Send(b.User, "kek", keyboard)
	if err != nil {
		panic(err)
	}
}
