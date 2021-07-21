package bot

import (
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
)

func (b *Bot) updateTransactionCategoryHandler(c *telebot.Callback) {
	fmt.Println(c)
	fmt.Println("updateTransactionCategory btn pressed")
	err := b.Telebot.Respond(c, &telebot.CallbackResponse{})
	if err != nil {
		panic(err)
	}
}

func (b *Bot) deleteTransactionHandler(c *telebot.Callback) {
	fmt.Println(c)
	fmt.Println("deleteTransaction btn pressed")
	err := b.Telebot.Respond(c, &telebot.CallbackResponse{Text: "not implemented yet"})
	if err != nil {
		panic(err)
	}
}

