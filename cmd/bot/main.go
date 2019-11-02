package main

import (
	"log"

	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	"github.com/KASthinker/TimeLordBot/configs"
	"github.com/KASthinker/TimeLordBot/internal/buttons"
	db "github.com/KASthinker/TimeLordBot/internal/database"
	lang "github.com/KASthinker/TimeLordBot/internal/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	// NewUserDataMap ...
	NewUserDataMap map[int64]*data.NewUserData
)

func init() {
	NewUserDataMap = make(map[int64]*data.NewUserData)
	data.Bot, data.Err = tgbotapi.NewBotAPI(configs.GetToken())
	if data.Err != nil {
		log.Println(data.Err)
	}
	data.Bot.Debug = true

	log.Printf("Authorized on account %s", data.Bot.Self.UserName)
}

func main() {
	typeText := "message"
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates, err := data.Bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {

		} else if update.Message != nil {
			message := update.Message
			sndMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			user, ok := NewUserDataMap[message.Chat.ID]
			if !ok {
				if db.IfUserExists(message.Chat.ID) {
					NewUserDataMap[message.Chat.ID] = new(data.NewUserData)
					user = NewUserDataMap[message.Chat.ID]
					// user[message.Chat.ID].Language =
					// user[message.Chat.ID].Timezone =
				} else {
					NewUserDataMap[message.Chat.ID] = new(data.NewUserData)
					user = NewUserDataMap[message.Chat.ID]
					user.Language = "en_EN"
				}
			}

			switch message.Command() {
			case "start":
				if db.IfUserExists(message.Chat.ID) {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Hello! Good to see you again! Your task list is uploaded.")
					sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
					go data.Bot.Send(sndMsg)
					sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					continue
				} else {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Hello! Welcome. Choose your language please:")
					sndMsg.ReplyMarkup = buttons.Language()
					go data.Bot.Send(sndMsg)
					sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					continue
				}
			}
			switch message.Text {
			default:
				if db.IfUserExists(message.Chat.ID) {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"I don't understand this command!")
					sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
					go data.Bot.Send(sndMsg)
					sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				} else {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Account not found! Please register! To register, enter /start.")
					go data.Bot.Send(sndMsg)
					sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				}
			}
		}
	}
}
