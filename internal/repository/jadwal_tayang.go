package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ChandraWahyuR/be-latihan_mkp/common/util"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/entity"
)

type JadwalTayangRepo struct {
	db *sql.DB
}

func NewJadwalTayangRepository(db *sql.DB) *JadwalTayangRepo {
	return &JadwalTayangRepo{
		db: db,
	}
}

func (r *JadwalTayangRepo) CreateJadwalTayang(ctx context.Context, m *entity.JadwalTayang) error {
	query := `INSERT INTO studio_schedules (id, movie_id, name_studio, starting, ending) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.ExecContext(ctx, query, m.ID, m.MovieID, m.NameStudio, m.Starting, m.Ending)
	if err != nil {
		return util.ParsePQError(err)
	}

	return nil
}

func (r *JadwalTayangRepo) GetJadwalTayang(ctx context.Context) ([]entity.JadwalTayang, error) {
	query := `SELECT 
		id, 
		movie_id, 
		name_studio, 
		starting, 
		ending 
		FROM studio_schedules 
		WHERE deleted_at IS NULL`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, util.ParsePQError(err)
	}
	defer rows.Close()

	var ms []entity.JadwalTayang
	for rows.Next() {
		var m entity.JadwalTayang
		if err := rows.Scan(&m.ID, &m.MovieID, &m.NameStudio, &m.Starting, &m.Ending); err != nil {
			return nil, util.ParsePQError(err)
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, util.ParsePQError(err)
	}

	return ms, nil
}

func (r *JadwalTayangRepo) GetJadwalTayangByID(ctx context.Context, id string) (*entity.JadwalTayang, error) {
	query := `SELECT 
		id, 
		movie_id, 
		name_studio, 
		starting, 
		ending 
		FROM studio_schedules  
		WHERE id = $1 AND deleted_at IS NULL`
	var m entity.JadwalTayang
	err := r.db.QueryRowContext(ctx, query, id).Scan(&m.ID, &m.MovieID, &m.NameStudio, &m.Starting, &m.Ending)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.ParsePQError(err)
		}
		return nil, util.ParsePQError(err)
	}

	return &m, nil
}

func (r *JadwalTayangRepo) EditJadwalTayang(ctx context.Context, m *entity.JadwalTayang) error {
	query := `UPDATE studio_schedules SET movie_id = $1, name_studio = $2, starting = $3, ending = $4 WHERE id = $5 AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, m.MovieID, m.NameStudio, m.Starting, m.Ending, m.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return util.ParsePQError(err)
		}
		return util.ParsePQError(err)
	}

	return nil
}

func (r *JadwalTayangRepo) DeleteJadwalTayang(ctx context.Context, id string) error {
	query := `DELETE FROM studio_schedules WHERE id = $1 AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return util.ParsePQError(err)
	}

	return nil
}
