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
	var debug bool

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
		}
		end := time.Now()
		_, _, sec := end.Clock()
		time.Sleep(time.Minute - time.Duration(sec)*time.Second) // Спать до начала следующей минуты
	}
}

func checkTasks(user db.Users) {
	tasks, err := db.TodayTasks(user.UserID, user.TimeZone)
	if err != nil {
		log.Fatalln(err)
	}

	loctime, err := methods.LocTime(user.TimeZone)
	if err != nil {
		log.Fatalln(err)
	}

	for _, task := range tasks {
		if strings.Contains(task.Time, loctime) {
			text := fmt.Sprintf("⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️⚜️\n%v🔰🔰🔰🔰🔰🔰🔰🔰🔰🔰🔰🔰",
				task.GetTask(user.Language))
			sendMessage(user, text)
		}
	}

}

type dataMessage struct {
	Result struct {
		MessageID int `json:"message_id"`
	} `json:"result"`
}

func sendMessage(user db.Users, text string) {
	keyboard := buttons.StartButtons(user.Language)
	byteKeyboard, _ := json.Marshal(keyboard)

	resp, err := http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", configs.GetToken()),
		url.Values{"chat_id": {strconv.Itoa(int(user.UserID))},
			"text": {text}, "parse_mode": {"Markdown"},
			"reply_markup": {string(byteKeyboard)}})

	if err != nil {
		log.Printf("Send message error -> %v", err)
	}

	data := dataMessage{}
	json.NewDecoder(resp.Body).Decode(&data)

	strMessageID := strconv.Itoa(data.Result.MessageID - 1)
	_, err = http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%v/deleteMessage", configs.GetToken()),
		url.Values{"chat_id": {strconv.Itoa(int(user.UserID))}, "message_id": {strMessageID}})

	if err != nil {
		log.Printf("Delete message error -> %v", err)
	}
}
