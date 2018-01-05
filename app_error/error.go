package appError

import (
	"github.com/labstack/echo"
)

func NewAppError(errorCode ErrorCode, err error) *echo.HTTPError {
	if err != nil {
		switch err.(type) {
		case *echo.HTTPError:
			return err.(*echo.HTTPError)
		}
	}
	return &echo.HTTPError{
		Code:    errorCode.Code,
		Message: errorCode.Message,
		Inner:   err,
	}
}
