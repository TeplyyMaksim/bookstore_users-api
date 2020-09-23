package users

import (
	"fmt"
	"github.com/TeplyyMaksim/bookstore_users-api/utils"
)

var (
	usersDB = make(map[int]*User)
)

func (user *User) Get() *utils.HttpError {
	result := usersDB[user.Id]

	if result == nil {
		return utils.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *utils.HttpError {
	current := usersDB[user.Id]

	if current != nil {
		return utils.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	usersDB[user.Id] = user

	return nil
}
