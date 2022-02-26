package Stores

import (
	"errors"
	"phonebook-api/model"
)

// Implementation of UserStore based on slice.
type arrayUserStore struct {
	list []*model.User
}

// GetUser returns model.User depending on his id.
func (us *arrayUserStore) GetUser(id int) (*model.User, error) {
	for _, user := range us.list {
		if user.UserID == uint(id) {
			return user, nil
		}
	}
	return nil, errors.New("no such id")
}

// GetAllUser returns slice with all users in memory.
func (us *arrayUserStore) GetAllUser() []*model.User {
	return us.list
}

var freeUserId uint = 0

// AddUser adds new user to memory and give him new unused id.
func (us *arrayUserStore) AddUser(name string, city string) {
	us.list = append(us.list, &model.User{UserID: freeUserId, Name: name, City: city})
	freeUserId++
}

// DeleteUser deletes  user from memory or throw error if no user with such id.
func (us *arrayUserStore) DeleteUser(id int) error {
	for i, user := range us.list {
		if user.UserID == uint(id) {
			us.list = append(us.list[:i], us.list[i+1:]...)
			return nil
		}
	}
	return errors.New("no such id")
}
