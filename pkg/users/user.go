package users

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
	Phone      string `json:"phone"`
	MacAddress string `json:"mac_address"`
	Name       string `json:"name"`
	Role       string `json:"role"`
}

func (u *User) String() string {
	return fmt.Sprintf(`User{Phone="%s", MacAddress=%s, Name=%s}`, u.Phone, u.MacAddress, u.Name)
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
		log.Printf("User: %s", user)
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
