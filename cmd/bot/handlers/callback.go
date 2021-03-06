package handlers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/KASthinker/TimeLordBot/internal/buttons"
	"github.com/KASthinker/TimeLordBot/internal/data"
	db "github.com/KASthinker/TimeLordBot/internal/database"
	lang "github.com/KASthinker/TimeLordBot/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

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
			sndMsg.DisableNotification = true
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
	case "delete_task":
		user.Stage = "delete_task_type"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select the type of task you want to delete:")
		sndMsg.ReplyMarkup = buttons.TypeTasks(user.Language)
	case "step_back_menu":
		user.Stage = ""
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select an action:")
		sndMsg.ReplyMarkup = buttons.Menu(user.Language)
	case "personal_tasks":
		user.Stage = "personal_tasks"
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Select a task type:")
		sndMsg.ReplyMarkup = buttons.TypeTasks(user.Language)
	case "today_tasks":
		user.Stage = "today_tasks"
		tasks, err := db.TodayTasks(message.Chat.ID, user.Timezone, user.TimeFormat)
		if err != nil {
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Error getting tasks!")
			user.Stage = ""
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			go data.Bot.Send(sndMsg)
		} else {
			if len(tasks) > 0 {
				data.StartHideMessage = make(map[int64]int)
				data.StartHideMessage[message.Chat.ID] = message.MessageID + 1
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(
						callback.Message.Chat.ID, callback.Message.MessageID))
				sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
				sndMsg.ParseMode = "Markdown"
				sndMsg.DisableNotification = true
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Task list:")
				data.Bot.Send(sndMsg)
				for i := 0; i < len(tasks); i++ {
					sndMsg.Text = tasks[i].GetTask(user.Language)
					if i == len(tasks)-1 {
						sndMsg.ReplyMarkup = buttons.HideButton(user.Language)
					}
					data.Bot.Send(sndMsg)
				}
			} else {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"The task list is empty. Select an action:")
				user.Stage = ""
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			}
		}
		// Type task buttons
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
		} else if user.Stage == "personal_tasks" {
			typeTask := ""
			switch callback.Data {
			case "common_task":
				typeTask = "Common"
			case "everyday_task":
				typeTask = "Everyday"
			case "holiday_task":
				typeTask = "Holiday"
			}
			tasks, err := db.GetTasks(message.Chat.ID, typeTask, user.TimeFormat)
			if err != nil {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Error getting tasks!")
				user.Stage = ""
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				go data.Bot.Send(sndMsg)
			} else {
				if len(tasks) > 0 {
					data.StartHideMessage = make(map[int64]int)
					data.StartHideMessage[message.Chat.ID] = message.MessageID + 1
					go data.Bot.DeleteMessage(
						tgbotapi.NewDeleteMessage(
							callback.Message.Chat.ID, callback.Message.MessageID))
					sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
					sndMsg.ParseMode = "Markdown"
					sndMsg.DisableNotification = true
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"Task list:")
					data.Bot.Send(sndMsg)
					for i := 0; i < len(tasks); i++ {
						sndMsg.Text = tasks[i].GetTask(user.Language)
						if i == len(tasks)-1 {
							sndMsg.ReplyMarkup = buttons.HideButton(user.Language)
						}
						data.Bot.Send(sndMsg)
					}
				} else {
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"The task list is empty. Select an action:")
					user.Stage = ""
					sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				}
			}
		} else if user.Stage == "delete_task_type" {
			typeTask := ""
			switch callback.Data {
			case "common_task":
				typeTask = "Common"
			case "everyday_task":
				typeTask = "Everyday"
			case "holiday_task":
				typeTask = "Holiday"
			}
			tasks, err := db.GetTasks(message.Chat.ID, typeTask, user.TimeFormat)
			if err != nil {
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Error deleting tasks. Select an action:")
				user.Stage = ""
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			} else {
				stateDel, ok := data.StateDelete[message.Chat.ID]
				if !ok {
					data.StateDelete[message.Chat.ID] = new(data.StateDel)
					stateDel = data.StateDelete[message.Chat.ID]
				}
				stateDel.Selected = make(map[int]bool)
				data.DeleteTasksMap[message.Chat.ID] = tasks
				if len(tasks) > 0 {
					user.Stage = "select_delete_tasks"
					go data.Bot.DeleteMessage(
						tgbotapi.NewDeleteMessage(
							callback.Message.Chat.ID, callback.Message.MessageID))
					sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
					stateDel.StartMessage = message.MessageID
					for i := 0; i < len(tasks); i++ {
						sndMsg.Text = fmt.Sprintf("№%v\n%v", i+1, tasks[i].GetTask(user.Language))
						sndMsg.ReplyMarkup = buttons.SelectDelTask(user.Language, i+1, stateDel)
						sndMsg.ParseMode = "Markdown"
						sndMsg.DisableNotification = true
						data.Bot.Send(sndMsg)
					}
					stateDel.StatusMessage = message.MessageID + len(tasks) + 1
					sndMsg.Text = lang.Translate(user.Language, typeText, "Stage:")
					sndMsg.ReplyMarkup = buttons.DelTask(user.Language, 0)
					sndMsg.ParseMode = "Markdown"
					sndMsg.DisableNotification = true
					data.Bot.Send(sndMsg)
				} else {
					sndMsg.Text = lang.Translate(user.Language, typeText,
						"The task list is empty. Select an action:")
					user.Stage = ""
					sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
					sndMsg.ParseMode = "Markdown"
					go data.Bot.Send(sndMsg)
				}
			}
		}
		// Time buttons
	case "upHours": // Не работает сука
		if user.Stage == "new_task_time" {
			if user.TimeFormat == 24 {
				if state.Hours+state.Step < 24 {
					state.Hours += state.Step
				} else {
					state.Hours = state.Hours + state.Step - 24
				}
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				sndMsg.ReplyMarkup = buttons.InputTime24(state)
			} else if user.TimeFormat == 12 {
				if state.Hours+state.Step <= 12 {
					state.Hours += state.Step
				} else {
					state.Hours = state.Hours + state.Step - 12
				}
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				sndMsg.ReplyMarkup = buttons.InputTime12(state)
			}
		}
	case "downHours":
		if user.Stage == "new_task_time" {
			if user.TimeFormat == 24 {
				if state.Hours-state.Step >= 0 {
					state.Hours -= state.Step
				} else {
					state.Hours = state.Hours - state.Step + 24
				}
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				sndMsg.ReplyMarkup = buttons.InputTime24(state)
			} else if user.TimeFormat == 12 {
				if state.Hours-state.Step >= 1 {
					state.Hours -= state.Step
				} else {
					state.Hours = state.Hours - state.Step + 12
				}
				sndMsg.Text = lang.Translate(user.Language, typeText,
					"Select the notification time:")
				sndMsg.ReplyMarkup = buttons.InputTime12(state)
			}
		}
	case "upMinute":
		if user.Stage == "new_task_time" {
			if state.Minute+state.Step < 60 {
				state.Minute += state.Step
			} else {
				state.Minute = state.Minute + state.Step - 60
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the notification time:")
			if user.TimeFormat == 24 {
				sndMsg.ReplyMarkup = buttons.InputTime24(state)
			} else if user.TimeFormat == 12 {
				sndMsg.ReplyMarkup = buttons.InputTime12(state)
			}
		}
	case "downMinute":
		if user.Stage == "new_task_time" {
			if state.Minute-state.Step >= 0 {
				state.Minute -= state.Step
			} else {
				state.Minute = state.Minute - state.Step + 60
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the notification time:")
			if user.TimeFormat == 24 {
				sndMsg.ReplyMarkup = buttons.InputTime24(state)
			} else if user.TimeFormat == 12 {
				sndMsg.ReplyMarkup = buttons.InputTime12(state)
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
			sndMsg.ReplyMarkup = buttons.InputTime12(state)
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
	case "done_delete_task":
		if user.Stage == "select_delete_tasks" {
			stateDel, ok := data.StateDelete[message.Chat.ID]
			if !ok {
				data.StateDelete[message.Chat.ID] = new(data.StateDel)
				stateDel = data.StateDelete[message.Chat.ID]
			}
			if len(stateDel.Selected) > 0 {
				tasks := data.DeleteTasksMap[message.Chat.ID]
				for key := range stateDel.Selected {
					err := db.DeleteTask(message.Chat.ID, tasks[key-1].ID)
					if err != nil {
						sndMsg.Text = strconv.Itoa(key) + " error delete!"
						sndMsg.ParseMode = "Markdown"
						data.Bot.Send(sndMsg)
					}
				}
				for i := stateDel.StartMessage; i <= stateDel.StatusMessage; i++ {
					go data.Bot.DeleteMessage(
						tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, i))
				}
				sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
				sndMsg.Text = lang.Translate(user.Language, typeText, "Deleted!")
				sndMsg.ParseMode = "Markdown"
				sndMsg.DisableNotification = true
				sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
				data.Bot.Send(sndMsg)
				user.Stage = ""
				data.DeleteTasksMap = make(map[int64][]data.Task)
				data.StateDelete = make(map[int64]*data.StateDel)
				return
			}
		}

	case "cancel":
		user.Stage = ""
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Action canceled! Select an action:")
		sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
	case "cancel_delete_task":
		user.Stage = ""
		stateDel, ok := data.StateDelete[message.Chat.ID]
		if !ok {
			data.StateDelete[message.Chat.ID] = new(data.StateDel)
			stateDel = data.StateDelete[message.Chat.ID]
		}
		for i := stateDel.StartMessage; i <= stateDel.StatusMessage; i++ {
			go data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, i))
		}
		sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
		sndMsg.Text = lang.Translate(user.Language, typeText,
			"Action canceled! Select an action:")
		sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
		sndMsg.ParseMode = "Markdown"
		sndMsg.DisableNotification = true
		go data.Bot.Send(sndMsg)
		return
	case "hide":
		if user.Stage == "personal_tasks" || user.Stage == "today_tasks" {
			for i := data.StartHideMessage[message.Chat.ID]; i <= message.MessageID; i++ {
				go data.Bot.DeleteMessage(
					tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, i))
			}
			sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select an action:")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			sndMsg.ParseMode = "Markdown"
			sndMsg.DisableNotification = true
			go data.Bot.Send(sndMsg)
			return
		}
	case "step":
		if user.Stage == "new_task_time" {
			switch state.Step {
			case 1:
				state.Step = 5
			case 5:
				state.Step = 10
			default:
				state.Step = 1
			}
			sndMsg.Text = lang.Translate(user.Language, typeText,
				"Select the notification time:")
			if user.TimeFormat == 24 {
				sndMsg.ReplyMarkup = buttons.InputTime24(state)
			} else if user.TimeFormat == 12 {
				sndMsg.ReplyMarkup = buttons.InputTime12(state)
			}
		}
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
		} else if temp[0] == "hide" && len(temp) == 2 {
			messageID, err := strconv.Atoi(temp[1])
			if err != nil {
				log.Printf("Error Atoi -> %v\n", err)
			}
			go data.Bot.DeleteMessage(tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, messageID))
		} else if temp[0] == "deletetask" && user.Stage == "select_delete_tasks" {
			stateDel, ok := data.StateDelete[message.Chat.ID]
			if !ok {
				data.StateDelete[message.Chat.ID] = new(data.StateDel)
				stateDel = data.StateDelete[message.Chat.ID]
			}
			messageID, _ := strconv.Atoi(temp[1])
			_, ok = stateDel.Selected[messageID]
			if ok {
				delete(stateDel.Selected, messageID)
				sndMsg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID,
					stateDel.StartMessage+messageID,
					*buttons.SelectDelTask(user.Language, messageID, stateDel))
				go data.Bot.Send(sndMsg)
				sndMsg = tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID,
					stateDel.StatusMessage, *buttons.DelTask(user.Language, len(stateDel.Selected)))
				go data.Bot.Send(sndMsg)
			} else {
				stateDel.Selected[messageID] = true
				sndMsg := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID,
					stateDel.StartMessage+messageID,
					*buttons.SelectDelTask(user.Language, messageID, stateDel))
				go data.Bot.Send(sndMsg)
				sndMsg = tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID,
					stateDel.StatusMessage, *buttons.DelTask(user.Language, len(stateDel.Selected)))
				go data.Bot.Send(sndMsg)
			}
			return
		}

	}
	sndMsg.ParseMode = "Markdown"
	go data.Bot.Send(sndMsg)
}
