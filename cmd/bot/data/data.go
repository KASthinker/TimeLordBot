package data

import (
	"strings"
	"time"

	lang "github.com/KASthinker/TimeLordBot/internal/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// UserData ...
type UserData struct {
	Stage      string
	Language   string
	Timezone   string
	TimeFormat int
}

var (
	// UserDataMap ...
	UserDataMap map[int64]*UserData
	// TasksMap ...
	TasksMap map[int64]*Task
	// Bot ...
	Bot *tgbotapi.BotAPI
	// Err ...
	Err error
)

//Weekday ...
var Weekday = map[string]time.Weekday{
	"1": time.Monday,
	"2": time.Tuesday,
	"3": time.Wednesday,
	"4": time.Thursday,
	"5": time.Friday,
	"6": time.Saturday,
	"7": time.Sunday,
}

// Priority ...
var Priority = map[string]string{
	"Важно и срочно":       "Do",
	"Важно, но не срочно":  "Schedule",
	"Не важно, но срочно":  "Delegate",
	"Не важно и не срочно": "Eliminate",
}

// Task ...
type Task struct {
	ID       int
	TypeTask string
	Text     string
	Date     string
	Time     string
	WeekDay  string
	Priority string
}

//GetTask ...
func (task *Task) GetTask(language string) string {
	typeText := "task"
	weekday := strings.Split(task.WeekDay, ",")
	for i := 0; i < len(weekday); i++ {
		weekday[i] = lang.TrMess(language, "weekday", Weekday[weekday[i]].String())
	}

	wkd := strings.Join(weekday[:], ", ")
	priority := lang.TrMess(language, "buttons", task.Priority)
	if task.TypeTask == "Everyday" {
		temp := lang.TrMess(language, typeText, "Task:") + task.Text + "\n" +
			lang.TrMess(language, typeText, "Time:") + task.Time + "\n" +
			lang.TrMess(language, typeText, "Weekdays:") + wkd + "\n" +
			lang.TrMess(language, typeText, "Priority:") + priority

		return temp
	}
	date := convDate(task.Date, language)
	temp := lang.TrMess(language, typeText, "Task:") + task.Text + "\n" +
		lang.TrMess(language, typeText, "Time:") + task.Time + "\n" +
		lang.TrMess(language, typeText, "Date:") + date + "\n" +
		lang.TrMess(language, typeText, "Priority:") + priority

	return temp
}

func convDate(strDate, language string) string {
	strDate = strings.TrimSpace(strDate)
	tm, _ := time.Parse("02-01-2006", strDate)
	if language == "en_EN" {
		return tm.Format("02/01/2006")
	} else if language == "ru_RU" {
		return tm.Format("02.01.2006")
	}
	return ""
}
