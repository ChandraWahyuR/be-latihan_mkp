package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ChandraWahyuR/be-latihan_mkp/common/util"
	"github.com/ChandraWahyuR/be-latihan_mkp/internal/entity"
)

type MoviesRepo struct {
	db *sql.DB
}

func NewMoviesRepository(db *sql.DB) *MoviesRepo {
	return &MoviesRepo{
		db: db,
	}
}

func (r *MoviesRepo) CreateMovie(ctx context.Context, m *entity.Movie) error {
	query := `INSERT INTO movies (id, title) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, m.ID, m.Title)
	if err != nil {
		return util.ParsePQError(err)
	}

	return nil
}

func (r *MoviesRepo) GetMovie(ctx context.Context) ([]entity.Movie, error) {
	query := `SELECT id, title FROM movies WHERE deleted_at IS NULL`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, util.ParsePQError(err)
	}
	defer rows.Close()

	var ms []entity.Movie
	for rows.Next() {
		var m entity.Movie
		if err := rows.Scan(&m.ID, &m.Title); err != nil {
			return nil, util.ParsePQError(err)
		}
		ms = append(ms, m)
	}

	if err := rows.Err(); err != nil {
		return nil, util.ParsePQError(err)
	}

	return ms, nil
}

func (r *MoviesRepo) GetMovieByID(ctx context.Context, id string) (*entity.Movie, error) {
	query := `SELECT id, title FROM movies WHERE id = $1 AND deleted_at IS NULL`
	var m entity.Movie
	err := r.db.QueryRowContext(ctx, query, id).Scan(&m.ID, &m.Title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.ParsePQError(err)
		}
		return nil, util.ParsePQError(err)
	}

	return &m, nil
}

func (r *MoviesRepo) GetMovieAndStudio(ctx context.Context, id string) (*entity.MovieAndStudio, error) {
	query := `SELECT
			m.id,
			m.title,
			s.name_studio,
			s.starting,
			s.ending
			FROM movies m
			INNER JOIN studio_schedules s ON m.id = s.movie_id
			WHERE m.id = $1 
			ORDER BY s.starting`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, util.ParsePQError(err)
	}
	defer rows.Close()

	var movie entity.MovieAndStudio
	var sch []entity.JadwalTayang
	for rows.Next() {
		var s entity.JadwalTayang
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&s.NameStudio,
			&s.Starting,
			&s.Ending,
		)
		if err != nil {
			return nil, err
		}
		sch = append(sch, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	if movie.ID == "" {
		return nil, sql.ErrNoRows
	}

	movie.JadwalTayang = sch
	return &movie, nil
}

func (r *MoviesRepo) EditMovie(ctx context.Context, m *entity.Movie) error {
	query := `UPDATE movies SET title = $1 WHERE id = $2 AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, m.Title, m.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return util.ParsePQError(err)
		}
		return util.ParsePQError(err)
	}

	return nil
}

func (r *MoviesRepo) DeleteMovie(ctx context.Context, id string) error {
	query := `DELETE FROM movies WHERE id = $1 AND deleted_at IS NULL`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return util.ParsePQError(err)
	}

	return nil
}
