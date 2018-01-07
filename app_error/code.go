package appError

import (
	"net/http"
)

type ErrorCode struct {
	Code    int
	Message string
}

var (
	ErrorInternalServer     = ErrorCode{Code: http.StatusInternalServerError, Message: "Unexpected error."}
	ErrorInvalidUserID      = ErrorCode{Code: http.StatusBadRequest, Message: "Invalid user id."}
	ErrorInvalidUserPayload = ErrorCode{Code: http.StatusBadRequest, Message: "Invalid user payload."}
	ErrorFailedReadUser     = ErrorCode{Code: http.StatusInternalServerError, Message: "Failed read user."}
	ErrorFailedReadUsers    = ErrorCode{Code: http.StatusInternalServerError, Message: "Failed read users."}
	ErrorFailedCreateUser   = ErrorCode{Code: http.StatusInternalServerError, Message: "Failed create user."}
	ErrorFailedUpdateUser   = ErrorCode{Code: http.StatusInternalServerError, Message: "Failed update user."}
	ErrorFailedDeleteUser   = ErrorCode{Code: http.StatusInternalServerError, Message: "Failed delete user."}
	ErrorNotFoundUser       = ErrorCode{Code: http.StatusNotFound, Message: "Not found user."}
	ErrorUnauthorized       = ErrorCode{Code: http.StatusInternalServerError, Message: "UnAuthorized"}
)
