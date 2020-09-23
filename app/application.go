package app

import "github.com/labstack/echo"

var router = echo.New()

func StartApplication() {
	mapUrls()
	router.Start(":1323")
}
