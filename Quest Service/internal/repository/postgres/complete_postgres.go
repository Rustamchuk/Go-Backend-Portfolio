package postgres

import (
	"VK-Quest/internal/model"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CompleteQuestPostgres struct {
	db *sqlx.DB
}

func NewCompleteQuestPostgres(db *sqlx.DB) *CompleteQuestPostgres {
	return &CompleteQuestPostgres{db: db}
}

func (c *CompleteQuestPostgres) AddCompleteQuest(ctx context.Context, completeQuest *model.CompletedQuest) error {
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, balance) VALUES ($1, $2) RETURNING id`, completeQuestTable)

	row := tx.QueryRowContext(ctx, query, completeQuest.UserID, completeQuest.QuestID)
	err = row.Scan(&completeQuest.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
