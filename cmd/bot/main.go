package main

import (
	"github.com/joho/godotenv"
	"github.com/kvderevyanko/bot/internal/service/product"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	godotenv.Load()

	tokenEnv := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(tokenEnv)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   0,
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help \n"+
			"/list - list products")
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	outputMessageText := "Here all the products are:\n\n"
	products := productService.List()
	for _, p := range products {
		outputMessageText += p.Title
		outputMessageText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessageText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Ты написал "+inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID //ответ на сообщение
	bot.Send(msg)
}
