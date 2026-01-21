package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/kvderevyanko/bot/internal/app/commands"
	"github.com/kvderevyanko/bot/internal/service/product"
	"log"
	"os"
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
	commander := commands.NewCommander(bot, productService)
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "help":
				commander.Help(update.Message)
			case "list":
				commander.List(update.Message)
			default:
				commander.Default(update.Message)
			}
		}
	}
}
