package app

import (
	"github.com/TeplyyMaksim/bookstore_users-api/ctrl/ping"
	"github.com/TeplyyMaksim/bookstore_users-api/ctrl/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/search", users.SearchUser)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
