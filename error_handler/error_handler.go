package errorHandler

import (
	"fmt"
	"net/http"

	"github.com/nuts300/test-echo/app_error"

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
	innerMessage := ""
	var inner error
	message := err.Error()
	if ae, ok := err.(*appError.AppError); ok {
		httpError := appError.ConvertToHttpError(ae)
		code = httpError.Code
		inner = httpError.Inner
		message = httpError.Message.(string)
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		inner = he.Inner
		message = he.Message.(string)
	}
	if inner != nil {
		innerMessage = inner.Error()
	}

	statusMessge := fmt.Sprint("[status]", code)
	ridMessage := fmt.Sprint("[rid]", c.Response().Header().Get(echo.HeaderXRequestID))
	logger.Error(ridMessage, statusMessge, message, innerMessage)
	res := errorResponse{Message: message, Error: innerMessage}
	c.JSON(code, res)
}
