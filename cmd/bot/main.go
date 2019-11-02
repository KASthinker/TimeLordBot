package main

import (
	"log"

	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	"github.com/KASthinker/TimeLordBot/cmd/bot/handlers"
	"github.com/KASthinker/TimeLordBot/configs"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func init() {
	data.NewUserDataMap = make(map[int64]*data.NewUserData)
	data.Bot, data.Err = tgbotapi.NewBotAPI(configs.GetToken())
	if data.Err != nil {
		log.Println(data.Err)
	}
	data.Bot.Debug = true

	log.Printf("Authorized on account %s", data.Bot.Self.UserName)
}

func main() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates, err := data.Bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			go handlers.CallbackHandler(update.CallbackQuery)
		} else if update.Message != nil {
			go handlers.MessageHandler(update.Message)
		}
	}
}
