package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMessageText := "Here all the products are:\n\n"
	products := c.productService.List()
	for _, p := range products {
		outputMessageText += p.Title
		outputMessageText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessageText)
	c.bot.Send(msg)
}
