package appError

type ErrorCode struct {
	Code    int
	Message string
}

var (
	INTERNAL_ERROR       = ErrorCode{Code: 500, Message: "Unexpected error"}
	INVALID_USER_ID      = ErrorCode{Code: 400, Message: "Invalid user id"}
	INVALID_USER_PAYLOAD = ErrorCode{Code: 400, Message: "Invalid user payload"}
)
