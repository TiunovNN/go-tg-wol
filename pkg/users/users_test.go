package users

import (
	"testing"
)

func TestUserStorage(t *testing.T) {
	user1 := &User{
		Phone:      "1234",
		MacAddress: "01:02:03:04:05:06",
		Name:       "User1",
	}
	user2 := &User{
		Phone:      "4321",
		MacAddress: "06:05:04:03:02:01",
		Name:       "User2",
	}

	users := []*User{user1, user2}
	storage := NewUserStorage(users)
	result, err := storage.GetUser("1234")
	if err != nil {
		t.Errorf("Occur error %s", err)
	}
	if result != user1 {
		t.Errorf("Unexpected user %v", result)
	}
	result, err = storage.GetUser("4321")
	if err != nil {
		t.Errorf("Occur error %s", err)
	}
	if result != user2 {
		t.Errorf("Unexpected user %v", result)
	}
	result, err = storage.GetUser("42354")
	if err == nil {
		t.Error("Expected error 'unkonw user'")
	}
	if result != nil {
		t.Errorf("Unexpected user %v", result)
	}
}
