package mysql_utils

import (
	"fmt"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors_utils.HttpError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors_utils.NewNotFoundError("No record matches given id")
		}

		return errors_utils.NewInternalServerError("Error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors_utils.NewBadRequestError(fmt.Sprintf("Field should be unique %s", sqlErr.Error()))
	default:
		return errors_utils.NewInternalServerError(sqlErr.Error())
	}
}
