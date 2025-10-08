package error

import (
	"errors"
	"net/http"

	errUsr "github.com/ChandraWahyuR/be-latihan_mkp/constant/error/user"
)

func ConvertErrorToCode(err error) int {
	switch {
	// 400
	case errors.Is(err, ErrFieldEmpty) ||
		errors.Is(err, ErrIDEmpty) ||
		errors.Is(err, ErrInvalidTime) ||
		errors.Is(err, errUsr.ErrFieldEmailEmpty) ||
		errors.Is(err, errUsr.ErrFieldNameEmpty) ||
		errors.Is(err, errUsr.ErrFieldPasswordEmpty) ||
		errors.Is(err, errUsr.ErrFormatEmail) ||
		errors.Is(err, errUsr.ErrPasswordDoesNotMatch):
		return http.StatusBadRequest

	// 401
	case errors.Is(err, ErrUnauthorized) ||
		errors.Is(err, ErrInvalidToken) ||
		errors.Is(err, errUsr.ErrPasswordInCorrect):
		return http.StatusUnauthorized

	// 403
	case errors.Is(err, ErrForbidden):
		return http.StatusForbidden

	// 404
	case errors.Is(err, ErrNotFound) ||
		errors.Is(err, errUsr.ErrUserNotFound):
		return http.StatusNotFound

	// 409
	case errors.Is(err, errUsr.ErrEmailExist):
		return http.StatusConflict

	default:
		return http.StatusInternalServerError
	}
}
