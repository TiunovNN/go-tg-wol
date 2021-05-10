package bot

import (
	"fmt"
	"log"

	"github.com/TiunovNN/go-tg-wol/pkg/wol"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type State interface {
	ProcessMessage(entity *ChatEntity, command *tgbotapi.Message) (*tgbotapi.MessageConfig, error)
}

type NewUserState struct{}
type RegisteredUserState struct{}
type AdminUserState struct{}
type AdminMenuState struct{}

func (s NewUserState) register(entity *ChatEntity, phoneNumber string) (*tgbotapi.MessageConfig, error) {
	if phoneNumber[0] != '+' {
		phoneNumber = "+" + phoneNumber
	}
	user, err := entity.userStorage.GetUser(phoneNumber)
	if err != nil {
		return nil, err
	}
	entity.user = user
	text := fmt.Sprintf("Welcome %s!", user.Name)
	message := tgbotapi.NewMessage(entity.id, text)
	var state State
	if user.Role == "admin" {
		message.ReplyMarkup = makeMenuForAdmin()
		state = AdminUserState{}
	} else {
		message.ReplyMarkup = makeMenu()
		state = RegisteredUserState{}
	}
	entity.UpdateState(&state)
	return &message, nil
}

func (s NewUserState) ProcessMessage(entity *ChatEntity, command *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	log.Printf("NewuserState process message")
	if command.Contact != nil && command.Contact.UserID == command.From.ID {
		message, err := s.register(entity, command.Contact.PhoneNumber)
		return message, err
	} else {
		message := tgbotapi.NewMessage(entity.id, "Please, introduce yourself!")
		message.ReplyMarkup = makeRegisterMenu()
		return &message, nil
	}
}

func (s RegisteredUserState) ProcessMessage(entity *ChatEntity, command *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	if command.Text == "Enable computer" {
		wol.Send(entity.user.MacAddress)
		text := fmt.Sprintf("Sending wake on lan to %s!", entity.user.MacAddress)
		message := tgbotapi.NewMessage(entity.id, text)
		message.ReplyMarkup = makeMenu()
		return &message, nil
	} else if command.Text == "Sign out" {
		entity.user = nil
		var state State = NewUserState{}
		entity.UpdateState(&state)
		message := tgbotapi.NewMessage(entity.id, "Please, introduce yourself!")
		message.ReplyMarkup = makeRegisterMenu()
		return &message, nil
	} else {
		message := tgbotapi.NewMessage(entity.id, "Please, choose one of the commands!")
		message.ReplyMarkup = makeMenu()
		return &message, nil
	}
	return nil, nil
}

func (s AdminUserState) ProcessMessage(entity *ChatEntity, command *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {

	return nil, nil
}

func (s AdminMenuState) ProcessMessage(entity *ChatEntity, command *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	return nil, nil
}
