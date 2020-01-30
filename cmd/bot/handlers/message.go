package handlers

import (
	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// MessageHandler ...
func MessageHandler(message *tgbotapi.Message) {
	sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
	sndMsg.ParseMode = "Markdown"
	sndMsg.DisableNotification = true
	switch message.Command() {
	case "start":
		sndMsg.Text = "Hello!"
		go data.Bot.Send(sndMsg)
	default:
		sndMsg.Text = "I'm don't understand this command!"
		go data.Bot.Send(sndMsg)
	}
}
