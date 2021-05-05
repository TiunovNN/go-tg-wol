package users

import "errors"

type User struct {
	Phone      string
	MacAddress string
	Name       string
}

var ErrUnknownUser = errors.New("unknown user")

type UserStorage struct {
	users map[string]*User
}

func NewUserStorage(users []*User) *UserStorage {
	storage := &UserStorage{
		make(map[string]*User, len(users)),
	}
	for _, user := range users {
		storage.users[user.Phone] = user
	}
	return storage
}

func (s *UserStorage) GetUser(phoneNumber string) (*User, error) {
	user, ok := s.users[phoneNumber]
	if !ok {
		return nil, ErrUnknownUser
	}
	return user, nil
}
