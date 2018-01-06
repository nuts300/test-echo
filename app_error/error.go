package appError

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func NewHttpError(status int, message string, err error) *echo.HTTPError {
	return &echo.HTTPError{
		Code:    status,
		Message: message,
		Inner:   err,
	}
}

type AppError struct {
	ErrorCode string
	Inner     error
}

func (ae *AppError) Error() string {
	return fmt.Sprint("[", ae.ErrorCode, "]", ae.Inner.Error())
}

func NewAppError(errorCode string, err error) *AppError {
	return &AppError{ErrorCode: errorCode, Inner: err}
}

func ErrorInternalServer(err error) *AppError {
	return NewAppError(CodeInternalServer, err)
}

func ConvertToHttpError(appError *AppError) *echo.HTTPError {
	if httpError := ConvertToAuthHttpError(appError); httpError != nil {
		return httpError
	}
	if httpError := ConvertToUserHttpError(appError); httpError != nil {
		return httpError
	}
	return NewHttpError(http.StatusInternalServerError, MessageInternalServer, appError)
}
