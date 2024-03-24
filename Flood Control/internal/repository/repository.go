package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"task/internal/model"
	"task/internal/repository/postgres"
)

type FloodControlRepo interface {
	GetStatWithInterval(ctx context.Context, stat model.RequestStat) (int, error)
	AddResponse(ctx context.Context, stat model.RequestStat) error
}

type Repository struct {
	FloodControlRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		FloodControlRepo: postgres.NewFloodPostgres(db),
	}
}
