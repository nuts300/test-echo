package appError

import (
	"github.com/labstack/echo"
)

func NewAppError(errorCode ErrorCode, err error) *echo.HTTPError {
	return &echo.HTTPError{
		Code:    errorCode.Code,
		Message: errorCode.Message,
		Inner:   err,
	}
}
