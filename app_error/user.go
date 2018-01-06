package appError

import "github.com/labstack/echo"
import "net/http"

func NewErrorInvalidUserID(err error) *AppError {
	return newAppError(codeInvalidUserID.ID, err)
}

func NewErrorInvalidUserPayload(err error) *AppError {
	return newAppError(codeInvalidUserPayload.ID, err)
}

func NewErrorFailedReadUser(err error) *AppError {
	return newAppError(codeFailedReadUser.ID, err)
}

func NewErrorFailedReadUsers(err error) *AppError {
	return newAppError(codeFailedReadUsers.ID, err)
}

func NewErrorFailedCreateUser(err error) *AppError {
	return newAppError(codeFailedCreateUser.ID, err)
}

func NewErrorFailedUpdateUser(err error) *AppError {
	return newAppError(codeFailedUpdateUser.ID, err)
}

func NewErrorFailedDeleteUser(err error) *AppError {
	return newAppError(codeFailedDeleteUser.ID, err)
}

func NewErrorNotFoundUser(err error) *AppError {
	return newAppError(codeNotFoundUser.ID, err)
}

func convertToUserHttpError(appError *AppError) *echo.HTTPError {
	switch appError.ErrorCode {
	case codeFailedCreateUser.ID:
		return newHttpError(http.StatusInternalServerError, codeFailedCreateUser.Message, appError)
	case codeFailedDeleteUser.ID:
		return newHttpError(http.StatusInternalServerError, codeFailedDeleteUser.Message, appError)
	case codeFailedReadUser.ID:
		return newHttpError(http.StatusInternalServerError, codeFailedReadUser.Message, appError)
	case codeFailedReadUsers.ID:
		return newHttpError(http.StatusInternalServerError, codeFailedReadUsers.Message, appError)
	case codeFailedUpdateUser.ID:
		return newHttpError(http.StatusInternalServerError, codeFailedUpdateUser.Message, appError)
	case codeInvalidUserID.ID:
		return newHttpError(http.StatusBadRequest, codeInvalidUserID.Message, appError)
	case codeInvalidUserPayload.ID:
		return newHttpError(http.StatusBadRequest, codeInvalidUserPayload.Message, appError)
	case codeNotFoundUser.ID:
		return newHttpError(http.StatusNotFound, codeNotFoundUser.Message, appError)
	default:
		return nil
	}
}
