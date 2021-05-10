package bot

import (
	"github.com/TiunovNN/go-tg-wol/pkg/users"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ChatEntity struct {
	id          int64
	user        *users.User
	state       *State
	userStorage *users.UserStorage
}

type ChatStorage struct {
	storage     map[int64]*ChatEntity
	userStorage *users.UserStorage
}

func NewChatStorage(userStorage *users.UserStorage) *ChatStorage {
	return &ChatStorage{
		make(map[int64]*ChatEntity, 1),
		userStorage,
	}
}

func (s *ChatStorage) GetChat(chatId int64) *ChatEntity {
	entity, ok := s.storage[chatId]
	if ok {
		return entity
	}

	var state State = NewUserState{}
	entity = &ChatEntity{
		chatId,
		nil,
		&state,
		s.userStorage,
	}
	s.storage[chatId] = entity
	return entity
}

func (e *ChatEntity) ProcessMessage(command *tgbotapi.Message) (*tgbotapi.MessageConfig, error) {
	return (*e.state).ProcessMessage(e, command)
}

func (e *ChatEntity) UpdateState(state *State) {
	e.state = state
}
