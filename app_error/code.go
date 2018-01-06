package appError

type errorCode struct {
	ID string
}

const (
	codeInternalServer     = "INTERNAL_SERVER_ERROR"
	codeInvalidUserID      = "INVALID_USER_ID"
	codeInvalidUserPayload = "INVALID_USER_PAYLOAD"
	codeFailedReadUser     = "FAILD_READ_USER"
	codeFailedReadUsers    = "FAILED_READ_USERS"
	codeFailedCreateUser   = "FAILED_CREATE_USER"
	codeFailedUpdateUser   = "FAILED_UPDATE_USER"
	codeFailedDeleteUser   = "FAILED_DELETE_USER"
	codeNotFoundUser       = "NOT_FOUND_USER"
	codeUnAuthorized       = "UNAUTHORIZED"
)

const (
	messageInternalServer     = "Unexpected error."
	messageInvalidUserID      = "Invalid user id."
	messageInvalidUserPayload = "Invalid user payload."
	messageFailedReadUser     = "Failed read user."
	messageFailedReadUsers    = "Failed read users."
	messageFailedCreateUser   = "Failed create user."
	messageFailedUpdateUser   = "Failed update user."
	messageFailedDeleteUser   = "Failed delete user."
	messageNotFoundUser       = "Not found user."
	messageUnAuthorized       = "UnAuthorized."
)
