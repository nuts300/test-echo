package appError

import "github.com/labstack/echo"
import "net/http"

func ErrorInvalidUserID(err error) *AppError {
	return NewAppError(CodeInvalidUserID, err)
}

func ErrorInvalidUserPayload(err error) *AppError {
	return NewAppError(CodeInvalidUserPayload, err)
}

func ErrorFailedReadUser(err error) *AppError {
	return NewAppError(CodeFailedReadUser, err)
}

func ErrorFailedReadUsers(err error) *AppError {
	return NewAppError(CodeFailedReadUsers, err)
}

func ErrorFailedCreateUser(err error) *AppError {
	return NewAppError(CodeFailedCreateUser, err)
}

func ErrorFailedUpdateUser(err error) *AppError {
	return NewAppError(CodeFailedUpdateUser, err)
}

func ErrorFailedDeleteUser(err error) *AppError {
	return NewAppError(CodeFailedDeleteUser, err)
}

func ErrorNotFoundUser(err error) *AppError {
	return NewAppError(CodeNotFoundUser, err)
}

func ConvertToUserHttpError(appError *AppError) *echo.HTTPError {
	switch appError.ErrorCode {
	case CodeFailedCreateUser:
		return NewHttpError(http.StatusInternalServerError, MessageFailedCreateUser, appError)
	case CodeFailedDeleteUser:
		return NewHttpError(http.StatusInternalServerError, MessageFailedDeleteUser, appError)
	case CodeFailedReadUser:
		return NewHttpError(http.StatusInternalServerError, MessageFailedReadUser, appError)
	case CodeFailedReadUsers:
		return NewHttpError(http.StatusInternalServerError, MessageFailedReadUser, appError)
	case CodeFailedUpdateUser:
		return NewHttpError(http.StatusInternalServerError, MessageFailedUpdateUser, appError)
	case CodeInvalidUserID:
		return NewHttpError(http.StatusBadRequest, MessageInvalidUserID, appError)
	case CodeInvalidUserPayload:
		return NewHttpError(http.StatusBadRequest, MessageInvalidUserID, appError)
	case CodeNotFoundUser:
		return NewHttpError(http.StatusNotFound, MessageNotFoundUser, appError)
	default:
		return nil
	}
}
