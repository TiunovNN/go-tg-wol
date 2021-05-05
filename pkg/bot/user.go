package bot

import "errors"

type userState int

const (
	newUser userState = iota
	registeredUser
)

type User struct {
	state       userState
	phoneNumber string
	computerMac string
	name        string
}

var ErrUnknownUser = errors.New("unknown user")

func GetUser(phoneNumber string) (*User, error) {
	if phoneNumber != "" {
		return nil, ErrUnknownUser
	}
	user := &User{registeredUser, phoneNumber, "01:02:03:04:05:06", "gogi"}
	return user, nil
}
