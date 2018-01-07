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
	innerMessage := ""
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		inner = he.Inner
		message = he.Message.(string)
	}
	if inner != nil {
		innerMessage = inner.Error()
	}
	statusMessge := fmt.Sprint("[status]", code)
	ridMessage := fmt.Sprint("[rid]", c.Response().Header().Get(echo.HeaderXRequestID))
	logger.Error(ridMessage, statusMessge, err.Error(), innerMessage)
	res := errorResponse{Message: message, Error: innerMessage}
	c.JSON(code, res)
}
