package repository

import (
	"biFebriansyah/back/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoMovie struct {
	*sqlx.DB
}

func NewMovie(db *sqlx.DB) *RepoMovie {
	return &RepoMovie{db}
}

func (r *RepoMovie) CreateMovie(data *models.Movie) (string, error) {
	q := `INSERT INTO public.movie(
		movie_name,
		slug_movie,
		movie_banner,
		release_date)
	VALUES(
		:movie_name,
		:slug_movie,
		:movie_banner,
		:release_date
	)`

	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 data movie created", nil

}
