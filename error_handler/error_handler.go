package errorHandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/nuts300/test-echo/app_logger"
)

var logger = appLogger.GetLogger()

type errorResponse struct {
	Message string `json:"message" yaml:"message"`
	Error   string `json:"error" yaml:"error"`
}

func AppErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	inner := errors.New("Internal error")
	message := err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		inner = he.Inner
		message = he.Message.(string)
	}
	statusMessge := fmt.Sprint("[status]", code)
	ridMessage := fmt.Sprint("[rid]", c.Response().Header().Get(echo.HeaderXRequestID))
	logger.Error(ridMessage, statusMessge, err.Error(), inner.Error())
	res := errorResponse{Message: message, Error: inner.Error()}
	c.JSON(code, res)
}
