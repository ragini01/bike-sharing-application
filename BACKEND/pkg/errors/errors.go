package errors

import "fmt"

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s (code %d)", e.Message, e.Code)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

var (
	ErrInvalidRequest = NewError(400, "Invalid request")
	ErrUnauthorized   = NewError(401, "Unauthorized")
	ErrNotFound       = NewError(404, "Not found")
	ErrInternalServer = NewError(500, "Internal server error")
)
