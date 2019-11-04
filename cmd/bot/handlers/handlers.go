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

	// User location GPS
	if message.Location != nil {
		if user.Stage == "registration_2" {
			user.Stage = "registration_3"
			loctime, tz := methods.TimeZoneGPS(message.Location.Longitude, message.Location.Latitude)
			user.Timezone = tz
			sndMsg.Text = lang.TrMess(user.Language, typeText,
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
		if db.IfUserExists(message.Chat.ID) {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Hello! Good to see you again! Your task list is uploaded.")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Hello! Welcome. Choose your language please:")
			user.Stage = "reg_language"
			sndMsg.ReplyMarkup = buttons.Language()
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}
		return
	case "cancel":
		user.Stage = ""
		data.Bot.DeleteMessage(
			tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Action canceled! Select an action:")
		sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
		sndMsg.ParseMode = "Markdown"
		go data.Bot.Send(sndMsg)
		sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		return
	}

	// Registration Replay buttons
	switch message.Text {
	case "Yes":
		if user.Stage == "registration_3" {
			err := db.NewUser(user, message.Chat.ID)
			if err != nil {
				sndMsg.Text = "База данных не доступна! /start"
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = ""
			} else {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Registration completed successfully. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = ""
			}
		} else {
			user.Stage = ""
		}
	case "Да":
		if user.Stage == "registration_3" {
			err := db.NewUser(user, message.Chat.ID)
			if err != nil {
				sndMsg.Text = "База данных не доступна! /start"
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = ""
			} else {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Registration completed successfully. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = ""
			}
		} else {
			user.Stage = ""
		}
	case "No":
		if user.Stage == "registration_3" {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Try again. Enter your time zone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			user.Stage = "registration_1"
		}
	case "Нет":
		if user.Stage == "registration_3" {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Try again. Enter your time zone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			user.Stage = "registration_1"
		}
	//-----------------------------------------------------------------\\

	//Delete account
	case lang.TrMess(user.Language, typeText,
		"Yes, I really want to delete my account!"):
		if user.Stage == "delete_my_account_1" {
			err := db.DeleteUserAccount(message.Chat.ID)
			if err != nil {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Error deleting account. Try again!")
				sndMsg.ReplyMarkup = buttons.StartButtons
				go data.Bot.Send(sndMsg)
				user.Stage = ""
			} else {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Your account has been deleted. Goodbye!")
				go data.Bot.Send(sndMsg)
				user.Stage = ""
			}
		} else {
			user.Stage = ""
		}
	//-----------------------------------------------------------------\\
	default:
		if db.IfUserExists(message.Chat.ID) {
			user.Stage = ""
			data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"I don't understand this command!")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			if user.Stage == "registration_2" {
				// Manually input timezone
				loctime, tz, err := methods.TimeZoneManually(message.Text)
				if err != nil {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Incorrect time zone entered! Try again:")
					sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
					user.Stage = "registration_1"
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
					user.Stage = "registration_3"
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
	// Registration Inline buttons
	switch callback.Data {
	case "en_EN":
		if user.Stage == "reg_language" {
			user.Language = callback.Data
			user.Stage = "registration_1"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your timezone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
		} else if user.Stage == "change_language" {
			err := db.ChangeLanguage(callback.Message.Chat.ID, callback.Data)
			if err != nil {
				user.Stage = ""
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Error changing language.Try again!")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
			} else {
				user.Language = callback.Data
				user.Stage = ""
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Language has changed! Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
			}
		} else {
			user.Stage = ""
		}
	case "ru_RU":
		if user.Stage == "reg_language" {
			user.Language = callback.Data
			user.Stage = "registration_1"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your timezone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
		} else if user.Stage == "change_language" {
			err := db.ChangeLanguage(callback.Message.Chat.ID, callback.Data)
			if err != nil {
				user.Stage = ""
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Error changing language.Try again!")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
			} else {
				user.Language = callback.Data
				user.Stage = ""
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Language has changed! Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
			}
		} else {
			user.Stage = ""
		}
	case "input_timezone":
		if user.Stage == "registration_1" {
			user.Stage = "registration_2"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your time zone. Example Los Angeles \"-8\", Moscow \"+3\":")
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = ""
		}
	case "use_GPS":
		if user.Stage == "registration_1" {
			data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID))
			sndMsg := tgbotapi.NewMessage(message.Chat.ID, "") // ReplyMarcup can't change
			user.Stage = "registration_2"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Press button:")
			sndMsg.ReplyMarkup = buttons.SendUserLocation(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			user.Stage = ""
		}
	//-----------------------------------------------------------------\\

	// Start buttons
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

	//Setting buttons
	case "delete_my_account":
		user.Stage = "delete_my_account_1"
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"If you really want to delete your account, enter without quotation marks: "+
				"“Yes, I really want to delete my account!” "+
				"This will permanently delete all your data!"+
				"\nCancel - /cancel")
		sndMsg.ParseMode = "Markdown"
		go data.Bot.Send(sndMsg)
	case "change_language":
		user.Stage = "change_language"
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Choose your language please:")
		sndMsg.ReplyMarkup = buttons.Language()
		go data.Bot.Send(sndMsg)
	}

}
