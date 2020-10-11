package users

import (
	"github.com/TeplyyMaksim/bookstore_users-api/datasources/mysql/users_db"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/date_utils"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"
)

func (user *User) Get() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (user *User) Save() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if saveError != nil {
		return mysql_utils.ParseError(saveError)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}

	user.Id = int(userId)

	return nil
}
