package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"task/internal/model"
)

type FloodPostgres struct {
	db *sqlx.DB
}

func NewFloodPostgres(db *sqlx.DB) *FloodPostgres {
	return &FloodPostgres{db: db}
}

func (f *FloodPostgres) GetStatWithInterval(ctx context.Context, stat model.RequestStat) (int, error) {
	var count int
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s WHERE user_id = $1 AND time >= $2`, floodTable)
	if err := f.db.QueryRowContext(ctx, query, stat.UserId, stat.Time).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (f *FloodPostgres) AddResponse(ctx context.Context, stat model.RequestStat) error {
	_, err := f.db.ExecContext(ctx,
		fmt.Sprintf("INSERT INTO %s (user_id, time) VALUES ($1, $2)", floodTable),
		stat.UserId, stat.Time)
	return err
}
