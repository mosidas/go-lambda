package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func initApp() *echo.Echo {
	log.Println("Server started")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}
