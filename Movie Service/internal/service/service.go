package service

import (
	"VK-Test_Ex/internal/model"
	"VK-Test_Ex/internal/repository"
	"context"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type ActorService interface {
	CreateActor(ctx context.Context, actor *model.Actor) error
	UpdateActor(ctx context.Context, actor *model.Actor) error
	DeleteActor(ctx context.Context, id int64) error
	GetActors(ctx context.Context) ([]model.Actor, error)
}

type MovieService interface {
	CreateMovie(ctx context.Context, movie *model.Movie) error
	UpdateMovie(ctx context.Context, movie *model.Movie) error
	DeleteMovie(ctx context.Context, id int64) error
	GetMovies(ctx context.Context, sortBy string) ([]model.Movie, error)
	SearchMovies(ctx context.Context, titleFragment string) ([]model.Movie, error)
}

type Service struct {
	Authorization
	ActorService
	MovieService
	userRoles []string
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		ActorService:  NewActorServiceImplementation(repos.ActorRepo),
		MovieService:  NewMovieServiceImplementation(repos.MovieRepo),
		userRoles:     make([]string, 1),
	}
}

func (s *Service) GetRole(id int) string {
	if id >= len(s.userRoles) {
		return ""
	}
	return s.userRoles[id]
}

func (s *Service) AddRole(role string) {
	s.userRoles = append(s.userRoles, role)
}
