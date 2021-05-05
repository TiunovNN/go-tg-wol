package bot

import (
	"errors"
	"fmt"

	"github.com/TiunovNN/go-tg-wol/pkg/wol"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ChatEntity struct {
	id    int64
	user  *User
	state userState
}

type ChatStorage struct {
	storage map[int64]*ChatEntity
}

var (
	ErrWrongOperation = errors.New("wrong operation")
)

func NewChatStorage() *ChatStorage {
	return &ChatStorage{
		make(map[int64]*ChatEntity, 1),
	}
}

func (s *ChatStorage) GetChat(chatId int64) *ChatEntity {
	entity, ok := s.storage[chatId]
	if ok {
		return entity
	}
	entity = &ChatEntity{
		chatId,
		nil,
		newUser,
	}
	s.storage[chatId] = entity
	return entity
}

func (e *ChatEntity) ProcessMessage(command *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	switch e.state {
	case newUser:
		if command.Contact != nil {
			message, err := e.Register(command.Contact.PhoneNumber)
			return message, err
		} else {
			message := tgbotapi.NewMessage(e.id, "Please, introduce yourselves!")
			message.ReplyMarkup = makeRegisterMenu()
			return &message, nil
		}
	case registeredUser:
		if command.Text == "Enable computer" {
			return e.SendWakeOnLan()
		} else if command.Text == "Sign out" {
			e.user = nil
			e.state = newUser
			message := tgbotapi.NewMessage(e.id, "Please, introduce yourselves!")
			message.ReplyMarkup = makeRegisterMenu()
			return &message, nil
		} else {
			message := tgbotapi.NewMessage(e.id, "Please, choose one of the commands!")
			message.ReplyMarkup = makeMenu()
			return &message, nil
		}
	}
	return nil, nil
}

func (e *ChatEntity) Register(phoneNumber string) (*tgbotapi.MessageConfig, error) {
	if e.state != newUser {
		return nil, ErrWrongOperation
	}
	user, err := GetUser(phoneNumber)
	if err != nil {
		return nil, err
	}
	e.user = user
	e.state = registeredUser
	text := fmt.Sprintf("Welcome %s!", e.user.name)
	message := tgbotapi.NewMessage(e.id, text)
	message.ReplyMarkup = makeMenu()
	return &message, nil
}

func (e *ChatEntity) SendWakeOnLan() (*tgbotapi.MessageConfig, error) {
	if e.state != registeredUser {
		return nil, ErrWrongOperation
	}
	wol.Send(e.user.computerMac)
	text := fmt.Sprintf("Sending wake on lan to %s!", e.user.computerMac)
	message := tgbotapi.NewMessage(e.id, text)
	message.ReplyMarkup = makeMenu()
	return &message, nil
}
