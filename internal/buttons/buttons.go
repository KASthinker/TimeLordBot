package buttons

import (
	loc "github.com/KASthinker/TimeLordBot/internal/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var typeText string = "buttons"

// StartButtons ...
func StartButtons(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📖 "+loc.TrMess(lang, typeText, "Menu"), "menu"),
			tgbotapi.NewInlineKeyboardButtonData(
				"⚙️ "+loc.TrMess(lang, typeText, "Setting"), "setting"),
		),
	)
	return &keyboard
}

// Menu ...
func Menu(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📝 "+loc.TrMess(lang, typeText, "New task"), "new_task"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🗑 "+loc.TrMess(lang, typeText, "Delete task"), "delete_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🧾 "+loc.TrMess(lang, typeText, "Tasks for today"), "today_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📜 "+loc.TrMess(lang, typeText, "Personal tasks"), "personal_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💼 "+loc.TrMess(lang, typeText, "Group tasks"), "group_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"👔 "+loc.TrMess(lang, typeText, "Groups"), "groups"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.TrMess(lang, typeText, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// Settings ...
func Settings(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.TrMess(lang, typeText, "Timezone"), "change_timezone"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.TrMess(lang, typeText, "Time format"), "change_format_time"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"‼️ "+loc.TrMess(lang, typeText, "Delete account"), "delete_my_account"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.TrMess(lang, typeText, "Language"), "change_language"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.TrMess(lang, typeText, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// TypeTasks ...
func TypeTasks(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"👨🏼‍💻 "+loc.TrMess(lang, typeText, "Common"), "common_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🏃🏼‍♂️ "+loc.TrMess(lang, typeText, "Everyday"), "everyday_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🍸 "+loc.TrMess(lang, typeText, "Holiday"), "holiday_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.TrMess(lang, typeText, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// Groups ...
func Groups(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💼 "+loc.TrMess(lang, typeText, "My groups"), "my_groups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🤴 "+loc.TrMess(lang, typeText, "Create group"), "create_groups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💣 "+loc.TrMess(lang, typeText, "Delete group"), "delete_group"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.TrMess(lang, typeText, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// InputTimeZone ...
func InputTimeZone(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.TrMess(lang, typeText, "Enter manually"), "input_timezone"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.TrMess(lang, typeText, "Use GPS"), "use_GPS"),
		),
	)
	return &keyboard
}

// SendUserLocation ...
func SendUserLocation(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonLocation(loc.TrMess(lang, typeText, "Submit your location")),
		),
	)
	return &keyboard
}

// Priority ...
func Priority(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, typeText, "Do")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, typeText, "Schedule")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, typeText, "Delegate")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, typeText, "Eliminate")),
		),
	)
	keyboard.OneTimeKeyboard = true
	return &keyboard
}

//YesORNot ...
func YesORNot(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, typeText, "Yes")),
			tgbotapi.NewKeyboardButton(loc.TrMess(lang, typeText, "No")),
		),
	)
	keyboard.OneTimeKeyboard = true
	return &keyboard
}


// Language ...
func Language() *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Russian", "ru_RU"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🇺🇸 English", "en_EN"),
		),
	)
	return &keyboard
}