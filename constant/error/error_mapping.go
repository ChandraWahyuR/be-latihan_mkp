package error

import (
	"errors"
	"net/http"

	errUsr "github.com/ChandraWahyuR/be-latihan_mkp/constant/error/user"
)

func ConvertErrorToCode(err error) int {
	// 400
	if errors.Is(err, errUsr.ErrFieldEmailEmpty) ||
		errors.Is(err, errUsr.ErrFieldNameEmpty) ||
		errors.Is(err, errUsr.ErrFieldPasswordEmpty) ||
		errors.Is(err, errUsr.ErrFormatEmail) ||
		errors.Is(err, errUsr.ErrPasswordDoesNotMatch) {
		return http.StatusBadRequest
	}

	// 401
	if errors.Is(err, errUsr.ErrPasswordInCorrect) {
		return http.StatusUnauthorized
	}

	// 404
	if errors.Is(err, errUsr.ErrUserNotFound) {
		return http.StatusNotFound
	}

	// 409
	if errors.Is(err, errUsr.ErrEmailExist) {
		return http.StatusConflict
	}

	return http.StatusInternalServerError
}
