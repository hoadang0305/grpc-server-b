package service

import (
	"context"

	"github.com/hoadang0305/grpc-server-b/public/domain/model"

	"github.com/hoadang0305/grpc-server-b/public/domain/entity"
)

type FilmService interface {
	GetFilmById(ctx context.Context, id int64) (*entity.Film, error)
	DeleteFilm(ctx context.Context, id int64) error
	CreateFilm(c context.Context, filmRequest model.FilmRequest) (*entity.Film, error)
	UpdateFilm(c context.Context, filmRequest model.FilmRequest, filmId int64) (*entity.Film, error)
	GetAllFilms(ctx context.Context) []entity.Film
}
