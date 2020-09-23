package users

import (
	"github.com/TeplyyMaksim/bookstore_users-api/utils"
	"net/http"
	"strings"
)

type User struct {
	Id 				int				`json:"id"`
	FirstName 		string			`json:"first_name"`
	LastName 		string			`json:"last_name"`
	Email 			string			`json:"email"`
	DateCreated 	string			`json:"date_created"`
}

func (user *User) Validate() *utils.HttpError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return &utils.HttpError{
			Message: "Wrong Email",
			Status: http.StatusBadRequest,
			Error: "bad_request",
		}
	}

	return nil
}