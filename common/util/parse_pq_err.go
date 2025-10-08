package util

import (
	"database/sql"
	"errors"

	errG "github.com/ChandraWahyuR/be-latihan_mkp/constant/error"
	errUsr "github.com/ChandraWahyuR/be-latihan_mkp/constant/error/user"
	"github.com/lib/pq"
)

func ParsePQError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return errG.ErrNotFound
	}
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		if pqErr.Code == "23505" {
			switch pqErr.Constraint {
			case "users_email_key":
				return errUsr.ErrEmailExist
			}
		}
	}
	return err
}
