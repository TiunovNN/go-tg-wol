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
var ErrUserExists = errors.New("user already exists")

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

func (s *UserStorage) AddUser(name, phone, mac string) (*User, error) {
	if _, ok := s.users[phone]; ok {
		return nil, ErrUserExists
	}

	user := &User{
		Name:       name,
		Phone:      phone,
		MacAddress: mac,
	}
	s.users[phone] = user
	return user, nil

}

func (s *UserStorage) GetUserList() []*User {
	userList := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		if user.Role == "Admin" {
			continue
		}
		userList = append(userList, user)
	}
	return userList
}

func (s *UserStorage) DeleteUser(phoneNumber string) error {
	_, ok := s.users[phoneNumber]
	if !ok {
		return ErrUnknownUser
	}
	delete(s.users, phoneNumber)
	return nil
}
