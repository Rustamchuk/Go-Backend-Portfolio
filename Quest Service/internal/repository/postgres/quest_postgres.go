package postgres

import (
	"VK-Quest/internal/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type QuestPostgres struct {
	db *sqlx.DB
}

func NewQuestPostgres(db *sqlx.DB) *QuestPostgres {
	return &QuestPostgres{db: db}
}

func (q *QuestPostgres) CreateQuest(ctx context.Context, quest *model.Quest) error {
	tx, err := q.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, cost) VALUES ($1, $2) RETURNING id`, questTable)

	row := tx.QueryRowContext(ctx, query, quest.Name, quest.Cost)
	err = row.Scan(&quest.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (q *QuestPostgres) GetQuest(ctx context.Context, id int) (*model.Quest, error) {
	var quest model.Quest

	query := fmt.Sprintf(`SELECT id, name, cost FROM %s WHERE id = $1`, questTable)
	row := q.db.QueryRow(query, id)

	err := row.Scan(&quest.Id, &quest.Name, &quest.Cost)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no quest with id %d", id)
		}
		return nil, err
	}

	return &quest, nil
}
