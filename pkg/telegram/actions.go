package telegram

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"gopkg.in/tucnak/telebot.v2"

	"github.com/randaalex/finance_bot/pkg/entities"
)

func (b *Bot) RenderTransaction(transaction *entities.Transaction) {
	keyboard := &telebot.ReplyMarkup{}

	keyboard.Inline(
		keyboard.Row(
			keyboard.Data("Change category", btnUpdateTransactionCategory, strconv.Itoa(int(transaction.Id)), strconv.Itoa(int(transaction.CategoryId))),
			//keyboard.Data("Delete transaction", btnDeleteTransaction, strconv.Itoa(int(transaction.Id))),
		),
	)

	content := b.renderTransactionBody(transaction)

	_, err := b.Telebot.Send(b.User, content, keyboard)
	if err != nil {
		panic(err) // TODO: fix panic
	}
}

func (b *Bot) UpdateTransaction(message telebot.Editable, transaction *entities.Transaction) {
	keyboard := &telebot.ReplyMarkup{}

	keyboard.Inline(
		keyboard.Row(
			keyboard.Data("Change category", btnUpdateTransactionCategory, strconv.Itoa(int(transaction.Id)), strconv.Itoa(int(transaction.CategoryId))),
			//keyboard.Data("Delete transaction", btnDeleteTransaction, strconv.Itoa(int(transaction.Id))),
		),
	)

	content := b.renderTransactionBody(transaction)

	_, err := b.Telebot.Edit(message, content, keyboard)
	if err != nil {
		panic(err) // TODO: fix panic
	}
}

func (b *Bot) renderTransactionBody(t *entities.Transaction) string {
	categoryName := t.CategoryName

	if categoryName == "" {
		categoryName = "❔ Unknown"
	}

	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "Transaction #%d (%s)\n\n", t.Id, t.Type)

	if t.ForeignAmount != 0 {
		fmt.Fprintf(buf, "💸 %.2f %s (%.2f %s)\n", t.Amount, t.CurrencyCode, t.ForeignAmount, t.ForeignCurrencyCode)
	} else {
		fmt.Fprintf(buf, "💸 %.2f %s\n", t.Amount, t.CurrencyCode)
	}

	fmt.Fprintf(buf, "%s\n", categoryName)
	fmt.Fprintf(buf, "💳 %s\n", t.SourceName)
	fmt.Fprintf(buf, "📝 %s\n", t.Description)
	fmt.Fprintf(buf, "🕓 %s", t.IssuedAt.Format("02.01.2006 15:04"))

	return buf.String()
}

func (b *Bot) UpdateTransactionWithCategoriesKeyboard(ctx context.Context, message telebot.Editable, transactionId, _ int) {
	keyboard := &telebot.ReplyMarkup{}

	i := 1
	var row []telebot.Btn
	var inline []telebot.Row

	for _, category := range *b.categories {
		if category.Type != entities.TransactionSplitTypeWithdrawal {
			continue
		}

		categoryId := category.Id
		categoryName := category.Name

		categoryIdString := strconv.Itoa(int(categoryId))
		transactionIdString := strconv.Itoa(transactionId)

		row = append(row, keyboard.Data(categoryName, btnSetTransactionCategory, transactionIdString, categoryIdString))

		if i%2 == 0 { //even
			inline = append(inline, keyboard.Row(row...))
			row = []telebot.Btn{}
		}
		i++
	}

	fireflyTransaction, _, err := b.FireflyClient.TransactionsApi.GetTransaction(ctx, int32(transactionId)).Execute()
	if err != nil {
		panic(err) // TODO: fix panic
	}
	transaction := entities.ConvertFireflyTransactionToTransaction(&fireflyTransaction)

	keyboard.Inline(inline...)
	content := b.renderTransactionBody(transaction)

	_, err = b.Telebot.Edit(message, content, keyboard)
	if err != nil {
		panic(err) // TODO: fix panic
	}
}
