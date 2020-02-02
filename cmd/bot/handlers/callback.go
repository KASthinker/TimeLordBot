package handlers

import (
	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	"github.com/KASthinker/TimeLordBot/internal/buttons"
	db "github.com/KASthinker/TimeLordBot/internal/database"
	lang "github.com/KASthinker/TimeLordBot/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// CallbackHandler ...
func CallbackHandler(callback *tgbotapi.CallbackQuery) {
	typeText := "message"
	message := callback.Message
	sndMsg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")
	sndMsg.ParseMode = "Markdown"
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
	// 	if db.IfUserExists(message.Chat.ID) {
	// 		data.TasksMap[message.Chat.ID] = new(data.Task)
	// 		task = data.TasksMap[message.Chat.ID]
	// 	}
	// }

	switch callback.Data {
	case "en_EN", "ru_RU":
		if user.Stage == "reg_language" {
			user.Language = callback.Data
			user.Stage = "reg_time_format"
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter the time format:")
			sndMsg.ReplyMarkup = buttons.TimeFormat(user.Language)
			go data.Bot.Send(sndMsg)
		} else if user.Stage == "change_language" {
			err := db.ChangeLanguage(callback.Message.Chat.ID, callback.Data)
			if err != nil {
				user.Stage = ""
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Error changing language.Try again!")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
			} else {
				user.Language = callback.Data
				user.Stage = ""
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Language has changed! Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
			}
		} else {
			user.Stage = ""
		}
	case "12_hour_clock", "24_hour_clock":
		if user.Stage == "reg_time_format" {
			user.Stage = "reg_timezone"
			if callback.Data == "12_hour_clock" {
				user.TimeFormat = 12
			} else if callback.Data == "24_hour_clock" {
				user.TimeFormat = 24
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter your timezone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
		} else if user.Stage == "change_time_format" {
			if callback.Data == "12_hour_clock" {
				user.TimeFormat = 12
			} else if callback.Data == "24_hour_clock" {
				user.TimeFormat = 24
			}
			err := db.ChangeTimeFormat(message.Chat.ID, user.TimeFormat)
			if err != nil {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Error changing time format.Try again!")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				user.Stage = ""
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"The time format has changed. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				user.Stage = ""
			}
		}
	case "input_timezone":
		if user.Stage == "reg_timezone" {
			user.Stage = "reg_check_timezone"
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter your time zone. Example Los Angeles \"-8\", Moscow \"+3\":")
			go data.Bot.Send(sndMsg)
		} else if user.Stage == "change_timezone" {
			user.Stage = "change_timezone_manually"
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter your time zone. Example Los Angeles \"-8\", Moscow \"+3\":")
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = ""
		}
	case "use_GPS":
		if user.Stage == "reg_timezone" || user.Stage == "change_timezone" {
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID))
			sndMsg := tgbotapi.NewMessage(message.Chat.ID, "") // ReplyMarcup can't change
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Press button:")
			sndMsg.ReplyMarkup = buttons.SendUserLocation(user.Language)
			sndMsg.ParseMode = "Markdown"
			go data.Bot.Send(sndMsg)
			if user.Stage == "reg_timezone" {
				user.Stage = "reg_check_timezone"
			} else if user.Stage == "change_timezone" {
				user.Stage = "change_timezone_GPS"
			}

		} else {
			user.Stage = ""
		}
		//------------------------------------------------------------------------------------------------\\
		// Start buttons
	case "menu", "setting":
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select an action:")
		if callback.Data == "menu" {
			sndMsg.ReplyMarkup = buttons.Menu(user.Language)
		} else if callback.Data == "setting" {
			sndMsg.ReplyMarkup = buttons.Settings(user.Language)
		}
		go data.Bot.Send(sndMsg)
	case "step_back_start":
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select an action:")
		sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
		go data.Bot.Send(sndMsg)
		//------------------------------------------------------------------------------------------------\\
		// Setting buttons
	case "delete_my_account":
		user.Stage = "check_delete_my_account"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"If you really want to delete your account, enter without quotation marks: "+
				"“Yes, I really want to delete my account!” "+
				"This will permanently delete all your data!"+
				"\nCancel - /cancel")
		sndMsg.ParseMode = "Markdown"
		go data.Bot.Send(sndMsg)
	case "change_language":
		user.Stage = "change_language"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Choose your language please:")
		sndMsg.ReplyMarkup = buttons.Language()
		go data.Bot.Send(sndMsg)
	case "change_timezone":
		user.Stage = "change_timezone"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Enter your timezone:")
		sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
		go data.Bot.Send(sndMsg)
	case "change_time_format":
		user.Stage = "change_time_format"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Enter the time format:")
		sndMsg.ReplyMarkup = buttons.TimeFormat(user.Language)
		go data.Bot.Send(sndMsg)
	}
}
