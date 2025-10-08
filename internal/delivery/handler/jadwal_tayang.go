package handler

import (
	"context"
	"net/http"

	"github.com/ChandraWahyuR/be-latihan_mkp/common/response"
	"github.com/ChandraWahyuR/be-latihan_mkp/constant"
	errMap "github.com/ChandraWahyuR/be-latihan_mkp/constant/error"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/delivery/middleware"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/model"
	"github.com/gin-gonic/gin"
)

type JadwalTayangUsecaseInterface interface {
	CreateJadwalTayang(ctx context.Context, req *model.CreateJadwalTayang) error
	GetJadwalTayang(ctx context.Context) ([]model.GetJadwalTayang, error)
	GetJadwalTayangByID(ctx context.Context, id string) (*model.GetJadwalTayang, error)
	EditJadwalTayang(ctx context.Context, req *model.EditJadwalTayang) error
	DeleteJadwalTayang(ctx context.Context, id string) error
}

type JadwalTayangHandler struct {
	uc JadwalTayangUsecaseInterface
}

func NewJadwalTayangHandler(uc JadwalTayangUsecaseInterface) *JadwalTayangHandler {
	return &JadwalTayangHandler{
		uc: uc,
	}
}

func (h *JadwalTayangHandler) CreateJadwalTayang(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	var data model.CreateJadwalTayang
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseHandler(constant.Error, "error memproses data", nil))
		return
	}

	err := h.uc.CreateJadwalTayang(c.Request.Context(), &data)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, response.ResponseHandler(constant.Success, "Berhasil menambahkan jadwal tayang", nil))
}

func (h *JadwalTayangHandler) GetAllJadwalTayang(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	result, err := h.uc.GetJadwalTayang(c.Request.Context())
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Berhasil mendapatkan semua jadwal tayang", result))
}

func (h *JadwalTayangHandler) GetJadwalTayangByID(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	id := c.Param("id")
	result, err := h.uc.GetJadwalTayangByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Berhasil mendapatkan detail jadwal tayang", result))
}

func (h *JadwalTayangHandler) UpdateJadwalTayang(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	id := c.Param("id")
	var data model.EditJadwalTayang
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseHandler(constant.Error, "error memproses data", nil))
		return
	}
	data.ID = id

	err := h.uc.EditJadwalTayang(c.Request.Context(), &data)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Berhasil memperbarui jadwal tayang", nil))
}

func (h *JadwalTayangHandler) DeleteJadwalTayang(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	id := c.Param("id")
	err := h.uc.DeleteJadwalTayang(c.Request.Context(), id)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Berhasil menghapus jadwal tayang", nil))
}
