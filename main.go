package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "here is the registration form")
	})

	//e.Logger.Fatal(e.Start(":8080"))
	e.HideBanner = true
	e.Start(":8080")
}
