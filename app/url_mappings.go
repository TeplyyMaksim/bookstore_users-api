package app

import (
	"github.com/TeplyyMaksim/bookstore_users-api/ctrl/ping"
	"github.com/TeplyyMaksim/bookstore_users-api/ctrl/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	//router.GET("/users/search", users.SearchUser)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)

	router.GET("/internal/users/search", users.Search)
}
