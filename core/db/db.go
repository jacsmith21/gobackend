package db

import (
	"errors"
	"fmt"

	"github.com/jacsmith21/lukabox"
)

var users = []*lukabox.User{
	{ID: 1, Email: "jacob.smith@unb.ca", FirstName: "Jacob", LastName: "Smith", Archived: false},
	{ID: 2, Email: "j.a.smith@live.ca", FirstName: "Jacob", LastName: "Smith", Archived: false},
	{ID: 3, Email: "jacobsmithunb@gmail.com", FirstName: "Jacob", LastName: "Smith", Archived: false},
}

//CreateUser creates a user in the database
func CreateUser(user *lukabox.User) (string, error) {
	user.ID = users[len(users)-1].ID + 1
	users = append(users, user)
	return fmt.Sprintf("%d", user.ID), nil
}

//GetUsers retrieves a user from the database
func GetUsers() ([]*lukabox.User, error) {
	return users, nil
}

//GetUser retrieves a user from the database
func GetUser(id int) (*lukabox.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}

//UpdateUser updates a user in the datbase
func UpdateUser(id int, user *lukabox.User) (*lukabox.User, error) {
	for i, u := range users {
		if u.ID == id {
			users[i] = user
			return u, nil
		}
	}
	return nil, errors.New("article not found")
}
