package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	u      User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User ID '%v' not found", id)
}

func AddUser() (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include an ID")
	}

	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func UpdateUser(u User) (User, error) {
	for i, updatedUser := range users {
		if updatedUser.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' could not be updated", u.ID)
}

func RemoveUserByID(id int) error {
	for i, u := range users {
		if id == u.ID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' could not be deleted", u.ID)

}
