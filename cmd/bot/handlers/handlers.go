package handlers

import (
	"fmt"

	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	"github.com/KASthinker/TimeLordBot/internal/buttons"
	db "github.com/KASthinker/TimeLordBot/internal/database"
	lang "github.com/KASthinker/TimeLordBot/internal/localization"
	"github.com/KASthinker/TimeLordBot/internal/methods"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// MessageHandler ...
func MessageHandler(message *tgbotapi.Message) {
	typeText := "message"
	sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
	user, ok := data.UserDataMap[message.Chat.ID]
	if !ok {
		if db.IfUserExists(message.Chat.ID) {
			data.UserDataMap[message.Chat.ID] = new(data.UserData)
			user = data.UserDataMap[message.Chat.ID]
			db.GetUserData(message.Chat.ID, user)
		} else {
			data.UserDataMap[message.Chat.ID] = new(data.UserData)
			user = data.UserDataMap[message.Chat.ID]
			user.Language = "en_EN"
		}
	}

	if message.Location != nil {
		if user.Stage == 2 {
			user.Stage = 3
			loctime, tz := methods.TimeZoneGPS(message.Location.Longitude, message.Location.Latitude)
			user.Timezone = tz
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Is your time ") + fmt.Sprintf("*%v*?", loctime)
			sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
			sndMsg.ParseMode = "Markdown"
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			user.Stage = 0
		}
		return
	}

	switch message.Command() {
	case "start":
		if db.IfUserExists(message.Chat.ID) {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Hello! Good to see you again! Your task list is uploaded.")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Hello! Welcome. Choose your language please:")
			sndMsg.ReplyMarkup = buttons.Language()
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}
		return
	}

	switch message.Text {
	case "Yes":
		if user.Stage == 3 {
			err := db.NewUser(user, message.Chat.ID)
			if err != nil {
				sndMsg.Text = "База данных не доступна! /start"
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = 0
			} else {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Registration completed successfully. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = 0
			}
		} else {
			user.Stage = 0
		}
	case "Да":
		if user.Stage == 3 {
			err := db.NewUser(user, message.Chat.ID)
			if err != nil {
				sndMsg.Text = "База данных не доступна! /start"
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = 0
			} else {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Registration completed successfully. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = 0
			}
		} else {
			user.Stage = 0
		}
	case "No":
		if user.Stage == 3 {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Try again. Enter your time zone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			user.Stage = 1
		}
	case "Нет":
		if user.Stage == 3 {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Try again. Enter your time zone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			user.Stage = 1
		}
		
	default:
		if db.IfUserExists(message.Chat.ID) {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"I don't understand this command!")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			if user.Stage == 2 {
				loctime, tz, err := methods.TimeZoneManualy(message.Text)
				if err != nil {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Incorrect time zone entered! Try again:")
					sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
					user.Stage = 1
					go data.Bot.Send(sndMsg)
					sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				} else {
					user.Timezone = tz
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Is your time ") + fmt.Sprintf("*%v*?", loctime)
					sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
					sndMsg.ParseMode = "Markdown"
					go data.Bot.Send(sndMsg)
					sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					user.Stage = 3
				}
			} else {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Account not found! Please register! To register, enter /start.")
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}
		}
	}
}

// CallbackHandler ...
func CallbackHandler(callback *tgbotapi.CallbackQuery) {
	typeText := "message"
	message := callback.Message
	sndMsg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")
	user, ok := data.UserDataMap[message.Chat.ID]
	if !ok {
		if db.IfUserExists(message.Chat.ID) {
			data.UserDataMap[message.Chat.ID] = new(data.UserData)
			user = data.UserDataMap[message.Chat.ID]
			db.GetUserData(message.Chat.ID, user)
		} else {
			data.UserDataMap[message.Chat.ID] = new(data.UserData)
			user = data.UserDataMap[message.Chat.ID]
			user.Language = "en_EN"
		}
	}

	switch callback.Data {
	case "en_EN":
		if user.Stage == 0 {
			user.Language = callback.Data
			user.Stage = 1
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your timezone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = 0
		}
	case "ru_RU":
		if user.Stage == 0 {
			user.Language = callback.Data
			user.Stage = 1
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your timezone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = 0
		}
	case "input_timezone":
		if user.Stage == 1 {
			user.Stage = 2
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your time zone. Example Los Angeles \"-8\", Moscow \"+3\":")
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = 0
		}
	case "use_GPS":
		if user.Stage == 1 {
			data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID))
			sndMsg := tgbotapi.NewMessage(message.Chat.ID, "") // ReplyMarcup can't change
			user.Stage = 2
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Press button:")
			sndMsg.ReplyMarkup = buttons.SendUserLocation(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			user.Stage = 0
		}

	case "menu":
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Select an action:")
		sndMsg.ReplyMarkup = buttons.Menu(user.Language)
		go data.Bot.Send(sndMsg)
	case "setting":
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Select an action:")
		sndMsg.ReplyMarkup = buttons.Settings(user.Language)
		go data.Bot.Send(sndMsg)
	case "step_back_start":
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Select an action:")
		sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
		go data.Bot.Send(sndMsg)
	}

}
