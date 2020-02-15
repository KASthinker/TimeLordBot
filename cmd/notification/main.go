package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	db "github.com/KASthinker/TimeLordBot/internal/database"

	"github.com/KASthinker/TimeLordBot/configs"
	"github.com/KASthinker/TimeLordBot/internal/methods"
)

func main() {
	file, err := os.OpenFile("notyfication_log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	log.SetOutput(file)

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
		time.Sleep(time.Minute - time.Duration(sec)*time.Second) // Спать до начала следующей	 минуты
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
			sendMessage(user.UserID, text)
		}
	}

}

func sendMessage(userID int64, text string) {
	http.PostForm(
		fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", configs.GetToken()),
		url.Values{"chat_id": {strconv.Itoa(int(userID))},
			"text": {text}, "parse_mode": {"Markdown"}})
}
