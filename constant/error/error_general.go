package error

import "errors"

const (
	Success = "success"
	Error   = "error"
)

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError            = errors.New("database server failed to execute query")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrTooManyReq          = errors.New("database server failded to execute query")
	ErrInvalidToken        = errors.New("invalid token")
	ErrForbidden           = errors.New("forbidden")
	ErrInvalidUploadFile   = errors.New("invalid upload file")
	ErrSizeTooBig          = errors.New("size too big")
	ErrNotFound            = errors.New("not found")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyReq,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
	ErrNotFound,
}
