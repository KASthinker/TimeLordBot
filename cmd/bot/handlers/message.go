package handlers

import (
	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/text/message"
	"golang.org/x/text/language"
)

// MessageHandler ...
func MessageHandler(msg *tgbotapi.Message) {
	p := message.NewPrinter(language.Russian)
	sndMsg := tgbotapi.NewMessage(msg.Chat.ID, "")
	sndMsg.ParseMode = "Markdown"
	sndMsg.DisableNotification = true
	switch msg.Command() {
	case "start":
		sndMsg.Text = p.Sprint("Hello!")
		go data.Bot.Send(sndMsg)
	default:
		sndMsg.Text = p.Sprint("I'm don't understand this command!")
		go data.Bot.Send(sndMsg)
	}
}
