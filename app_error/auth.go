package appError

import (
	"net/http"

	"github.com/labstack/echo"
)

func NewErrorUnAuthorized(err error) *AppError {
	return newAppError(codeUnAuthorized, err)
}

func convertToAuthHttpError(appError *AppError) *echo.HTTPError {
	switch appError.ErrorCode {
	case codeUnAuthorized:
		return newHttpError(http.StatusNotFound, messageNotFoundUser, appError)
	default:
		return nil
	}
}
