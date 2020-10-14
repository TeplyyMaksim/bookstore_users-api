package users

import (
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"strings"
)

const (
	StatusActive = "ACTIVE"
)

type User struct {
	Id 				int				`json:"id"`
	FirstName 		string			`json:"first_name"`
	LastName 		string			`json:"last_name"`
	Email 			string			`json:"email"`
	DateCreated 	string			`json:"date_created"`
	Status			string			`json:"status"`
	Password		string			`json:"password"`
}

type Users []User

func (user User) Validate() *errors_utils.HttpError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors_utils.NewBadRequestError("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors_utils.NewBadRequestError("Invalid password")
	}

	return nil
}