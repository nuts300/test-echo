package errorHandler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

type Response struct {
	Message string
	Error   string
}

func AppErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	inner := errors.New("Internal error")
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		inner = he.Inner
	}
	c.Logger().Error(code, err)
	res := Response{Message: err.Error(), Error: inner.Error()}
	c.JSON(code, res)
}
