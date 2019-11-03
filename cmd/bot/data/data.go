package data

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// NewUserData ...
type NewUserData struct {
	Stage    int
	Language string
	Timezone string
}

var (
	// NewUserDataMap ...
	NewUserDataMap map[int64]*NewUserData
	// Bot ...
	Bot     *tgbotapi.BotAPI
	// Err ...
	Err     error
)