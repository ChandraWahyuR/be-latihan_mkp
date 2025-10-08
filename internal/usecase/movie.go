package usecase

import (
	"context"

	"github.com/ChandraWahyuR/be-latihan_mkp/internal/entity"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/model"
	"github.com/google/uuid"

	errResp "github.com/ChandraWahyuR/be-latihan_mkp/constant/error"
)

type MoviesRepositoryInterface interface {
	CreateMovie(ctx context.Context, m *entity.Movie) error
	GetMovie(ctx context.Context) ([]entity.Movie, error)
	GetMovieByID(ctx context.Context, id string) (*entity.Movie, error)
	GetMovieAndStudio(ctx context.Context, id string) (*entity.MovieAndStudio, error)
	EditMovie(ctx context.Context, m *entity.Movie) error
	DeleteMovie(ctx context.Context, id string) error
}

type MovieUseCase struct {
	mri MoviesRepositoryInterface
}

func NewMoviesUsecase(mri MoviesRepositoryInterface) *MovieUseCase {
	return &MovieUseCase{
		mri: mri,
	}
}

func (s *MovieUseCase) CreateMovie(ctx context.Context, req *model.CreateMovie) error {
	if req.Title == "" {
		return errResp.ErrFieldEmpty
	}

	m := &entity.Movie{
		ID:    uuid.New().String(),
		Title: req.Title,
	}

	err := s.mri.CreateMovie(ctx, m)
	if err != nil {
		return err
	}

	return nil
}

func (s *MovieUseCase) GetMovie(ctx context.Context) ([]model.GetMovieResponse, error) {
	entity, err := s.mri.GetMovie(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]model.GetMovieResponse, 0)
	for _, v := range entity {
		resp = append(resp, model.GetMovieResponse{
			ID:    v.ID,
			Title: v.Title,
		})
	}

	return resp, nil
}

func (s *MovieUseCase) GetMovieDetail(ctx context.Context, id string) (*model.MovieAndStudio, error) {
	if id == "" {
		return nil, errResp.ErrIDEmpty
	}

	entity, err := s.mri.GetMovieAndStudio(ctx, id)
	if err != nil {
		return nil, err
	}

	jt := make([]model.JamTayang, 0)
	for _, v := range entity.JadwalTayang {
		jt = append(jt, model.JamTayang{
			MovieID:    entity.ID,
			NameStudio: v.NameStudio,
			Starting:   v.Starting,
			Ending:     v.Ending,
		})
	}

	return &model.MovieAndStudio{
		ID:        entity.ID,
		Title:     entity.Title,
		JamTayang: jt,
	}, nil
}

func (s *MovieUseCase) EditMovie(ctx context.Context, req *model.EditMovie) error {
	if req.ID == "" {
		return errResp.ErrIDEmpty
	}

	d, err := s.mri.GetMovieByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if req.Title != "" {
		d.Title = req.Title
	}

	err = s.mri.EditMovie(ctx, d)
	if err != nil {
		return err
	}
	return nil
}

func (s *MovieUseCase) DeleteMovie(ctx context.Context, id string) error {
	if id == "" {
		return errResp.ErrIDEmpty
	}

	_, err := s.mri.GetMovieByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.mri.DeleteMovie(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
