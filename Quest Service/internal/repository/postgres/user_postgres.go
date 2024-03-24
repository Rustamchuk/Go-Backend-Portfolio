package postgres

import (
	"VK-Quest/internal/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) CreateUser(ctx context.Context, user *model.User) error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, balance) VALUES ($1, $2) RETURNING id`, userTable)

	row := tx.QueryRowContext(ctx, query, user.Name, user.Balance)
	err = row.Scan(&user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (u *UserPostgres) IncreaseUserBalance(ctx context.Context, id, value int) error {
	query := fmt.Sprintf(`UPDATE %s SET balance = balance + $1 WHERE id = $2`, userTable)

	_, err := u.db.Exec(query, value, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserPostgres) GetUser(ctx context.Context, id int) (*model.User, error) {
	var user model.User

	query := fmt.Sprintf(`SELECT id, name, balance FROM %s WHERE id = $1`, userTable)
	row := u.db.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Name, &user.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user with id %d", id)
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserPostgres) GetAchievements(ctx context.Context, id int) (*model.Achievements, error) {
	user, err := u.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	questIDs := make([]int, 0)
	rows, err := u.db.Query(fmt.Sprintf("SELECT quest_id FROM %s WHERE user_id = $1", completeQuestTable), id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var questID int
		if err = rows.Scan(&questID); err != nil {
			return nil, err
		}
		questIDs = append(questIDs, questID)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	quests := make([]model.Quest, 0)
	for _, qID := range questIDs {
		var quest model.Quest
		err = u.db.QueryRow(
			fmt.Sprintf("SELECT id, name, cost FROM %s WHERE id = $1", questTable), qID).Scan(&quest.Id, &quest.Name, &quest.Cost)
		if err != nil {
			return nil, err
		}
		quests = append(quests, quest)
	}

	return &model.Achievements{
		Balance: user.Balance,
		Quests:  quests,
	}, nil
}
