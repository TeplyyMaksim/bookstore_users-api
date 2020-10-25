package users

import (
	"github.com/TeplyyMaksim/bookstore_users-api/datasources/mysql/users_db"
	"github.com/TeplyyMaksim/bookstore_users-api/logger"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser =
		"INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser =
		"SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE id=?;"
	queryUpdateUser =
		"UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	queryDeleteUser =
		"DELETE FROM users WHERE id=?;"
	queryFindByStatus =
		"SELECT id, first_name, last_name, email, date_created, password FROM users WHERE status=?;"
	queryFindByEmailAndPasswordAndStatus =
		"SELECT id, first_name, last_name, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

func (user *User) Get() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		error := errors_utils.NewInternalServerError(err.Error())
		logger.HttpError(error)
		return error
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.DateCreated,
		&user.Status,
		&user.Password,
	); getErr != nil {
		error := mysql_utils.ParseError(getErr)
		logger.HttpError(error)
		return error
	}

	return nil
}

func (user *User) Save() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		error := errors_utils.NewInternalServerError(err.Error())
		logger.HttpError(error)
		return error
	}
	defer stmt.Close()

	insertResult, saveError := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
		user.Status,
		user.Password,
	)
	if saveError != nil {
		return mysql_utils.ParseError(saveError)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		httpError := errors_utils.NewInternalServerError(err.Error())
		logger.HttpError(httpError)
		return httpError
	}

	user.Id = int(userId)

	return nil
}

func (user *User) Update() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		error := errors_utils.NewInternalServerError(err.Error())
		logger.HttpError(error)
		return error
	}
	defer stmt.Close()

	_, updateErr := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Status,
		user.Password,
		user.Id,
	)

	if updateErr != nil {
		httpError := mysql_utils.ParseError(updateErr)
		logger.HttpError(httpError)
		return httpError
	}

	return nil
}

func (user *User) Delete() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		error := errors_utils.NewInternalServerError(err.Error())
		logger.HttpError(error)
		return error
	}
	defer stmt.Close()

	if _, saveError := stmt.Exec(user.Id); saveError != nil {
		httpError := mysql_utils.ParseError(saveError)
		logger.HttpError(httpError)
		return httpError
	}

	return nil
}

func FindByStatus(status string) ([]User, *errors_utils.HttpError) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		error := errors_utils.NewInternalServerError(err.Error())
		logger.HttpError(error)
		return nil, error
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		httpError := errors_utils.NewInternalServerError(err.Error())
		logger.HttpError(httpError)
		return nil, httpError
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.DateCreated,
			&user.Password,
		); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		user.Status = status
		results = append(results, user)
	}

	return results, nil
}

func (user *User) FindByEmailAndPassword() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPasswordAndStatus)
	if err != nil {
		error := errors_utils.NewInternalServerError(err.Error())
		logger.HttpError(error)
		return error
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)

	if getErr := result.Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.DateCreated,
		&user.Status,
	); getErr != nil {
		error := mysql_utils.ParseError(getErr)
		logger.HttpError(error)
		return error
	}

	return nil
}