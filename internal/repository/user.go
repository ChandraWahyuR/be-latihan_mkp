package repository

import (
	"context"
	"database/sql"

	"github.com/ChandraWahyuR/be-latihan_mkp/common/util"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/entity"

	"github.com/sirupsen/logrus"
)

type UserRepo struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewUserRepository(db *sql.DB, log *logrus.Logger) *UserRepo {
	return &UserRepo{
		db:  db,
		log: log,
	}
}

func (r *UserRepo) Register(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return util.ParsePQError(err)
	}

	return nil
}

func (r *UserRepo) Login(ctx context.Context, email string) (*entity.User, error) {
	var u entity.User
	query := `SELECT id, name, email FROM users WHERE email = $1 AND deleted_at IS NULL`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, util.ParsePQError(err)
	}

	return &u, nil
}

func (r *UserRepo) GetDataFromEmail(ctx context.Context, email string) (*entity.User, error) {
	var u entity.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1 AND deleted_at IS NULL`
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return nil, util.ParsePQError(err)
	}

	return &u, nil
}

func (r *UserRepo) IsDataAvailable(ctx context.Context, email, name string) bool {
	var data string
	query := `SELECT email, name FROM users WHERE email = $1 OR username = $2`
	err := r.db.QueryRowContext(ctx, query, email, name).Scan(&data)
	return err != nil
}
