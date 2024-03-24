package service

import (
	"VK-Test_Ex/internal/model"
	"VK-Test_Ex/internal/repository"
	"context"
)

type MovieServiceImplementation struct {
	repo repository.MovieRepo
}

func NewMovieServiceImplementation(repo repository.MovieRepo) *MovieServiceImplementation {
	return &MovieServiceImplementation{repo: repo}
}

func (m *MovieServiceImplementation) CreateMovie(ctx context.Context, movie *model.Movie) error {
	return m.repo.CreateMovie(ctx, movie)
}

func (m *MovieServiceImplementation) UpdateMovie(ctx context.Context, movie *model.Movie) error {
	return m.repo.UpdateMovie(ctx, movie)
}

func (m *MovieServiceImplementation) DeleteMovie(ctx context.Context, id int64) error {
	return m.repo.DeleteMovie(ctx, id)
}

func (m *MovieServiceImplementation) GetMovies(ctx context.Context, sortBy string) ([]model.Movie, error) {
	return m.repo.GetMovies(ctx, sortBy)
}

func (m *MovieServiceImplementation) SearchMovies(ctx context.Context, titleFragment string) ([]model.Movie, error) {
	return m.repo.SearchMovies(ctx, titleFragment)
}
