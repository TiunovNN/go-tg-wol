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
		if update.Message == nil {
			continue
		}
		chat := chatStorage.GetChat(update.Message.Chat.ID)
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		msg, err := chat.ProcessMessage(update.Message)
		if err != nil {
			log.Printf("Error %v", err)
			continue
		}
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
	return nil
}
