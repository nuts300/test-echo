package appError

type errorCode struct {
	ID string
}

const (
	CodeInternalServer     = "INTERNAL_SERVER_ERROR"
	CodeInvalidUserID      = "INVALID_USER_ID"
	CodeInvalidUserPayload = "INVALID_USER_PAYLOAD"
	CodeFailedReadUser     = "FAILD_READ_USER"
	CodeFailedReadUsers    = "FAILED_READ_USERS"
	CodeFailedCreateUser   = "FAILED_CREATE_USER"
	CodeFailedUpdateUser   = "FAILED_UPDATE_USER"
	CodeFailedDeleteUser   = "FAILED_DELETE_USER"
	CodeNotFoundUser       = "NOT_FOUND_USER"
	CodeUnAuthorized       = "UNAUTHORIZED"
)

const (
	MessageInternalServer     = "Unexpected error."
	MessageInvalidUserID      = "Invalid user id."
	MessageInvalidUserPayload = "Invalid user payload."
	MessageFailedReadUser     = "Failed read user."
	MessageFailedReadUsers    = "Failed read users."
	MessageFailedCreateUser   = "Failed create user."
	MessageFailedUpdateUser   = "Failed update user."
	MessageFailedDeleteUser   = "Failed delete user."
	MessageNotFoundUser       = "Not found user."
	MessageUnAuthorized       = "UnAuthorized."
)
