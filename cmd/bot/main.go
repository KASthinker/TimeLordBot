package main

import (
	"log"
	"fmt"
	"github.com/KASthinker/TimeLordBot/configs"
	"github.com/KASthinker/TimeLordBot/internal/buttons"
	loc "github.com/KASthinker/TimeLordBot/internal/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	db "github.com/KASthinker/TimeLordBot/internal/database"
)

func main() {
	typeText := "message"
	
	bot, err := tgbotapi.NewBotAPI(configs.GetToken())
	if err != nil {
		log.Println(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(".")
	for update := range updates {
		lang := "en_EN"
		if update.CallbackQuery != nil{
			cmsg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
			switch update.CallbackQuery.Data {
			case "en_EN":
				bot.AnswerCallbackQuery(
					tgbotapi.NewCallback(update.CallbackQuery.ID, "English"))
				cmsg.Text = loc.TrMess(lang, typeText,
					"Enter your time zone.")
				cmsg.ReplyMarkup = buttons.InputTimeZone(lang)
			case "ru_RU":
				bot.AnswerCallbackQuery(
					tgbotapi.NewCallback(update.CallbackQuery.ID, "Russian"))
				cmsg.Text = loc.TrMess(lang, typeText,
					"Enter your time zone.")
				cmsg.ReplyMarkup = buttons.InputTimeZone(lang)
			}
			bot.Send(cmsg)
		}
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

			switch update.Message.Command() {
			case "start":
				if db.IfUserExists(update.Message.Chat.ID) {
					msg.Text = loc.TrMess(lang, typeText, 
						"Hello! Good to see you again! Your task list is uploaded.")
					msg.ReplyMarkup = buttons.StartButtons(lang)
				} else {
					msg.Text = loc.TrMess(lang, typeText, 
						"Hello! Welcome. Choose your language please.")
					msg.ReplyMarkup = buttons.Language()
				}
			default:
				if db.IfUserExists(update.Message.Chat.ID) {
					msg.Text = loc.TrMess(lang, typeText, 
						"I don't understand this command!")
					msg.ReplyMarkup = buttons.StartButtons(lang)
				} else {
					msg.Text = loc.TrMess(lang, typeText, 
						"Account not found! Please register! To register, enter /start.")
				}
			}

			bot.Send(msg)
		}
	}
}
