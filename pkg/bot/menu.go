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
