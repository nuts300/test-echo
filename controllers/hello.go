package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type (
	helloController struct{}

	HelloController interface {
		Hello(echo.Context) error
	}
)

func (h *helloController) Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, this is test-echo")
}

func NewHelloController() HelloController {
	return &helloController{}
}
