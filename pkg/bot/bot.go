package bot

import (
	"log"

	"github.com/TiunovNN/go-tg-wol/pkg/users"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Start(token string, userList []*users.User) error {
	bot, err := tgbotapi.NewBotAPI(token)
	userStorage := users.NewUserStorage(userList)
	chatStorage := NewChatStorage(userStorage)
	if err != nil {
		return err
	}
	log.Printf("Authorized account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updatesChannel, err := bot.GetUpdatesChan(u)
	if err != nil {
		return nil
	}
	for update := range updatesChannel {
		if update.Message != nil {
			chat := chatStorage.GetChat(update.Message.Chat.ID)
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			msg, err := chat.ProcessMessage(&update)
			if err != nil {
				log.Printf("Error %v", err)
				continue
			}
			bot.Send(msg)
			continue
		} else if update.CallbackQuery != nil {
			chat := chatStorage.GetChat(update.CallbackQuery.Message.Chat.ID)
			log.Printf("[%s] %s", update.CallbackQuery.Message.From.UserName, update.CallbackQuery.Message.Text)
			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			msg, err := chat.ProcessCallback(&update)
			if err != nil {
				log.Printf("Error %v", err)
				continue
			}
			bot.Send(msg)
			continue

		}

	}
	return nil
}
