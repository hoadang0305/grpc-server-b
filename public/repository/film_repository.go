package repository

import (
	"context"

	"github.com/hoadang0305/grpc-server-b/public/domain/entity"
)

type FilmRepository interface {
	GetFilmByID(c context.Context, id int64) (*entity.Film, error)
	DeleteFilm(c context.Context, id int64) error
	CreateFilm(c context.Context, film *entity.Film) error
	UpdateFilm(c context.Context, film *entity.Film, filmId int64) (*entity.Film, error)
	GetAllFilms(c context.Context) []entity.Film
}
