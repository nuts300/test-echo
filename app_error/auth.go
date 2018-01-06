package appError

import (
	"net/http"

	"github.com/labstack/echo"
)

func ErrorUnAuthorized(err error) *AppError {
	return NewAppError(CodeUnAuthorized, err)
}

func ConvertToAuthHttpError(appError *AppError) *echo.HTTPError {
	switch appError.ErrorCode {
	case CodeUnAuthorized:
		return NewHttpError(http.StatusNotFound, MessageNotFoundUser, appError)
	default:
		return nil
	}
}
