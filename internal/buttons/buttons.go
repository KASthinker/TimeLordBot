package buttons

import (
	"fmt"
	"strconv"
	"time"

	loc "github.com/KASthinker/TimeLordBot/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
)

var typeText string = "buttons"

// StartButtons ...
func StartButtons(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📖 "+loc.Translate(lang, typeText, "Menu"), "menu"),
			tgbotapi.NewInlineKeyboardButtonData(
				"⚙️ "+loc.Translate(lang, typeText, "Setting"), "setting"),
		),
	)
	return &keyboard
}

// Menu ...
func Menu(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📝 "+loc.Translate(lang, typeText, "New task"), "new_task"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🗑 "+loc.Translate(lang, typeText, "Delete task"), "delete_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🧾 "+loc.Translate(lang, typeText, "Tasks for today"), "today_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📜 "+loc.Translate(lang, typeText, "Personal tasks"), "personal_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💼 "+loc.Translate(lang, typeText, "Group tasks"), "group_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"👔 "+loc.Translate(lang, typeText, "Groups"), "groups"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.Translate(lang, typeText, "Back"), "step_back_start"),
		),
	)
	return &keyboard
}

// Settings ...
func Settings(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.Translate(lang, typeText, "Timezone"), "change_timezone"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.Translate(lang, typeText, "Time format"), "change_time_format"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"‼️ "+loc.Translate(lang, typeText, "Delete account"), "delete_my_account"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🏳️ "+loc.Translate(lang, typeText, "Language"), "change_language"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.Translate(lang, typeText, "Back"), "step_back_start"),
		),
	)
	return &keyboard
}

// TimeFormat ...
func TimeFormat(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.Translate(lang, typeText, "12-hour clock"), "12_hour_clock"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕤 "+loc.Translate(lang, typeText, "24-hour clock"), "24_hour_clock"),
		),
	)
	return &keyboard
}

// TypeTasks ...
func TypeTasks(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"👨🏼‍💻 "+loc.Translate(lang, typeText, "Common"), "common_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🏃🏼‍♂️ "+loc.Translate(lang, typeText, "Everyday"), "everyday_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🍸 "+loc.Translate(lang, typeText, "Holiday"), "holiday_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.Translate(lang, typeText, "Back"), "step_back_menu"),
		),
	)
	return &keyboard
}

// Groups ...
func Groups(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💼 "+loc.Translate(lang, typeText, "My groups"), "my_groups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🤴 "+loc.Translate(lang, typeText, "Create group"), "create_groups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💣 "+loc.Translate(lang, typeText, "Delete group"), "delete_group"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.Translate(lang, typeText, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// InputTimeZone ...
func InputTimeZone(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Enter manually"), "input_timezone"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Use GPS"), "use_GPS"),
		),
	)
	return &keyboard
}

// SendUserLocation ...
func SendUserLocation(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonLocation(loc.Translate(lang, typeText, "Submit your location")),
		),
	)
	keyboard.OneTimeKeyboard = true
	return &keyboard
}

// Priority ...
func Priority(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Do"), "Do"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Schedule"), "Schedule"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Delegate"), "Delegate"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Eliminate"), "Eliminate"),
		),
	)
	return &keyboard
}

//YesORNot ...
func YesORNot(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.Translate(lang, typeText, "Yes")),
			tgbotapi.NewKeyboardButton(loc.Translate(lang, typeText, "No")),
		),
	)
	keyboard.OneTimeKeyboard = true
	return &keyboard
}

// Language ...
func Language() *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Russian", "ru_RU"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇺🇸 English", "en_EN"),
		),
	)
	return &keyboard
}

// InputTime24 ...
func InputTime24(hours, minute int) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔼", "upHours"),
			tgbotapi.NewInlineKeyboardButtonData("🔼", "upMinute"),
			tgbotapi.NewInlineKeyboardButtonData(" ", "empty"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%02d", hours), "empty"),
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%02d", minute), "empty"),
			tgbotapi.NewInlineKeyboardButtonData("🆗", "TimeOK"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔽", "downHours"),
			tgbotapi.NewInlineKeyboardButtonData("🔽", "downMinute"),
			tgbotapi.NewInlineKeyboardButtonData(" ", "empty"),
		),
	)

	return &keyboard
}

// InputTime12 ...
func InputTime12(hours, minute int, meridiem string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔼", "upHours"),
			tgbotapi.NewInlineKeyboardButtonData("🔼", "upMinute"),
			tgbotapi.NewInlineKeyboardButtonData("🔼", "changeMeridiem"),
			tgbotapi.NewInlineKeyboardButtonData(" ", " "),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%02d", hours), "empty"),
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%02d", minute), "empty"),
			tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%-2s", meridiem), "empty"),
			tgbotapi.NewInlineKeyboardButtonData("🆗", "TimeOK"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🔽", "downHours"),
			tgbotapi.NewInlineKeyboardButtonData("🔽", "downMinute"),
			tgbotapi.NewInlineKeyboardButtonData("🔽", "changeMeridiem"),
			tgbotapi.NewInlineKeyboardButtonData(" ", "empty"),
		),
	)

	return &keyboard
}

