package bot

import (
	"fmt"

	"github.com/TiunovNN/go-tg-wol/pkg/users"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func makeMenu() *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Enable computer", "enable_comp")),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Sign out", "sign_out")),
	)
	return &keyboard
}

func makeRegisterMenu() *tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonContact("Send number")),
	)
	return &keyboard
}

func makeMenuForAdmin() *tgbotapi.InlineKeyboardMarkup {
	keyboard := makeMenu()
	row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Admin menu", "admin"))
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	return keyboard
}

func makeAdminPanel() *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Add user", "add")),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Delete user", "delete")),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("Back", "back")),
	)
	return &keyboard
}

func makeDeleteMenu(userList []*users.User, page int) *tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup()
	pages := len(userList)
	page_buttons := make([]tgbotapi.InlineKeyboardButton, 0, pages)
	var text string
	for i := 0; i < pages; i++ {
		if i == page {
			text = fmt.Sprintf("*%d*", i)
		} else {
			text = fmt.Sprintf("%d", i)
		}
		page_buttons = append(page_buttons, tgbotapi.NewInlineKeyboardButtonData(text, "page"))
	}
	page_row := tgbotapi.NewInlineKeyboardRow(page_buttons...)
	if page*5 > cap(userList) {
		page = 1
	}

	high := page * 5
	if high > cap(userList) {
		high = cap(userList)
	}
	for _, user := range userList[(page-1)*5 : high] {
		text := fmt.Sprintf("Name: %s, Phone: %s", user.Name, user.Phone)
		row := tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(text, user.Phone))
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	}
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, page_row)
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Back", "back"),
	))
	return &keyboard
}
