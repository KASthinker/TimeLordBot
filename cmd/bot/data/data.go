package data

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// UserData ...
type UserData struct {
	Stage    int
	Language string
	Timezone string
}

var (
	// UserDataMap ...
	UserDataMap map[int64]*UserData
	// Bot ...
	Bot     *tgbotapi.BotAPI
	// Err ...
	Err     error
)
