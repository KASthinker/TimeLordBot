package handlers

import (
	"fmt"

	"github.com/KASthinker/TimeLordBot/internal/buttons"
	"github.com/KASthinker/TimeLordBot/internal/data"
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
	task, ok := data.TasksMap[message.Chat.ID]
	if !ok {
		if user.Registered || db.IfUserExists(message.Chat.ID) {
			data.TasksMap[message.Chat.ID] = new(data.Task)
			task = data.TasksMap[message.Chat.ID]
		}
	}

	// User location GPS
	if message.Location != nil {
		if user.Stage == "reg_check_timezone" || user.Stage == "change_timezone_GPS" {
			loctime, tz := methods.TimeZoneGPS(message.Location.Longitude, message.Location.Latitude, user.TimeFormat)
			if user.Timezone == tz {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"The timezone hasn't been changed since you selected the current time zone.")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				sndMsg.ParseMode = "Markdown"
				go data.Bot.Send(sndMsg)
				user.Stage = ""
			} else {
				user.Timezone = tz
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Is your time ") + fmt.Sprintf("*%v*?", loctime)
				sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
				sndMsg.ParseMode = "Markdown"
				go data.Bot.Send(sndMsg)
				if user.Stage == "reg_check_timezone" {
					user.Stage = "reg_finaly"
				} else if user.Stage == "change_timezone_GPS" {
					user.Stage = "update_timezone"
				}
			}
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
		} else {
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Hello! Welcome. Choose your language please:")
			user.Stage = "reg_language"
			sndMsg.ReplyMarkup = buttons.Language()
		}
		go data.Bot.Send(sndMsg)
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

		} else {
			user.Stage = ""
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Action canceled!")
			sndMsg.ParseMode = "Markdown"
		}
		go data.Bot.Send(sndMsg)
		return
	case "send_admin_message":
		if db.IfUserAdmin(message.Chat.ID) {
			user.Stage = "send_admin_message"
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
			sndMsg.Text = "*Enter a message for all users:*"
			sndMsg.ParseMode = "Markdown"
			go data.Bot.Send(sndMsg)
			return
		}
	}

	switch message.Text {
	case "Yes", "Да":
		if user.Stage == "reg_finaly" {
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
			err := db.NewUser(user, message.Chat.ID)
			if err != nil {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Error input timezone.Try again! Enter - /start")

				user.Stage = ""
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Registration completed successfully. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)

				user.Stage = ""
			}
		} else if user.Stage == "update_timezone" {
			err := db.ChangeTimeZone(message.Chat.ID, user.Timezone)
			if err != nil {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Error changing timezone.Try again!")
				sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
				user.Stage = "change_timezone"
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"The time zone has changed. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
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

			user.Stage = "reg_timezone"
		} else if user.Stage == "update_timezone" {
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Try again. Enter your time zone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			user.Stage = "change_timezone"
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
				user.Stage = ""
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Your account has been deleted. Goodbye!")
				delete(data.UserDataMap, message.Chat.ID)
			}
		} else {
			user.Stage = ""
		}
	//-----------------------------------------------------------------\\
	default:
		if user.Registered || db.IfUserExists(message.Chat.ID) {
			if user.Stage == "change_timezone_manually" {
				// Manually change timezone
				loctime, tz, err := methods.TimeZoneManually(message.Text, user.TimeFormat)
				if err != nil {
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Incorrect time zone entered! Try again:")
					sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
					user.Stage = "change_timezone"
				} else {
					if user.Timezone == tz {
						sndMsg.Text = lang.Translate(user.Language, typeText,
							"The timezone hasn't been changed since you selected the current time zone.")
						sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
						sndMsg.ParseMode = "Markdown"
						user.Stage = ""
					} else {
						user.Timezone = tz
						sndMsg.Text = lang.Translate(user.Language, typeText,
							"Is your time ") + fmt.Sprintf("*%v*?", loctime)
						sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
						sndMsg.ParseMode = "Markdown"
						user.Stage = "update_timezone"
					}
				}
				// Enter new task
			} else if user.Stage == "new_task_text" {
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				state, ok := data.StateTime[message.Chat.ID]
				if !ok {
					data.StateTime[message.Chat.ID] = new(data.StateTm)
					state = data.StateTime[message.Chat.ID]
				}
				task.Text = message.Text
				user.Stage = "new_task_time"
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				if user.TimeFormat == 12 {
					state.Hours = 10
					state.Minute = 0
					state.Meridiem = "AM"
					state.Step = 1
					sndMsg.ReplyMarkup = buttons.InputTime12(state)
				} else if user.TimeFormat == 24 {
					state.Hours = 10
					state.Minute = 0
					state.Step = 1
					sndMsg.ReplyMarkup = buttons.InputTime24(state)
				}
				/////////////////////////////////////////////////////////////////////

			} else if user.Stage == "send_admin_message" && db.IfUserAdmin(message.Chat.ID) {
				users, err := db.GetUsers()
				if err != nil {
					sndMsg.Text = "*Error get users data!*"
					user.Stage = ""
					return
				}
				for _, user := range users {
					sndMsg := tgbotapi.NewMessage(user.UserID, "")
					sndMsg.Text = message.Text
					sndMsg.ParseMode = "Markdown"
					go data.Bot.Send(sndMsg)
				}
				user.Stage = ""
				sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
				sndMsg.Text = "*Messages have been sent!*"
				sndMsg.ParseMode = "Markdown"
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				return
			} else {
				data.TasksMap[message.Chat.ID] = new(data.Task)
				user.Stage = ""
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"I don't understand this command!")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			}
		} else {
			if user.Stage == "reg_check_timezone" {
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				// Manually input timezone
				loctime, tz, err := methods.TimeZoneManually(message.Text, user.TimeFormat)
				if err != nil {
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Incorrect time zone entered! Try again:")
					sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
					user.Stage = "reg_timezone"
				} else {
					user.Timezone = tz
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Is your time ") + fmt.Sprintf("*%v*?", loctime)
					sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
					sndMsg.ParseMode = "Markdown"
					user.Stage = "reg_finaly"
				}
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Account not found! Please register! To register, enter /start.")

			}
		}
	}
	sndMsg.ParseMode = "Markdown"
	sndMsg.DisableNotification = true
	go data.Bot.Send(sndMsg)
}
