package data

import (
	"fmt"
	"strings"
	"time"

	lang "github.com/KASthinker/TimeLordBot/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	// StateDelete ...
	StateDelete map[int64]*StateDel
	// StateWeekdays ...
	StateWeekdays map[int64]*StateWd
	// StateTime ...
	StateTime map[int64]*StateTm
	// StateDate ...
	StateDate map[int64]*StateDt
	// TimeFormat convert timeformat to int
	TimeFormat = map[string]int{
		"12_hour_clock": 12,
		"24_hour_clock": 24,
	}
	//StartHideMessage ...
	StartHideMessage map[int64]int
	//DeleteTasksMap ...
	DeleteTasksMap map[int64][]Task
	// UserDataMap ...
	UserDataMap map[int64]*UserData
	// TasksMap ...
	TasksMap map[int64]*Task
	// Bot ...
	Bot *tgbotapi.BotAPI
	// Err ...
	Err error
)

// UserData ...
type UserData struct {
	Stage      string
	Language   string
	Timezone   string
	TimeFormat int
	Registered bool
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

// StateTm ...
type StateTm struct {
	Meridiem string
	Hours    int
	Minute   int
	Step     int
}

// StateDt ...
type StateDt struct {
	Time     string
	Month    time.Month
	Selected [2]int
	Year     int
	Day      int
	Status   bool
}

// StateWd ...
type StateWd struct {
	Selected map[string]bool
	Time     string
	Status   bool
}

// StateDel ...
type StateDel struct {
	Selected      map[int]bool
	StartMessage  int
	StatusMessage int
}

// IntWeekday ...
var IntWeekday = map[time.Weekday]int{
	time.Monday:    0,
	time.Tuesday:   1,
	time.Wednesday: 2,
	time.Thursday:  3,
	time.Friday:    4,
	time.Saturday:  5,
	time.Sunday:    6,
}

//GetTask ...
func (task *Task) GetTask(language string) string {
	var temp string
	typeText := "task"
	priority := lang.Translate(language, "buttons", task.Priority)
	if task.TypeTask != "Everyday" {
		date := strings.Split(task.Date, "-")
		temp = fmt.Sprintf("%v %v\n%v %v\n%v %v\n%v %v\n",
			lang.Translate(language, typeText, "Task:"), task.Text,
			lang.Translate(language, typeText, "Time:"), task.Time,
			lang.Translate(language, typeText, "Date:"), fmt.Sprintf(
				"%02s-%02s-%04s", date[2], date[1], date[0]),
			lang.Translate(language, typeText, "Priority:"), priority)
	} else {
		selectedWeekday := strings.Split(task.WeekDay, ", ")
		weekdayList := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
		weekday := []string{}
		for _, val1 := range weekdayList {
			for _, val2 := range selectedWeekday {
				if val1 == val2 {
					weekday = append(weekday, lang.Translate(language, "buttons", val1))
				}
			}
		}
		temp = fmt.Sprintf("%v %v\n%v %v\n%v %v\n%v %v\n",
			lang.Translate(language, typeText, "Task:"), task.Text,
			lang.Translate(language, typeText, "Time:"), task.Time,
			lang.Translate(language, typeText, "Weekdays:"), strings.Join(weekday, ", "),
			lang.Translate(language, typeText, "Priority:"), priority)
	}

	return temp
}
