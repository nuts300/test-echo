package appError

type errorCode struct {
	ID      string
	Message string
}

var (
	codeInternalServer     = errorCode{ID: "INTERNAL_SERVER_ERROR", Message: "Unexpected error."}
	codeInvalidUserID      = errorCode{ID: "INVALID_USER_ID", Message: "Invalid user id."}
	codeInvalidUserPayload = errorCode{ID: "INVALID_USER_PAYLOAD", Message: "Invalid user payload."}
	codeFailedReadUser     = errorCode{ID: "FAILD_READ_USER", Message: "Failed read user."}
	codeFailedReadUsers    = errorCode{ID: "FAILED_READ_USERS", Message: "Failed read users."}
	codeFailedCreateUser   = errorCode{ID: "FAILED_CREATE_USER", Message: "Failed create user."}
	codeFailedUpdateUser   = errorCode{ID: "FAILED_UPDATE_USER", Message: "Failed update user."}
	codeFailedDeleteUser   = errorCode{ID: "FAILED_DELETE_USER", Message: "Failed delete user."}
	codeNotFoundUser       = errorCode{ID: "NOT_FOUND_USER", Message: "Not found user."}
	codeUnAuthorized       = errorCode{ID: "UNAUTHORIZED", Message: "UnAuthorized."}
)
