package users

import (
	"github.com/TeplyyMaksim/bookstore_users-api/domain/users"
	"github.com/TeplyyMaksim/bookstore_users-api/services"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func CreateUser(c echo.Context) error {
	var user users.User

	// c.Bind way of getting user from request
	if err := c.Bind(&user); err != nil {
		response := errors_utils.NewBadRequestError(err.Error())
		return c.JSON(response.Status, response)
	}

	// ioutil.ReadAll + json.Unmarshal way getting user from request
	//userBytes, err := ioutil.ReadAll(c.Request().Body)
	//if err != nil {
	//	return utils.NewBadRequestError(err.Error())
	//}
	//
	//if err = json.Unmarshal(userBytes, &user); err != nil {
	//	return utils.NewBadRequestError(err.Error())
	//}

	result, err := services.CreateUser(user)

	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func GetUser(c echo.Context) error {
	userId, parseIdErr := strconv.Atoi(c.Param("user_id"))

	if parseIdErr != nil {
		httpError := errors_utils.NewBadRequestError("Wrong user_id")
		return c.JSON(httpError.Status, httpError)
	}

	user, err := services.GetUser(userId)

	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusOK, user)
}

func SearchUser(c echo.Context) error {
	return c.String(http.StatusOK, "SearchUser")
}