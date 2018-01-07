package appError

type ErrorCode struct {
	Code    int
	Message string
}

var (
	ErrorInternalServer     = ErrorCode{Code: 500, Message: "Unexpected error."}
	ErrorInvalidUserID      = ErrorCode{Code: 400, Message: "Invalid user id."}
	ErrorInvalidUserPayload = ErrorCode{Code: 400, Message: "Invalid user payload."}
	ErrorFailedReadUser     = ErrorCode{Code: 500, Message: "Failed read user."}
	ErrorFailedReadUsers    = ErrorCode{Code: 500, Message: "Failed read users."}
	ErrorFailedCreateUser   = ErrorCode{Code: 500, Message: "Failed create user."}
	ErrorFailedUpdateUser   = ErrorCode{Code: 500, Message: "Failed update user."}
	ErrorFailedDeleteUser   = ErrorCode{Code: 500, Message: "Failed delete user."}
	ErrorNotFoundUser       = ErrorCode{Code: 404, Message: "Not found user."}
	ErrorUnauthorized       = ErrorCode{Code: 500, Message: "UnAuthorized"}
)
