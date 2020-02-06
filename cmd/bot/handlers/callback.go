package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	task, ok := data.TasksMap[message.Chat.ID]
	if !ok {
		if db.IfUserExists(message.Chat.ID) {
			data.TasksMap[message.Chat.ID] = new(data.Task)
			task = data.TasksMap[message.Chat.ID]
		}
	}

	var state *data.StateTm
	if user.Stage == "new_task_time" {
		state, ok = data.StateTime[message.Chat.ID]
		if !ok {
			data.StateTime[message.Chat.ID] = new(data.StateTm)
			state = data.StateTime[message.Chat.ID]
		}
	}

	switch callback.Data {
	case "en_EN", "ru_RU":
		if user.Stage == "reg_language" {
			user.Language = callback.Data
			user.Stage = "reg_time_format"
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter the time format:")
			sndMsg.ReplyMarkup = buttons.TimeFormat(user.Language)
		} else if user.Stage == "change_language" {
			if user.Language == callback.Data {
				user.Stage = ""
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"The language hasn't been changed since you selected the current language.")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			} else {
				err := db.ChangeLanguage(callback.Message.Chat.ID, callback.Data)
				if err != nil {
					user.Stage = ""
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Error changing language.Try again!")
					sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				} else {
					user.Language = callback.Data
					user.Stage = ""
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Language has changed! Select an action:")
					sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				}
			}
		} else {
			user.Stage = ""
		}
	case "12_hour_clock", "24_hour_clock":
		if user.Stage == "reg_time_format" {
			user.Stage = "reg_timezone"
			user.TimeFormat = data.TimeFormat[callback.Data]
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter your timezone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
		} else if user.Stage == "change_time_format" {
			timefmt := data.TimeFormat[callback.Data]
			if timefmt == user.TimeFormat {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"The time format hasn't been changed since you selected the current format.")
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				user.Stage = ""
			} else {
				err := db.ChangeTimeFormat(message.Chat.ID, timefmt)
				if err != nil {
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Error changing time format.Try again!")
					sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
					user.Stage = ""
				} else {
					user.TimeFormat = timefmt
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"The time format has changed. Select an action:")
					sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
					user.Stage = ""
				}
			}
		}
	case "input_timezone":
		if user.Stage == "reg_timezone" {
			user.Stage = "reg_check_timezone"
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter your time zone. Example Los Angeles \"-8\", Moscow \"+3\":")
		} else if user.Stage == "change_timezone" {
			user.Stage = "change_timezone_manually"
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter your time zone. Example Los Angeles \"-8\", Moscow \"+3\":")
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

		// Start buttons
	case "menu", "setting":
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select an action:")
		if callback.Data == "menu" {
			sndMsg.ReplyMarkup = buttons.Menu(user.Language)
		} else if callback.Data == "setting" {
			sndMsg.ReplyMarkup = buttons.Settings(user.Language)
		}
	case "step_back_start":
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select an action:")
		sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)

		// Setting buttons
	case "delete_my_account":
		user.Stage = "check_delete_my_account"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"If you really want to delete your account, enter without quotation marks: "+
				"“Yes, I really want to delete my account!” "+
				"This will permanently delete all your data!"+
				"\nCancel - /cancel")
		sndMsg.ParseMode = "Markdown"
	case "change_language":
		user.Stage = "change_language"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Choose your language please:")
		sndMsg.ReplyMarkup = buttons.Language()
	case "change_timezone":
		user.Stage = "change_timezone"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Enter your timezone:")
		sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
	case "change_time_format":
		user.Stage = "change_time_format"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Enter the time format:")
		sndMsg.ReplyMarkup = buttons.TimeFormat(user.Language)

	// Menu buttons
	case "new_task":
		user.Stage = "new_task_type"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select the type of new task:")
		sndMsg.ReplyMarkup = buttons.TypeTasks(user.Language)

	// Type task buttons
	case "step_back_menu":
		user.Stage = ""
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select an action:")
		sndMsg.ReplyMarkup = buttons.Menu(user.Language)
	case "common_task", "everyday_task", "holiday_task":
		if user.Stage == "new_task_type" {
			switch callback.Data {
			case "common_task":
				task.TypeTask = "Common"
			case "everyday_task":
				task.TypeTask = "Everyday"
			case "holiday_task":
				task.TypeTask = "Holiday"
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Enter the task text:")
			user.Stage = "new_task_text"
		}

		// Time buttons
	case "upHours": // Не работает сука
		if user.Stage == "new_task_time" {
			if user.TimeFormat == 24 {
				if state.Hours < 23 {
					state.Hours++
				} else {
					state.Hours = 0
				}
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				sndMsg.ReplyMarkup = buttons.InputTime24(state.Hours, state.Minute)
			} else if user.TimeFormat == 12 {
				if state.Hours < 12 {
					state.Hours++
				} else {
					state.Hours = 1
				}
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				sndMsg.ReplyMarkup = buttons.InputTime12(state.Hours, state.Minute, state.Meridiem)
			}
		}
	case "downHours":
		if user.Stage == "new_task_time" {
			if user.TimeFormat == 24 {
				if state.Hours > 0 {
					state.Hours--
				} else {
					state.Hours = 23
				}
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				sndMsg.ReplyMarkup = buttons.InputTime24(state.Hours, state.Minute)
			} else if user.TimeFormat == 12 {
				if state.Hours > 1 {
					state.Hours--
				} else {
					state.Hours = 12
				}
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				sndMsg.ReplyMarkup = buttons.InputTime12(state.Hours, state.Minute, state.Meridiem)
			}
		}
	case "upMinute":
		if user.Stage == "new_task_time" {
			if state.Minute < 59 {
				state.Minute++
			} else {
				state.Minute = 0
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the notification time:")
			if user.TimeFormat == 24 {
				sndMsg.ReplyMarkup = buttons.InputTime24(state.Hours, state.Minute)
			} else if user.TimeFormat == 12 {
				sndMsg.ReplyMarkup = buttons.InputTime12(state.Hours, state.Minute, state.Meridiem)
			}
		}
	case "downMinute":
		if user.Stage == "new_task_time" {
			if state.Minute > 0 {
				state.Minute--
			} else {
				state.Minute = 59
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the notification time:")
			if user.TimeFormat == 24 {
				sndMsg.ReplyMarkup = buttons.InputTime24(state.Hours, state.Minute)
			} else if user.TimeFormat == 12 {
				sndMsg.ReplyMarkup = buttons.InputTime12(state.Hours, state.Minute, state.Meridiem)
			}
		}
	case "changeMeridiem":
		if user.Stage == "new_task_time" {
			if state.Meridiem == "PM" {
				state.Meridiem = "AM"
			} else {
				state.Meridiem = "PM"
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the notification time:")
			sndMsg.ReplyMarkup = buttons.InputTime12(state.Hours, state.Minute, state.Meridiem)
		}
	case "TimeOK":
		if task.TypeTask == "Everyday" {
			weekdays, ok := data.StateWeekdays[message.Chat.ID]
			if !ok {
				data.StateWeekdays[message.Chat.ID] = new(data.StateWd)
				weekdays = data.StateWeekdays[message.Chat.ID]
			}
			if user.TimeFormat == 24 {
				weekdays.Time = fmt.Sprintf("%02d:%02d", state.Hours, state.Minute)
				task.Time = fmt.Sprintf("%02d:%02d", state.Hours, state.Minute)
			} else if user.TimeFormat == 12 {
				weekdays.Time = fmt.Sprintf("%02d:%02d %s", state.Hours, state.Minute, state.Meridiem)
				task.Time = fmt.Sprintf("%02d:%02d %s", state.Hours, state.Minute, state.Meridiem)
			}
			user.Stage = "new_task_weekdays"
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the days of the week you want to receive notifications:")
			weekdays.Selected = make(map[string]bool)
			sndMsg.ReplyMarkup = buttons.InputWeekdays(user.Language, weekdays)
		} else {
			date, ok := data.StateDate[message.Chat.ID]
			if !ok {
				data.StateDate[message.Chat.ID] = new(data.StateDt)
				date = data.StateDate[message.Chat.ID]
			}
			if user.TimeFormat == 24 {
				date.Time = fmt.Sprintf("%02d:%02d", state.Hours, state.Minute)
				task.Time = fmt.Sprintf("%02d:%02d", state.Hours, state.Minute)
			} else if user.TimeFormat == 12 {
				date.Time = fmt.Sprintf("%02d:%02d %s", state.Hours, state.Minute, state.Meridiem)
				task.Time = fmt.Sprintf("%02d:%02d %s", state.Hours, state.Minute, state.Meridiem)
			}
			user.Stage = "new_task_date"
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the date of notification:")
			now := time.Now()
			date.Year, date.Month, _ = now.Date()
			sndMsg.ReplyMarkup = buttons.InputDate(user.Language, date)
		}
	case "nextMonth":
		if user.Stage == "new_task_date" {
			date, ok := data.StateDate[message.Chat.ID]
			if !ok {
				data.StateDate[message.Chat.ID] = new(data.StateDt)
				date = data.StateDate[message.Chat.ID]
			}
			if date.Month != time.December {
				date.Month++
			} else {
				date.Month = time.January
				date.Year++
			}
			date.Selected[0] = 0
			date.Selected[1] = 0
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the date of notification:")
			sndMsg.ReplyMarkup = buttons.InputDate(user.Language, date)
		}
	case "prevMonth":
		if user.Stage == "new_task_date" {
			date, ok := data.StateDate[message.Chat.ID]
			if !ok {
				data.StateDate[message.Chat.ID] = new(data.StateDt)
				date = data.StateDate[message.Chat.ID]
			}
			if date.Month != time.January {
				date.Month--
			} else {
				date.Month = time.December
				date.Year--
			}
			date.Selected[0] = 0
			date.Selected[1] = 0
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the date of notification:")
			sndMsg.ReplyMarkup = buttons.InputDate(user.Language, date)
		}
	case "MonthOK":
		if user.Stage == "new_task_date" {
			date, ok := data.StateDate[message.Chat.ID]
			if !ok {
				data.StateDate[message.Chat.ID] = new(data.StateDt)
				date = data.StateDate[message.Chat.ID]
			}
			if date.Status {
				user.Stage = "new_task_priority"
				task.Date = fmt.Sprintf("%v-%v-%v", date.Year, int(date.Month), date.Day)
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Choose priority:")
				sndMsg.ReplyMarkup = buttons.Priority(user.Language)
				sndMsg.ParseMode = "Markdown"
			}
		}
	case "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun":
		if user.Stage == "new_task_weekdays" {
			weekdays, ok := data.StateWeekdays[message.Chat.ID]
			if !ok {
				data.StateWeekdays[message.Chat.ID] = new(data.StateWd)
				weekdays = data.StateWeekdays[message.Chat.ID]
			}
			if weekdays.Selected[callback.Data] == false {
				weekdays.Selected[callback.Data] = true
			} else {
				delete(weekdays.Selected, callback.Data)
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the days of the week you want to receive notifications:")
			sndMsg.ReplyMarkup = buttons.InputWeekdays(user.Language, weekdays)
		}
	case "WeekdaysOK":
		if user.Stage == "new_task_weekdays" {
			weekdays, ok := data.StateWeekdays[message.Chat.ID]
			if !ok {
				data.StateWeekdays[message.Chat.ID] = new(data.StateWd)
				weekdays = data.StateWeekdays[message.Chat.ID]
			}
			if weekdays.Status {
				user.Stage = "new_task_priority"
				help := []string{}
				for key := range weekdays.Selected {
					help = append(help, key)
				}
				task.WeekDay = strings.Join(help, ", ")
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Choose priority:")
				sndMsg.ReplyMarkup = buttons.Priority(user.Language)
				sndMsg.ParseMode = "Markdown"
				data.StateWeekdays[message.Chat.ID] = new(data.StateWd)
				data.StateWeekdays[message.Chat.ID].Selected = make(map[string]bool)
			}
		}
	case "Do", "Schedule", "Delegate", "Eliminate":
		if user.Stage == "new_task_priority" {
			task.Priority = callback.Data
			sndMsg.Text = fmt.Sprintf("%v%v", lang.Translate(user.Language, typeText, "Save this task?"),
				task.GetTask(user.Language))
			sndMsg.ReplyMarkup = buttons.OKorCancel(user.Language)
			user.Stage = "new_task_finaly"
		}
	case "OKbutton":
		if user.Stage == "new_task_finaly" {
			user.Stage = ""
			err := db.AddNewTask(message.Chat.ID, task)
			if err == nil {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Task added:")
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Failed to add task. Try again.")
			}
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			task = new(data.Task)
			data.StateDate[message.Chat.ID] = new(data.StateDt)
			data.StateTime[message.Chat.ID] = new(data.StateTm)
		}
	case "cancel":
		user.Stage = ""
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Action canceled! Select an action:")
		sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
	default:
		temp := strings.Split(callback.Data, "/")
		if temp[0] == "calendar" && user.Stage == "new_task_date" {
			date, ok := data.StateDate[message.Chat.ID]
			if !ok {
				data.StateDate[message.Chat.ID] = new(data.StateDt)
				date = data.StateDate[message.Chat.ID]
			}
			date.Selected[0], _ = strconv.Atoi(temp[1])
			date.Selected[1], _ = strconv.Atoi(temp[2])
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the date of notification:")
			sndMsg.ReplyMarkup = buttons.InputDate(user.Language, date)
		}

	}
	sndMsg.ParseMode = "Markdown"
	go data.Bot.Send(sndMsg)
}
