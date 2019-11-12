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
	sndMsg.ParseMode = "Markdown"
	sndMsg.DisableNotification = true
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
	task, ok := data.TasksMap[message.Chat.ID]
	if !ok {
		if db.IfUserExists(message.Chat.ID) {
			data.TasksMap[message.Chat.ID] = new(data.Task)
			task = data.TasksMap[message.Chat.ID]
		}
	}

	// User location GPS
	if message.Location != nil {
		if user.Stage == "registration_2" {
			user.Stage = "registration_3"
			loctime, tz := methods.TimeZoneGPS(message.Location.Longitude, message.Location.Latitude, user.TimeFormat)
			user.Timezone = tz
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Is your time ") + fmt.Sprintf("*%v*?", loctime)
			sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
			sndMsg.ParseMode = "Markdown"
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else if user.Stage == "change_timezone_GPS" {
			user.Stage = "change_timezone_1"
			loctime, tz := methods.TimeZoneGPS(message.Location.Longitude, message.Location.Latitude, user.TimeFormat)
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
		go data.Bot.DeleteMessage(
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
	case "Yes", "Да":
		if user.Stage == "registration_3" {
			err := db.NewUser(user, message.Chat.ID)
			if err != nil {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Error input timezone.Try again! Enter - /start")
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
		} else if user.Stage == "change_timezone_1" {
			err := db.ChangeTimeZone(message.Chat.ID, user.Timezone)
			if err != nil {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Error changing timezone.Try again!")
				sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
				go data.Bot.Send(sndMsg)
				user.Stage = "change_timezone"
			} else {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"The time zone has changed. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				user.Stage = ""
			}
		} else {
			user.Stage = ""
		}
	case "No", "Нет":
		if user.Stage == "registration_3" {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Try again. Enter your time zone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			user.Stage = "registration_1"
		} else if user.Stage == "change_timezone_1" {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Try again. Enter your time zone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			user.Stage = "change_timezone"
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
			if user.Stage == "change_timezone_manually" {
				// Manually change timezone
				loctime, tz, err := methods.TimeZoneManually(message.Text, user.TimeFormat)
				if err != nil {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Incorrect time zone entered! Try again:")
					sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
					user.Stage = "change_timezone"
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
					user.Stage = "change_timezone_1"
				}
			} else if user.Stage == "new_task_text" {
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				task.Text = message.Text
				user.Stage = "new_task_time"
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Enter the notification time:")
				sndMsg.ParseMode = "Markdown"
				go data.Bot.Send(sndMsg)
			} else if user.Stage == "new_task_time" {
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				if methods.CheckTime(message.Text) != nil {
					user.Stage = "new_task_time"
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Incorrect time entered. Try again.")
					sndMsg.ParseMode = "Markdown"
					go data.Bot.Send(sndMsg)
				} else {
					task.Time = message.Text
					if task.TypeTask == "Everyday" {
						user.Stage = "new_task_weekday"
						sndMsg.Text = lang.TrMess(user.Language, typeText,
							"Enter weeks numbers separated by commas:")
						sndMsg.ParseMode = "Markdown"
						go data.Bot.Send(sndMsg)
					} else {
						user.Stage = "new_task_date"
						sndMsg.Text = lang.TrMess(user.Language, typeText,
							"Enter the date of notification:")
						sndMsg.ParseMode = "Markdown"
						go data.Bot.Send(sndMsg)
					}
				}
			} else if user.Stage == "new_task_date" {
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				date, err := methods.CheckDate(message.Text)
				if err != nil {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Incorrect date entered. Try again.")
					sndMsg.ParseMode = "Markdown"
					go data.Bot.Send(sndMsg)
				} else {
					task.Date = date
					user.Stage = "new_task_priority"
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Choose priority:")
					sndMsg.ParseMode = "Markdown"
					sndMsg.ReplyMarkup = buttons.Priority(user.Language)
					go data.Bot.Send(sndMsg)
				}
			} else if user.Stage == "new_task_weekday" {
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				weekday := methods.CheckWeekday(message.Text)
				if len(weekday) > 0 {
					task.WeekDay = weekday
					user.Stage = "new_task_priority"
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Choose priority:")
					sndMsg.ParseMode = "Markdown"
					sndMsg.ReplyMarkup = buttons.Priority(user.Language)
					go data.Bot.Send(sndMsg)
				} else {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Weeks numbers are not entered correctly. Try again.")
					sndMsg.ParseMode = "Markdown"
					go data.Bot.Send(sndMsg)
				}
			} else if user.Stage == "new_task_priority" {
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				priority := ""
				ok := false
				if user.Language == "ru_RU" {
					priority, ok = data.Priority[message.Text]
				} else if user.Language == "en_EN" {
					priority = lang.TrMess(user.Language, "buttons", message.Text)
					if len(priority) > 0 {
						ok = true
					}
				}
				if ok {
					user.Stage = ""
					task.Priority = priority
					err := db.AddNewTask(message.Chat.ID, task)
					if err != nil {
						sndMsg.Text = lang.TrMess(user.Language, typeText,
							"Failed to add task. Try again.")
						sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
						sndMsg.ParseMode = "Markdown"
						go data.Bot.Send(sndMsg)
					} else {
						sndMsg.Text = lang.TrMess(user.Language, typeText,
							"Task added:") + task.GetTask(user.Language) +
							lang.TrMess(user.Language, typeText, "Select an action:")
						sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
						sndMsg.ParseMode = "Markdown"
						go data.Bot.Send(sndMsg)
					}
					data.TasksMap[message.Chat.ID] = new(data.Task)
				} else {
					sndMsg.Text = lang.TrMess(user.Language, typeText,
						"Priority entered incorrectly. Try again.")
					sndMsg.ReplyMarkup = buttons.Priority(user.Language)
					sndMsg.ParseMode = "Markdown"
					go data.Bot.Send(sndMsg)
				}
			} else {
				data.TasksMap[message.Chat.ID] = new(data.Task)
				user.Stage = ""
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID-1))
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"I don't understand this command!")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}
		} else {
			if user.Stage == "registration_2" {
				// Manually input timezone
				loctime, tz, err := methods.TimeZoneManually(message.Text, user.TimeFormat)
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
	sndMsg.ParseMode = "Markdown"
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
	task, ok := data.TasksMap[message.Chat.ID]
	if !ok {
		if db.IfUserExists(message.Chat.ID) {
			data.TasksMap[message.Chat.ID] = new(data.Task)
			task = data.TasksMap[message.Chat.ID]
		}
	}

	// Registration Inline buttons
	switch callback.Data {
	case "en_EN", "ru_RU":
		if user.Stage == "reg_language" {
			user.Language = callback.Data
			user.Stage = "reg_time_format"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter the time format:")
			sndMsg.ReplyMarkup = buttons.TimeFormat(user.Language)
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
		} else if user.Stage == "change_timezone" {
			user.Stage = "change_timezone_manually"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your time zone. Example Los Angeles \"-8\", Moscow \"+3\":")
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = ""
		}
	case "use_GPS":
		if user.Stage == "registration_1" {
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID))
			sndMsg := tgbotapi.NewMessage(message.Chat.ID, "") // ReplyMarcup can't change
			user.Stage = "registration_2"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Press button:")
			sndMsg.ReplyMarkup = buttons.SendUserLocation(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else if user.Stage == "change_timezone" {
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID))
			sndMsg := tgbotapi.NewMessage(message.Chat.ID, "") // ReplyMarcup can't change
			user.Stage = "change_timezone_GPS"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Press button:")
			sndMsg.ReplyMarkup = buttons.SendUserLocation(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			user.Stage = ""
		}
	case "12_hour_clock", "24_hour_clock":
		if user.Stage == "reg_time_format" {
			user.Stage = "registration_1"
			if callback.Data == "12_hour_clock" {
				user.TimeFormat = 12
			} else if callback.Data == "24_hour_clock" {
				user.TimeFormat = 24
			}
			sndMsg.Text = lang.TrMess(user.Language, typeText,
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
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"Error changing time format.Try again!")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				user.Stage = ""
			} else {
				sndMsg.Text = lang.TrMess(user.Language, typeText,
					"The time format has changed. Select an action:")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
				user.Stage = ""
			}
		}

	//-----------------------------------------------------------------\\

	// Start buttons
	case "menu", "setting":
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Select an action:")
		sndMsg.ParseMode = "Markdown"
		if callback.Data == "menu" {
			sndMsg.ReplyMarkup = buttons.Menu(user.Language)
		} else if callback.Data == "setting" {
			sndMsg.ReplyMarkup = buttons.Settings(user.Language)
		}
		go data.Bot.Send(sndMsg)
	case "step_back_start":
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Select an action:")
		sndMsg.ParseMode = "Markdown"
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
	case "change_timezone":
		user.Stage = "change_timezone"
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Enter your timezone:")
		sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
		go data.Bot.Send(sndMsg)
	case "change_time_format":
		user.Stage = "change_time_format"
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Enter the time format:")
		sndMsg.ReplyMarkup = buttons.TimeFormat(user.Language)
		go data.Bot.Send(sndMsg)

	// Menu buttons
	case "new_task":
		user.Stage = "new_task"
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Select the type of new task:")
		sndMsg.ReplyMarkup = buttons.TypeTasks(user.Language)
		go data.Bot.Send(sndMsg)
	case "delete_task":
		user.Stage = "delete_task"
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Select the type of task you want to delete:")
		sndMsg.ReplyMarkup = buttons.TypeTasks(user.Language)
		go data.Bot.Send(sndMsg)

	// Type Task buttons
	case "common_task", "everyday_task", "holiday_task":
		if user.Stage == "new_task" {
			if callback.Data == "common_task" {
				task.TypeTask = "Common"
			} else if callback.Data == "everyday_task" {
				task.TypeTask = "Everyday"
			} else if callback.Data == "holiday_task" {
				task.TypeTask = "Holiday"
			}
			task.TypeTask = "Common"
			user.Stage = "new_task_text"
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter the task text:")
			sndMsg.ParseMode = "Markdown"
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = ""
		}
	case "step_back_menu":
		user.Stage = ""
		sndMsg.Text = lang.TrMess(user.Language, typeText,
			"Select an action:")
		sndMsg.ParseMode = "Markdown"
		sndMsg.ReplyMarkup = buttons.Menu(user.Language)
		go data.Bot.Send(sndMsg)
	}

}
