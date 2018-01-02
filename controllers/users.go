package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetUser(c echo.Context) (err error) {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
