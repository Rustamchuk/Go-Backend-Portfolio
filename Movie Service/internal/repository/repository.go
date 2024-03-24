package repository

import (
	"VK-Test_Ex/internal/model"
	"VK-Test_Ex/internal/repository/postgres"
	"context"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type ActorRepo interface {
	CreateActor(ctx context.Context, actor *model.Actor) error
	UpdateActor(ctx context.Context, actor *model.Actor) error
	DeleteActor(ctx context.Context, id int64) error
	GetActors(ctx context.Context) ([]model.Actor, error)
}

type MovieRepo interface {
	CreateMovie(ctx context.Context, movie *model.Movie) error
	UpdateMovie(ctx context.Context, movie *model.Movie) error
	DeleteMovie(ctx context.Context, id int64) error
	GetMovies(ctx context.Context, sortBy string) ([]model.Movie, error)
	SearchMovies(ctx context.Context, titleFragment string) ([]model.Movie, error)
}

type Repository struct {
	Authorization
	ActorRepo
	MovieRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		ActorRepo:     postgres.NewActorPostgres(db),
		MovieRepo:     postgres.NewMoviePostgres(db),
	}
}
