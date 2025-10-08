package user

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordInCorrect    = errors.New("password incorrect")
	ErrEmailExist           = errors.New("email already exist")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
	ErrFieldEmailEmpty      = errors.New("field email is empty")
	ErrFieldNameEmpty       = errors.New("field name is empty")
	ErrFieldPasswordEmpty   = errors.New("field password is empty")
	ErrFormatEmail          = errors.New("format email is not correct")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrPasswordDoesNotMatch,
}
