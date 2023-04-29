package adapter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo) {
	e.GET("/account", hello)
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Text: "Hello, world!"})
}

type response struct {
	Text string `json:"text"`
}
