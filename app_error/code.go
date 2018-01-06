package appError

type ErrorCode struct {
	Code    int
	Message string
}

var (
	INTERNAL_ERROR       = ErrorCode{Code: 500, Message: "Unexpected error."}
	INVALID_USER_ID      = ErrorCode{Code: 400, Message: "Invalid user id."}
	INVALID_USER_PAYLOAD = ErrorCode{Code: 400, Message: "Invalid user payload."}
	FAILED_READ_USER     = ErrorCode{Code: 500, Message: "Failed read user."}
	FAILED_READ_USERS    = ErrorCode{Code: 500, Message: "Failed read users."}
	FAILED_CREATE_USER   = ErrorCode{Code: 500, Message: "Failed create user."}
	FAILED_UPDATE_USER   = ErrorCode{Code: 500, Message: "Failed update user."}
	FAILED_DELETE_USER   = ErrorCode{Code: 500, Message: "Failed delete user."}
	NOT_FOUND_USER       = ErrorCode{Code: 404, Message: "Not found user."}
	UNAUTHORIZED_ERROR   = ErrorCode{Code: 500, Message: "UnAuthorized"}
)
