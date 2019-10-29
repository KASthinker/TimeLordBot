package buttons

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	loc "github.com/KASthinker/TimeLordBot/internal/localization"
)

// StartButtons ...
func StartButtons(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📖 " + loc.TrMess(lang, "Menu"), "menu"),
			tgbotapi.NewInlineKeyboardButtonData(
				"⚙️ " + loc.TrMess(lang, "Setting"), "setting"),
		),
	)
	return &keyboard
}

// Menu ...
func Menu(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📝 " + loc.TrMess(lang, "New task"), "new_task"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🗑 " + loc.TrMess(lang, "Delete task"), "delete_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🧾 " + loc.TrMess(lang, "Tasks for today"), "today_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📜 " + loc.TrMess(lang, "Personal tasks"), "personal_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💼 " + loc.TrMess(lang, "Group tasks"), "group_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"👔 " + loc.TrMess(lang, "Groups"), "groups"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 " + loc.TrMess(lang, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// Settings ...
func Settings(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 " + loc.TrMess(lang, "Timezone"), "change_timezone"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 " + loc.TrMess(lang, "Time format"), "change_format_time"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"‼️ " + loc.TrMess(lang, "Delete account"), "delete_my_account"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 " + loc.TrMess(lang, "Language"), "change_language"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 " + loc.TrMess(lang, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// TypeTasks ...
func TypeTasks(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"👨🏼‍💻 " + loc.TrMess(lang, "Common"), "common_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🏃🏼‍♂️ " + loc.TrMess(lang, "Everyday"), "everyday_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🍸 " + loc.TrMess(lang, "Holiday"), "holiday_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 " + loc.TrMess(lang, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// Groups ...
func Groups(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💼 " + loc.TrMess(lang, "My groups"), "my_groups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🤴 " + loc.TrMess(lang, "Create group"), "create_groups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💣 " + loc.TrMess(lang, "Delete group"), "delete_group"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 " + loc.TrMess(lang, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// InputTimeZone ...
func InputTimeZone(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.TrMess(lang, "Enter manually"), "input_timezone"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.TrMess(lang, "Use GPS"), "use_GPS"),
		),
	)
	return &keyboard
}

// SendUserLocation ...
func SendUserLocation(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonLocation(loc.TrMess(lang, "Submit your location")),
		),
	)
	return &keyboard
}

// Priority ...
func Priority(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, "Do")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, "Schedule")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, "Delegate")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, "Eliminate")),
		),
	)
	return &keyboard
}

//YesORNot ...
func YesORNot(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, "Yes")),
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, "No")),
		),
	)
	return &keyboard
}
