package users

import (
	"github.com/TeplyyMaksim/bookstore_users-api/datasources/mysql/users_db"
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
	queryFindUserByStatus =
		"SELECT id, first_name, last_name, email, date_created, password FROM users WHERE status=?;"
)

func (user *User) Get() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
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
		return errors_utils.NewInternalServerError(err.Error())
	}

	user.Id = int(userId)

	return nil
}

func (user *User) Update() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
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
		return mysql_utils.ParseError(updateErr)
	}

	return nil
}

func (user *User) Delete() *errors_utils.HttpError {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, saveError := stmt.Exec(user.Id); saveError != nil {
		return mysql_utils.ParseError(saveError)
	}

	return nil
}

func FindByStatus(status string) ([]User, *errors_utils.HttpError) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors_utils.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors_utils.NewInternalServerError(err.Error())
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