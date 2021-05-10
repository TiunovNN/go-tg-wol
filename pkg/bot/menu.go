package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func makeStartMenu() *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("/start")),
	)
	return &keyboard
}

func makeMenu() *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Enable computer")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Sign out")),
	)
	return &keyboard
}

func makeRegisterMenu() *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonContact("Send number")),
	)
	return &keyboard
}

func makeMenuForAdmin() *tgbotapi.ReplyKeyboardMarkup {
	keyboard := makeMenu()
	row := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Admin menu"))
	keyboard.Keyboard = append(keyboard.Keyboard, row)
	return keyboard
}

func makeAdminPanel() *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonContact("Add user")),
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonContact("Delete user")),
	)
	return &keyboard
}
