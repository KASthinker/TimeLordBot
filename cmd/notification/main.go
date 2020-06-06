package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"flag"

	db "github.com/KASthinker/TimeLordBot/internal/database"

	"github.com/KASthinker/TimeLordBot/configs"
	"github.com/KASthinker/TimeLordBot/internal/buttons"
	"github.com/KASthinker/TimeLordBot/internal/methods"
)

func main() {
	log.Println("Started!!!")
	var debug bool

	file, err := os.OpenFile("notyfication_log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	flag.BoolVar(&debug, "debug", false, "Usage")
	flag.Parse()

	if !debug {
		file, err := os.OpenFile("notification_log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			os.Exit(1)
		}
		defer file.Close()

		log.SetOutput(file)
	}

	for {
		users, err := db.GetUsers()
		if err != nil {
			log.Fatalln(err)
		}

		for _, user := range users {
			go checkTasks(user)
			go db.DeleteOldCommonTask(user.UserID, user.TimeZone)
		}
		end := time.Now()
		_, _, sec := end.Clock()
		time.Sleep(time.Minute - time.Duration(sec)*time.Second) // Спать до начала следующей минуты
	}
}

func checkTasks(user db.Users) {
	tasks, err := db.TodayTasks(user.UserID, user.TimeZone, user.TimeFormat)
	if err != nil {
		log.Fatalln(err)
	}

	loctime, err := methods.LocTime(user.TimeZone)
	if err != nil {
		log.Fatalln(err)
	}

	for _, task := range tasks {
		tm, err := methods.ConvTimeFormat(task.Time, 24)
		if err != nil {
			log.Printf("Convert time format error -> %v", err)
		}

		if tm == loctime {
			text := fmt.Sprintf("⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️\n%v🔰🔰🔰🔰🔰🔰🔰🔰🔰🔰🔰🔰",
				task.GetTask(user.Language))
			sendMessage(user, text)
		}

		if task.TypeTask == "Holiday" {
			date := strings.Split(task.Date, "-")
			year, _ := strconv.Atoi(date[0])
			year++
			task.Date = fmt.Sprintf("%d-%s-%s", year, date[1], date[2])
			go db.UpdateHolidayTask(user.UserID, task.ID, task.Date)
		}
	}

}

type dataMessage struct {
	Result struct {
		MessageID int `json:"message_id"`
	} `json:"result"`
}

func sendMessage(user db.Users, text string) {
	resp, err := http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", configs.GetToken()),
		url.Values{
			"chat_id":    {strconv.Itoa(int(user.UserID))},
			"text":       {text},
			"parse_mode": {"Markdown"}})

	if err != nil {
		log.Printf("Send message error -> %v", err)
	}

	data := dataMessage{}
	json.NewDecoder(resp.Body).Decode(&data)

	strMessageID := strconv.Itoa(data.Result.MessageID)
	keyboard := buttons.HideByMessageID(user.Language, strMessageID)
	byteKeyboard, _ := json.Marshal(keyboard)

	_, err = http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%v/editMessageReplyMarkup", configs.GetToken()),
		url.Values{
			"chat_id":      {strconv.Itoa(int(user.UserID))},
			"message_id":   {strMessageID},
			"reply_markup": {string(byteKeyboard)}})

	if err != nil {
		log.Printf("Edit message error -> %v", err)
	}
}

func deleteMessage(userID int64, messageID int) {
	strMessageID := strconv.Itoa(messageID)
	_, err := http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%v/deleteMessage", configs.GetToken()),
		url.Values{"chat_id": {strconv.Itoa(int(userID))}, "message_id": {strMessageID}})
	if err != nil {
		log.Printf("Delete message error -> %v", err)
	}
}
