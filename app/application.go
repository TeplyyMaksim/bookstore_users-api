package app

import (
	"github.com/TeplyyMaksim/bookstore_users-api/logger"
	"github.com/labstack/echo"
)

var router = echo.New()

func StartApplication() {
	mapUrls()
	logger.Info("Starting application on :8000 port")
	router.Start(":8000")
}
