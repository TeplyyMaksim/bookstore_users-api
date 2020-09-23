package services

import (
	"github.com/TeplyyMaksim/bookstore_users-api/domain/users"
	"github.com/TeplyyMaksim/bookstore_users-api/utils"
)

func CreateUser(user users.User) (*users.User, *utils.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int) (*users.User, *utils.HttpError) {
	result := users.User{ Id: userId }

	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}