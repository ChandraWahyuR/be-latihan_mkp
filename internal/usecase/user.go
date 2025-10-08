package usecase

import (
	"context"
	"strings"

	"github.com/ChandraWahyuR/be-latihan_mkp/common/util"
	errUsr "github.com/ChandraWahyuR/be-latihan_mkp/constant/error/user"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/auth/jwt"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/entity"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UsersRepositoryInterface interface {
	Register(ctx context.Context, user *entity.User) error
	Login(ctx context.Context, email string) (*entity.User, error)
	GetDataFromEmail(ctx context.Context, email string) (*entity.User, error)
	IsDataAvailable(ctx context.Context, email, name string) bool
}

type UserUsecase struct {
	log *logrus.Logger
	ur  UsersRepositoryInterface
	jwt jwt.JWTService
}

func NewUserUsecase(log *logrus.Logger, ur UsersRepositoryInterface, jwt jwt.JWTService) *UserUsecase {
	return &UserUsecase{
		log: log,
		ur:  ur,
		jwt: jwt,
	}
}

func (s *UserUsecase) Register(ctx context.Context, user *model.Register) error {
	switch {
	case user.Email == "":
		return errUsr.ErrFieldEmailEmpty
	case user.Name == "":
		return errUsr.ErrFieldNameEmpty
	case user.Password == "":
		return errUsr.ErrFieldPasswordEmpty
	case user.Password != user.ConfirmPassword:
		return errUsr.ErrPasswordDoesNotMatch
	}

	if !util.ValidasiEmail(user.Email) {
		return errUsr.ErrFormatEmail
	}

	if !s.ur.IsDataAvailable(ctx, user.Email, user.Name) {
		return errUsr.ErrEmailExist
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	resData := entity.User{
		ID:       uuid.New().String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	err = s.ur.Register(ctx, &resData)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserUsecase) Login(ctx context.Context, email, password string) (*string, error) {
	switch {
	case email == "":
		return nil, errUsr.ErrFieldEmailEmpty
	case password == "":
		return nil, errUsr.ErrFieldPasswordEmpty
	}

	if !util.ValidasiEmail(email) {
		return nil, errUsr.ErrFormatEmail
	}

	email = strings.ToLower(email)

	usr, err := s.ur.GetDataFromEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !util.VerifyHashedPassword(password, usr.Password) {
		return nil, errUsr.ErrPasswordInCorrect
	}

	res, err := s.ur.Login(ctx, email)
	if err != nil {
		return nil, err
	}

	tokenData := model.Login{
		ID:    res.ID,
		Email: res.Email,
	}

	token, err := s.jwt.GenerateToken(&tokenData)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
