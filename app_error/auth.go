package appError

import (
	"net/http"

	"github.com/labstack/echo"
)

func NewErrorUnAuthorized(err error) *AppError {
	return newAppError(codeUnAuthorized.ID, err)
}

func convertToAuthHttpError(appError *AppError) *echo.HTTPError {
	switch appError.ErrorCode {
	case codeUnAuthorized.ID:
		return newHttpError(http.StatusNotFound, codeUnAuthorized.Message, appError)
	default:
		return nil
	}
}
