package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/KASthinker/TimeLordBot/configs"
	"github.com/KASthinker/TimeLordBot/internal/buttons"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(configs.GetToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery != nil{
			cmsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID,"")
			cmsg.Text = update.CallbackQuery.Data
			bot.Send(cmsg)
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message {
			default:
				msg.ReplyMarkup = buttons.YesORNot()
			}

			bot.Send(msg)
		}
	}
}