func getCalendar(currentYear int, currentMonth time.Month) [6][7]string {
	now := time.Now()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	calendar := [6][7]string{}

	k := int(firstOfMonth.Day())
	end := int(lastOfMonth.Day())

	j := data.IntWeekday[firstOfMonth.Weekday()]
	for i := 0; i < 6; i++ {
		for ; j < 7 && k <= end; j++ {
			calendar[i][j] = fmt.Sprintf("%02d", k)
			k++
		}
		j = 0
	}

	return calendar
}

// InputDate ...
func InputDate(lang string, date *data.StateDt) *tgbotapi.InlineKeyboardMarkup {
	cld := getCalendar(date.Year, date.Month)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				fmt.Sprintf("%v | %v | %v", date.Time, 
					loc.Translate(lang, typeText, date.Month.String()), date.Year), "-"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(loc.Translate(lang, typeText, "Mon"), " "),
			tgbotapi.NewInlineKeyboardButtonData(loc.Translate(lang, typeText, "Tue"), " "),
			tgbotapi.NewInlineKeyboardButtonData(loc.Translate(lang, typeText, "Wed"), " "),
			tgbotapi.NewInlineKeyboardButtonData(loc.Translate(lang, typeText, "Thu"), " "),
			tgbotapi.NewInlineKeyboardButtonData(loc.Translate(lang, typeText, "Fri"), " "),
			tgbotapi.NewInlineKeyboardButtonData(loc.Translate(lang, typeText, "Sat"), " "),
			tgbotapi.NewInlineKeyboardButtonData(loc.Translate(lang, typeText, "Sun"), " "),
		),
	)

	var rows [][]tgbotapi.InlineKeyboardButton

	if cld[date.Selected[0]][date.Selected[1]] != "" {
		i := date.Selected[0]
		j := date.Selected[1]
		date.Day, _ = strconv.Atoi(cld[i][j])
		cld[i][j] = fmt.Sprintf("·%2s·", cld[i][j])
		date.Status = true
	} else {
		date.Status = false
	}

	for i := 0; i < 6; i++ {
		if i == 5 && cld[5][0] == "" {
			break
		}
		row := []tgbotapi.InlineKeyboardButton(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%4s", cld[i][0]), fmt.Sprintf("calendar/%d/%d", i, 0)),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%4s", cld[i][1]), fmt.Sprintf("calendar/%d/%d", i, 1)),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%4s", cld[i][2]), fmt.Sprintf("calendar/%d/%d", i, 2)),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%4s", cld[i][3]), fmt.Sprintf("calendar/%d/%d", i, 3)),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%4s", cld[i][4]), fmt.Sprintf("calendar/%d/%d", i, 4)),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%4s", cld[i][5]), fmt.Sprintf("calendar/%d/%d", i, 5)),
				tgbotapi.NewInlineKeyboardButtonData(
					fmt.Sprintf("%4s", cld[i][6]), fmt.Sprintf("calendar/%d/%d", i, 6)),
			),
		)
		rows = append(rows, row)
	}

	futter := []tgbotapi.InlineKeyboardButton(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("<", "prevMonth"),
			tgbotapi.NewInlineKeyboardButtonData("OK", "MonthOK"),
			tgbotapi.NewInlineKeyboardButtonData(">", "nextMonth"),
		),
	)

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows...)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, futter)

	return &keyboard
}

// OKorCancel ...
func OKorCancel(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("OK", "OKbutton"),
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Cancel"), "cancel"),
		),
	)
	return &keyboard
}

// InputWeekdays ...
func InputWeekdays(lang string, weekdays *data.StateWd) *tgbotapi.InlineKeyboardMarkup {
	temp := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for i, val := range(temp) {
		_, ok := weekdays.Selected[val]
		if ok {
			temp[i] = fmt.Sprintf("·%s·", loc.Translate(lang, typeText, temp[i]))
			weekdays.Status = true
		} else {
			temp[i] = fmt.Sprintf("%s", loc.Translate(lang, typeText, temp[i]))
		}
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(temp[0], "Mon"),
			tgbotapi.NewInlineKeyboardButtonData(temp[1], "Tue"),
			tgbotapi.NewInlineKeyboardButtonData(temp[2], "Wed"),
			tgbotapi.NewInlineKeyboardButtonData(temp[3], "Thu"),
			tgbotapi.NewInlineKeyboardButtonData(temp[4], "Fri"),
			tgbotapi.NewInlineKeyboardButtonData(temp[5], "Sat"),
			tgbotapi.NewInlineKeyboardButtonData(temp[6], "Sun"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("OK", "WeekdaysOK"),
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Cancel"), "cancel"),
		),
	)
	return &keyboard
}