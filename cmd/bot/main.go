package main

import (
	"flag"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/KASthinker/TimeLordBot/cmd/bot/handlers"
	"github.com/KASthinker/TimeLordBot/configs"
	"github.com/KASthinker/TimeLordBot/internal/data"
)

func main() {
	data.UserDataMap = make(map[int64]*data.UserData)
	data.TasksMap = make(map[int64]*data.Task)
	data.DeleteTasksMap = make(map[int64][]data.Task)
	data.StateDate = make(map[int64]*data.StateDt)
	data.StateTime = make(map[int64]*data.StateTm)
	data.StateWeekdays = make(map[int64]*data.StateWd)
	data.StateDelete = make(map[int64]*data.StateDel)

	var debug bool

	flag.BoolVar(&debug, "debug", false, "Usage")
	flag.Parse()

	if !debug {
		file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			os.Exit(1)
		}
		defer file.Close()

		log.SetOutput(file)
	}

	data.Bot, data.Err = tgbotapi.NewBotAPI(configs.GetToken())
	if data.Err != nil {
		log.Fatalln(data.Err)
	}

	data.Bot.Debug = debug

	log.Printf("Authorized on account %s", data.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates, err := data.Bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			go handlers.CallbackHandler(update.CallbackQuery)
		} else if update.Message != nil {
			go handlers.MessageHandler(update.Message)
		}
	}
}
