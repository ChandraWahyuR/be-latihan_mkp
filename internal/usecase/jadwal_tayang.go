package usecase

import (
	"context"

	errResp "github.com/ChandraWahyuR/be-latihan_mkp/constant/error"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/entity"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/model"
	"github.com/google/uuid"
)

type JadwalTayangRepositoryInterface interface {
	CreateJadwalTayang(ctx context.Context, m *entity.JadwalTayang) error
	GetJadwalTayang(ctx context.Context) ([]entity.JadwalTayang, error)
	GetJadwalTayangByID(ctx context.Context, id string) (*entity.JadwalTayang, error)
	EditJadwalTayang(ctx context.Context, m *entity.JadwalTayang) error
	DeleteJadwalTayang(ctx context.Context, id string) error
}

type JadwalTayangUseCase struct {
	jtri JadwalTayangRepositoryInterface
	mri  MoviesRepositoryInterface
}

func NewJadwalTayangUsecase(jtri JadwalTayangRepositoryInterface, mri MoviesRepositoryInterface) *JadwalTayangUseCase {
	return &JadwalTayangUseCase{
		jtri: jtri,
		mri:  mri,
	}
}

func (s *JadwalTayangUseCase) CreateJadwalTayang(ctx context.Context, req *model.CreateJadwalTayang) error {
	if req.MovieID == "" {
		return errResp.ErrIDEmpty
	}
	if req.NameStudio == "" {
		return errResp.ErrFieldEmpty
	}

	if req.Starting.IsZero() || req.Ending.IsZero() {
		return errResp.ErrFieldEmpty
	}
	if req.Starting.After(*req.Ending) {
		return errResp.ErrInvalidTime
	}

	_, err := s.mri.GetMovieByID(ctx, req.MovieID)
	if err != nil {
		return err
	}

	jt := &entity.JadwalTayang{
		ID:         uuid.New().String(),
		MovieID:    req.MovieID,
		NameStudio: req.NameStudio,
		Starting:   req.Starting,
		Ending:     req.Ending,
	}

	err = s.jtri.CreateJadwalTayang(ctx, jt)
	if err != nil {
		return err
	}

	return nil
}

func (s *JadwalTayangUseCase) GetJadwalTayang(ctx context.Context) ([]model.GetJadwalTayang, error) {
	entities, err := s.jtri.GetJadwalTayang(ctx)
	if err != nil {
		return nil, err
	}

	resp := make([]model.GetJadwalTayang, 0)
	for _, v := range entities {
		resp = append(resp, model.GetJadwalTayang{
			ID:         v.ID,
			MovieID:    v.MovieID,
			NameStudio: v.NameStudio,
			Starting:   v.Starting,
			Ending:     v.Ending,
		})
	}

	return resp, nil
}

func (s *JadwalTayangUseCase) GetJadwalTayangByID(ctx context.Context, id string) (*model.GetJadwalTayang, error) {
	if id == "" {
		return nil, errResp.ErrIDEmpty
	}

	entity, err := s.jtri.GetJadwalTayangByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.GetJadwalTayang{
		ID:         entity.ID,
		MovieID:    entity.MovieID,
		NameStudio: entity.NameStudio,
		Starting:   entity.Starting,
		Ending:     entity.Ending,
	}, nil
}

func (s *JadwalTayangUseCase) EditJadwalTayang(ctx context.Context, req *model.EditJadwalTayang) error {
	if req.ID == "" {
		return errResp.ErrIDEmpty
	}

	d, err := s.jtri.GetJadwalTayangByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if req.MovieID != "" {
		d.MovieID = req.MovieID
	}
	if req.NameStudio != "" {
		d.NameStudio = req.NameStudio
	}

	if req.Starting != nil && !req.Starting.IsZero() {
		d.Starting = req.Starting
	}

	if req.Ending != nil && !req.Ending.IsZero() {
		d.Ending = req.Ending
	}

	if d.Starting.After(*d.Ending) {
		return errResp.ErrInvalidTime
	}

	err = s.jtri.EditJadwalTayang(ctx, d)
	if err != nil {
		return err
	}

	return nil
}

func (s *JadwalTayangUseCase) DeleteJadwalTayang(ctx context.Context, id string) error {
	if id == "" {
		return errResp.ErrIDEmpty
	}

	_, err := s.jtri.GetJadwalTayangByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.jtri.DeleteJadwalTayang(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
