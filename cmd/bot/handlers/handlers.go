package handlers

import (
	"fmt"
	"github.com/KASthinker/TimeLordBot/cmd/bot/data"
	"github.com/KASthinker/TimeLordBot/internal/buttons"
	db "github.com/KASthinker/TimeLordBot/internal/database"
	lang "github.com/KASthinker/TimeLordBot/internal/localization"
	"github.com/KASthinker/TimeLordBot/internal/methods"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// MessageHandler ...
func MessageHandler(message *tgbotapi.Message) {
	typeText := "message"
	sndMsg := tgbotapi.NewMessage(message.Chat.ID, "")
	user, ok := data.NewUserDataMap[message.Chat.ID]
	if !ok {
		if db.IfUserExists(message.Chat.ID) {
			data.NewUserDataMap[message.Chat.ID] = new(data.NewUserData)
			user = data.NewUserDataMap[message.Chat.ID]
			// user[message.Chat.ID].Language =
			// user[message.Chat.ID].Timezone =
		} else {
			data.NewUserDataMap[message.Chat.ID] = new(data.NewUserData)
			user = data.NewUserDataMap[message.Chat.ID]
			user.Language = "en_EN"
		}
	}

	if message.Location != nil {
		if user.Stage == 2 {
			user.Stage = 3
			loctime, tz := methods.TimeZone(message.Location.Longitude, message.Location.Latitude)
			user.Timezone = tz
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Is your time \"") + fmt.Sprintf("\"%v\"?", loctime)
			sndMsg.ReplyMarkup = buttons.YesORNot(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			user.Stage = 0
		}
		return
	}

	switch message.Command() {
	case "start":
		if db.IfUserExists(message.Chat.ID) {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Hello! Good to see you again! Your task list is uploaded.")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Hello! Welcome. Choose your language please:")
			sndMsg.ReplyMarkup = buttons.Language()
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}
		return
	}
	switch message.Text {
	default:
		if db.IfUserExists(message.Chat.ID) {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"I don't understand this command!")
			sndMsg.ReplyMarkup = buttons.StartButtons(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Account not found! Please register! To register, enter /start.")
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}
	}
}

// CallbackHandler ...
func CallbackHandler(callback *tgbotapi.CallbackQuery) {
	typeText := "message"
	message := callback.Message
	sndMsg := tgbotapi.NewEditMessageText(message.Chat.ID, message.MessageID, "")
	user, ok := data.NewUserDataMap[message.Chat.ID]
	if !ok {
		if db.IfUserExists(message.Chat.ID) {
			data.NewUserDataMap[message.Chat.ID] = new(data.NewUserData)
			user = data.NewUserDataMap[message.Chat.ID]
			// user[message.Chat.ID].Language =
			// user[message.Chat.ID].Timezone =
		} else {
			data.NewUserDataMap[message.Chat.ID] = new(data.NewUserData)
			user = data.NewUserDataMap[message.Chat.ID]
			user.Language = "en_EN"
		}
	}

	switch callback.Data {
	case "en_EN":
		if user.Stage == 0 {
			user.Language = callback.Data
			user.Stage = 1
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your timezone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = 0
		}
	case "ru_RU":
		if user.Stage == 0 {
			user.Language = callback.Data
			user.Stage = 1
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Enter your timezone:")
			sndMsg.ReplyMarkup = buttons.InputTimeZone(user.Language)
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = 0
		}
	case "input_timezone":
		if user.Stage == 1 {
			user.Stage = 0 // 2
			sndMsg.Text = "В разработке!"
			go data.Bot.Send(sndMsg)
		} else {
			user.Stage = 0
		}
	case "use_GPS":
		if user.Stage == 1 {
			data.Bot.DeleteMessage(
				tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID))
			sndMsg := tgbotapi.NewMessage(message.Chat.ID, "") // ReplyMarcup can't change
			user.Stage = 2 // 2
			sndMsg.Text = lang.TrMess(user.Language, typeText,
				"Press button:")
			sndMsg.ReplyMarkup = buttons.SendUserLocation(user.Language)
			go data.Bot.Send(sndMsg)
			sndMsg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		} else {
			user.Stage = 0
		}
	}

}
