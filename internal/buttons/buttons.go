package buttons

import (
	loc "github.com/KASthinker/TimeLordBot/localization"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var typeText string = "buttons"

// StartButtons ...
func StartButtons(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📖 "+loc.Translate(lang, typeText, "Menu"), "menu"),
			tgbotapi.NewInlineKeyboardButtonData(
				"⚙️ "+loc.Translate(lang, typeText, "Setting"), "setting"),
		),
	)
	return &keyboard
}

// Menu ...
func Menu(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📝 "+loc.Translate(lang, typeText, "New task"), "new_task"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🗑 "+loc.Translate(lang, typeText, "Delete task"), "delete_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🧾 "+loc.Translate(lang, typeText, "Tasks for today"), "today_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"📜 "+loc.Translate(lang, typeText, "Personal tasks"), "personal_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💼 "+loc.Translate(lang, typeText, "Group tasks"), "group_tasks"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"👔 "+loc.Translate(lang, typeText, "Groups"), "groups"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.Translate(lang, typeText, "Back"), "step_back_start"),
		),
	)
	return &keyboard
}

// Settings ...
func Settings(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.Translate(lang, typeText, "Timezone"), "change_timezone"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.Translate(lang, typeText, "Time format"), "change_time_format"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"‼️ "+loc.Translate(lang, typeText, "Delete account"), "delete_my_account"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🏳️ "+loc.Translate(lang, typeText, "Language"), "change_language"),
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.Translate(lang, typeText, "Back"), "step_back_start"),
		),
	)
	return &keyboard
}

// TimeFormat ...
func TimeFormat(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕑 "+loc.Translate(lang, typeText, "12-hour clock"), "12_hour_clock"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🕤 "+loc.Translate(lang, typeText, "24-hour clock"), "24_hour_clock"),
		),
	)
	return &keyboard
}

// TypeTasks ...
func TypeTasks(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"👨🏼‍💻 "+loc.Translate(lang, typeText, "Common"), "common_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🏃🏼‍♂️ "+loc.Translate(lang, typeText, "Everyday"), "everyday_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🍸 "+loc.Translate(lang, typeText, "Holiday"), "holiday_task"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.Translate(lang, typeText, "Back"), "step_back_menu"),
		),
	)
	return &keyboard
}

// Groups ...
func Groups(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💼 "+loc.Translate(lang, typeText, "My groups"), "my_groups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🤴 "+loc.Translate(lang, typeText, "Create group"), "create_groups"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"💣 "+loc.Translate(lang, typeText, "Delete group"), "delete_group"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"🔙 "+loc.Translate(lang, typeText, "Back"), "step_back"),
		),
	)
	return &keyboard
}

// InputTimeZone ...
func InputTimeZone(lang string) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Enter manually"), "input_timezone"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				loc.Translate(lang, typeText, "Use GPS"), "use_GPS"),
		),
	)
	return &keyboard
}

// SendUserLocation ...
func SendUserLocation(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonLocation(loc.Translate(lang, typeText, "Submit your location")),
		),
	)
	keyboard.OneTimeKeyboard = true
	return &keyboard
}

// Priority ...
func Priority(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.Translate(lang, typeText, "Do")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.Translate(lang, typeText, "Schedule")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.Translate(lang, typeText, "Delegate")),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.Translate(lang, typeText, "Eliminate")),
		),
	)
	keyboard.OneTimeKeyboard = true
	return &keyboard
}

//YesORNot ...
func YesORNot(lang string) *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(loc.Translate(lang, typeText, "Yes")),
			tgbotapi.NewKeyboardButton(loc.Translate(lang, typeText, "No")),
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