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

type MovieUsecaseInterface interface {
	CreateMovie(ctx context.Context, req *model.CreateMovie) error
	GetMovie(ctx context.Context) ([]model.GetMovieResponse, error)
	GetMovieDetail(ctx context.Context, id string) (*model.MovieAndStudio, error)
	EditMovie(ctx context.Context, req *model.EditMovie) error
	DeleteMovie(ctx context.Context, id string) error
}

type MovieHandler struct {
	uc MovieUsecaseInterface
}

func NewMovieHandler(uc MovieUsecaseInterface) *MovieHandler {
	return &MovieHandler{
		uc: uc,
	}
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	var data model.CreateMovie
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseHandler(constant.Error, "error memproses data", nil))
		return
	}

	err := h.uc.CreateMovie(c.Request.Context(), &data)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, response.ResponseHandler(constant.Success, "Berhasil menambahkan film", nil))
}

func (h *MovieHandler) GetAllMovies(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	result, err := h.uc.GetMovie(c.Request.Context())
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Berhasil mendapatkan semua film", result))
}

func (h *MovieHandler) GetMovieDetailByID(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	id := c.Param("id")
	result, err := h.uc.GetMovieDetail(c.Request.Context(), id)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Berhasil mendapatkan detail film", result))
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	id := c.Param("id")
	var data model.EditMovie
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ResponseHandler(constant.Error, "error memproses data", nil))
		return
	}

	data.ID = id

	err := h.uc.EditMovie(c.Request.Context(), &data)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Berhasil memperbarui film", nil))
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	_, ok := middleware.GetUser(c)
	if !ok {
		middleware.HandleUnauthorizedError(c)
		return
	}

	id := c.Param("id")
	err := h.uc.DeleteMovie(c.Request.Context(), id)
	if err != nil {
		c.JSON(errMap.ConvertErrorToCode(err), response.ResponseHandler(constant.Error, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.ResponseHandler(constant.Success, "Berhasil menghapus film", nil))
}
