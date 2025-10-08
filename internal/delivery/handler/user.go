package handler

import (
	"context"

	"net/http"

	"github.com/ChandraWahyuR/be-latihan_mkp/common/response"
	"github.com/ChandraWahyuR/be-latihan_mkp/constant"
	errMap "github.com/ChandraWahyuR/be-latihan_mkp/constant/error"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/auth/jwt"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/model"
	"github.com/gin-gonic/gin"
)

type RegisterUsecaseInterface interface {
	Register(ctx context.Context, user *model.Register) error
	Login(ctx context.Context, email, password string) (*string, error)
}

type UserHandler struct {
	jwt jwt.JWTService
	uc  RegisterUsecaseInterface
}

func NewUserHandler(jwt jwt.JWTService, uc RegisterUsecaseInterface) *UserHandler {
	return &UserHandler{
		jwt: jwt,
		uc:  uc,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	var data model.Register
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseHandler(constant.Error, "error memproses data", nil))
		return
	}

	err = h.uc.Register(ctx, &data)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, response.ResponseHandler(constant.Success, "Berhasil membuat akun", nil))
}

func (h *UserHandler) Login(c *gin.Context) {
	var data model.Login
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseHandler(constant.Error, "error memproses data", nil))
		return
	}

	ctx := c.Request.Context()
	result, err := h.uc.Login(ctx, data.Email, data.Password)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	res := map[string]string{
		"token": *result,
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Login berhasil", res))
}
