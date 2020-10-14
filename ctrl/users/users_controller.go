package users

import (
	"github.com/TeplyyMaksim/bookstore_users-api/domain/users"
	"github.com/TeplyyMaksim/bookstore_users-api/services"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func getUserId(userIdParam string) (int, *errors_utils.HttpError) {
	userId, parseIdErr := strconv.Atoi(userIdParam)

	if parseIdErr != nil {
		httpError := errors_utils.NewBadRequestError("Wrong user_id")
		return 0, httpError
	}

	return userId, nil
}

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
	userId, userIdErr := getUserId(c.Param("user_id"))

	if userIdErr != nil {
		return c.JSON(userIdErr.Status, userIdErr)
	}

	user, err := services.GetUser(userId)

	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusOK, user.Marshall(c.Request().Header.Get("X-Public") == "true"))
}

func UpdateUser(c echo.Context) error {
	userId, userIdErr := getUserId(c.Param("user_id"))

	if userIdErr != nil {
		return c.JSON(userIdErr.Status, userIdErr)
	}

	var user users.User
	// c.Bind way of getting user from request
	if err := c.Bind(&user); err != nil {
		response := errors_utils.NewBadRequestError(err.Error())
		return c.JSON(response.Status, response)
	}
	user.Id = userId

	isPartial := c.Request().Method == http.MethodPatch

	result, err := services.UpdateUser(user, isPartial)
	if err != nil {
		c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteUser(c echo.Context) error {
	userId, userIdErr := getUserId(c.Param("user_id"))

	if userIdErr != nil {
		return c.JSON(userIdErr.Status, userIdErr)
	}

	if err := services.DeleteUser(userId); err != nil {
		return c.JSON(err.Status, err)
	}

	serverResponse := struct { Status string `json:"status"` }{
		Status: "deleted",
	}

	return c.JSON(
		http.StatusOK,
		&serverResponse,
	)
}

func Search(c echo.Context) error {
	status := c.QueryParam("status")

	users, err := services.Search(status)

	if err != nil {
		return c.JSON(err.Status, err)
	}




	isPublic := c.Request().Header.Get("X-Public") == "true"
	return c.JSON(http.StatusOK, users.Marshall(isPublic))
}