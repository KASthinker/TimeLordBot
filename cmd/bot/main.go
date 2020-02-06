package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	"github.com/KASthinker/TimeLordBot/cmd/bot/handlers"
	"github.com/KASthinker/TimeLordBot/configs"
)

func init() {
	data.UserDataMap = make(map[int64]*data.UserData)
	data.TasksMap = make(map[int64]*data.Task)
	data.DeleteTasksMap = make(map[int64][]data.Task)
	data.StateDate = make(map[int64]*data.StateDt)
	data.StateTime = make(map[int64]*data.StateTm)
	data.StateWeekdays = make(map[int64]*data.StateWd)

	data.Bot, data.Err = tgbotapi.NewBotAPI(configs.GetToken())
	if data.Err != nil {
		log.Println(data.Err)
	}
	data.Bot.Debug = true

	log.Printf("Authorized on account %s", data.Bot.Self.UserName)
}

func main() {
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
