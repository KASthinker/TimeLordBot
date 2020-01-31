package handlers

import (
	"fmt"
	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	"github.com/KASthinker/TimeLordBot/internal/buttons"
	db "github.com/KASthinker/TimeLordBot/internal/database"
	"github.com/KASthinker/TimeLordBot/internal/methods"
	lang "github.com/KASthinker/TimeLordBot/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// MessageHandler ...
func MessageHandler(message *tgbotapi.Message) {
	typeText := "message"
	sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
	sndMsg.ParseMode = "Markdown"
	sndMsg.DisableNotification = true
	user, ok := data.UserDataMap[message.Chat.ID]
	if !ok {
		if db.IfUserExists(message.Chat.ID) {
			data.UserDataMap[message.Chat.ID] = new(data.UserData)
			user = data.UserDataMap[message.Chat.ID]
			db.GetUserData(message.Chat.ID, user)
			user.Registered = true
		} else {
			data.UserDataMap[message.Chat.ID] = new(data.UserData)
			user = data.UserDataMap[message.Chat.ID]
			user.Language = "en_EN" //Default language
			user.Registered = false
		}
	}
	// task, ok := data.TasksMap[message.Chat.ID]
	// if !ok {
	// 	if user.Registered || db.IfUserExists(message.Chat.ID) {
	// 		data.TasksMap[message.Chat.ID] = new(data.Task)
	// 		task = data.TasksMap[message.Chat.ID]
	// 	}
	// }

	// User location GPS
	if message.Location != nil {
		if user.Stage == "reg_check_timezone" {
			user.Stage = "reg_finaly"
			loctime, tz := methods.TimeZoneGPS(message.Location.Longitude, message.Location.Latitude, user.TimeFormat)
			user.Timezone = tz
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Is your time ") + fmt.Sprintf("*%v*?", loctime)
			sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
			sndMsg.ParseMode = "Markdown"
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			user.Stage = ""
		}
		return
	}

	switch message.Command() {
	case "start":
		if user.Registered || db.IfUserExists(message.Chat.ID) {
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Hello! Good to see you again! Your task list is uploaded.")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Hello! Welcome. Choose your language please:")
			user.Stage = "reg_language"
			sndMsg.ReplyMarkup = buttons.Language()
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}
		return
	case "cancel":
		if user.Registered || db.IfUserExists(message.Chat.ID) {
			user.Stage = ""
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Action canceled! Select an action:")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			sndMsg.ParseMode = "Markdown"
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			user.Stage = ""
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Action canceled!")
			sndMsg.ParseMode = "Markdown"
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}
		return

	}

	switch message.Text {
	case "Yes", "Да":
		if user.Stage == "reg_finaly" {
			err := db.NewUser(user, message.Chat.ID)
			if err != nil {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Error input timezone.Try again! Enter - /start")
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = ""
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Registration completed successfully. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = ""
			}
		} else {
			user.Stage = ""
		}
	case "No", "Нет":
		if user.Stage == "reg_finaly" {
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Try again. Enter your time zone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			user.Stage = "reg_timezone"
		}
	
	//Delete account
	case lang.Translate(user.Language, typeText,
		"Yes, I really want to delete my account!"):
		if user.Stage == "check_delete_my_account" {
			err := db.DeleteUserAccount(message.Chat.ID)
			if err != nil {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Error deleting account. Try again!")
				sndMsg.ReplyMarkup = buttons.StartButtons
				go data.Bot.Send(sndMsg)
				user.Stage = ""
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Your account has been deleted. Goodbye!")
				go data.Bot.Send(sndMsg)
				user.Stage = ""
				user.Registered = false
			}
		} else {
			user.Stage = ""
		}
	//-----------------------------------------------------------------\\

	default:
		if user.Registered || db.IfUserExists(message.Chat.ID) {
			sndMsg.Text = "Найден"
			go data.Bot.Send(sndMsg)
		} else {
			if user.Stage == "reg_check_timezone" {
				// Manually input timezone
				loctime, tz, err := methods.TimeZoneManually(message.Text, user.TimeFormat)
				if err != nil {
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Incorrect time zone entered! Try again:")
					sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
					user.Stage = "reg_timezone"
					go data.Bot.Send(sndMsg)
					sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				} else {
					user.Timezone = tz
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Is your time ") + fmt.Sprintf("*%v*?", loctime)
					sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
					sndMsg.ParseMode = "Markdown"
					go data.Bot.Send(sndMsg)
					sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					user.Stage = "reg_finaly"
				}
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Account not found! Please register! To register, enter /start.")
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}
		}
	}
}
