package bot

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/TiunovNN/go-tg-wol/pkg/wol"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type State interface {
	OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error)
	OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error)
}

type NewUserState struct{}
type RegisteredUserState struct{}
type AdminUserState struct{}
type AdminMenuState struct{}
type AddUserName struct{}
type AddUserPhone struct{}
type AddUserMac struct{}
type RemoveUsersList struct{}

var (
	NotImplementedErr = errors.New("not implemented")
)

func (s NewUserState) register(entity *ChatEntity, phoneNumber string) (tgbotapi.Chattable, error) {
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

func (s NewUserState) OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	log.Printf("NewuserState process message")
	command := update.Message
	if command.Contact != nil && command.Contact.UserID == command.From.ID {
		message, err := s.register(entity, command.Contact.PhoneNumber)
		return message, err
	} else {
		message := tgbotapi.NewMessage(entity.id, "Please, introduce yourself!")
		message.ReplyMarkup = makeRegisterMenu()
		return &message, nil
	}
}

func (s NewUserState) OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, NotImplementedErr
}

func (s RegisteredUserState) OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, NotImplementedErr
}

func (s RegisteredUserState) OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	if update.CallbackQuery.Data == "enable_comp" {
		wol.Send(entity.user.MacAddress)
		text := fmt.Sprintf("Sending wake on lan to %s!", entity.user.MacAddress)
		message := tgbotapi.NewEditMessageText(
			entity.id,
			update.CallbackQuery.Message.MessageID,
			text,
		)
		message.ReplyMarkup = makeMenu()
		return &message, nil
	} else if update.CallbackQuery.Data == "sign_out" {
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
}
func (s AdminUserState) OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, NotImplementedErr
}

func (s AdminUserState) OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	log.Printf("admin message: %v", update)
	// log.Printf("admin message: %s", update.Message.Text)
	if update.CallbackQuery.Data == "admin" {
		var state State = AdminMenuState{}
		entity.UpdateState(&state)
		message := tgbotapi.NewEditMessageText(
			entity.id,
			update.CallbackQuery.Message.MessageID,
			"Chose action",
		)
		message.ReplyMarkup = makeAdminPanel()
		return &message, nil

	} else {
		return RegisteredUserState{}.OnCallback(entity, update)
	}
}

func (s AdminMenuState) OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, NotImplementedErr
}

func (s AdminMenuState) OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	data := update.CallbackQuery.Data
	if data == "add" {
		var state State = AddUserName{}
		entity.UpdateState(&state)
		message := tgbotapi.NewEditMessageText(
			entity.id,
			update.CallbackQuery.Message.MessageID,
			"Please, write a name!",
		)
		return &message, nil
	} else if data == "back" {
		var state State = AdminUserState{}
		entity.UpdateState(&state)
		message := tgbotapi.NewEditMessageReplyMarkup(
			entity.id,
			update.CallbackQuery.Message.MessageID,
			*makeMenuForAdmin(),
		)
		return &message, nil
	} else if data == "delete" {
		var state State = RemoveUsersList{}
		entity.UpdateState(&state)
		message := tgbotapi.NewEditMessageReplyMarkup(
			entity.id,
			update.CallbackQuery.Message.MessageID,
			*makeDeleteMenu(entity.userStorage.GetUserList(), 1),
		)
		return &message, nil
	}
	return nil, NotImplementedErr
}

func (s AddUserName) OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, NotImplementedErr
}

func (s AddUserName) OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	name := update.Message.Text
	entity.Data["userName"] = name
	var state State = AddUserPhone{}
	entity.UpdateState(&state)
	message := tgbotapi.NewMessage(entity.id, "Please, write their phone number!")
	return &message, nil
}

func (s AddUserPhone) OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, NotImplementedErr
}

func (s AddUserPhone) OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	phone := update.Message.Text
	entity.Data["userPhone"] = phone
	var state State = AddUserMac{}
	entity.UpdateState(&state)
	message := tgbotapi.NewMessage(entity.id, "Please, write their MAC-address!")
	return &message, nil
}

func (s AddUserMac) OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, NotImplementedErr
}

func (s AddUserMac) OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	mac := update.Message.Text
	var state State = AdminMenuState{}
	entity.UpdateState(&state)

	user, _ := entity.userStorage.AddUser(
		entity.Data["userName"].(string),
		entity.Data["userPhone"].(string),
		mac,
	)
	delete(entity.Data, "userName")
	delete(entity.Data, "userPhone")
	text := fmt.Sprintf("User has been created successfuly: %s", user)
	message := tgbotapi.NewMessage(entity.id, text)
	message.ReplyMarkup = makeAdminPanel()
	return &message, nil
}

func (s RemoveUsersList) OnCallback(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	data := update.CallbackQuery.Data
	if data == "back" {
		var state State = AdminUserState{}
		entity.UpdateState(&state)
		message := tgbotapi.NewEditMessageReplyMarkup(
			entity.id,
			update.CallbackQuery.Message.MessageID,
			*makeMenuForAdmin(),
		)
		return &message, nil
	} else if data == "page" {
		pageNumberStr := update.CallbackQuery.Message.Text
		pageNumber, err := strconv.ParseInt(pageNumberStr, 10, 32)
		if err != nil {
			pageNumber = 1
		}
		message := tgbotapi.NewEditMessageReplyMarkup(
			entity.id,
			update.CallbackQuery.Message.MessageID,
			*makeDeleteMenu(entity.userStorage.GetUserList(), int(pageNumber)),
		)

		return &message, nil
	} else {
		if err := entity.userStorage.DeleteUser(data); err != nil {
			log.Printf("Couldn't delete user %s: %v", data, err)
		}
		message := tgbotapi.NewEditMessageReplyMarkup(
			entity.id,
			update.Message.MessageID,
			*makeDeleteMenu(entity.userStorage.GetUserList(), 1),
		)
		return &message, nil
	}
}
func (s RemoveUsersList) OnMessage(entity *ChatEntity, update *tgbotapi.Update) (tgbotapi.Chattable, error) {
	return nil, NotImplementedErr
}
