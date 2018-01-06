package appError

import "github.com/labstack/echo"
import "net/http"

func NewErrorInvalidUserID(err error) *AppError {
	return newAppError(codeInvalidUserID, err)
}

func NewErrorInvalidUserPayload(err error) *AppError {
	return newAppError(codeInvalidUserPayload, err)
}

func NewErrorFailedReadUser(err error) *AppError {
	return newAppError(codeFailedReadUser, err)
}

func NewErrorFailedReadUsers(err error) *AppError {
	return newAppError(codeFailedReadUsers, err)
}

func NewErrorFailedCreateUser(err error) *AppError {
	return newAppError(codeFailedCreateUser, err)
}

func NewErrorFailedUpdateUser(err error) *AppError {
	return newAppError(codeFailedUpdateUser, err)
}

func NewErrorFailedDeleteUser(err error) *AppError {
	return newAppError(codeFailedDeleteUser, err)
}

func NewErrorNotFoundUser(err error) *AppError {
	return newAppError(codeNotFoundUser, err)
}

func convertToUserHttpError(appError *AppError) *echo.HTTPError {
	switch appError.ErrorCode {
	case codeFailedCreateUser:
		return newHttpError(http.StatusInternalServerError, messageFailedCreateUser, appError)
	case codeFailedDeleteUser:
		return newHttpError(http.StatusInternalServerError, messageFailedDeleteUser, appError)
	case codeFailedReadUser:
		return newHttpError(http.StatusInternalServerError, messageFailedReadUser, appError)
	case codeFailedReadUsers:
		return newHttpError(http.StatusInternalServerError, messageFailedReadUser, appError)
	case codeFailedUpdateUser:
		return newHttpError(http.StatusInternalServerError, messageFailedUpdateUser, appError)
	case codeInvalidUserID:
		return newHttpError(http.StatusBadRequest, messageInvalidUserID, appError)
	case codeInvalidUserPayload:
		return newHttpError(http.StatusBadRequest, messageInvalidUserID, appError)
	case codeNotFoundUser:
		return newHttpError(http.StatusNotFound, messageNotFoundUser, appError)
	default:
		return nil
	}
}
