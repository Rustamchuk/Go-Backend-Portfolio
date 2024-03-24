package postgres

import (
	"VK-Test_Ex/internal/model"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ActorPostgres struct {
	db *sqlx.DB
}

func NewActorPostgres(db *sqlx.DB) *ActorPostgres {
	return &ActorPostgres{db: db}
}

func (a *ActorPostgres) CreateActor(ctx context.Context, actor *model.Actor) error {
	tx, err := a.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO %s (name, gender, birth_date) VALUES ($1, $2, $3) RETURNING id`, actorsTable)

	row := tx.QueryRowContext(ctx, query, actor.Name, actor.Gender, actor.BirthDate)
	err = row.Scan(&actor.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (a *ActorPostgres) UpdateActor(ctx context.Context, actor *model.Actor) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if actor.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, actor.Name)
		argId++
	}

	if actor.Gender != "" {
		setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
		args = append(args, actor.Gender)
		argId++
	}

	if actor.BirthDate.IsZero() {
		setValues = append(setValues, fmt.Sprintf("birth_date=$%d", argId))
		args = append(args, actor.BirthDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, actorsTable, setQuery, argId)
	args = append(args, actor.ID)
	_, err := a.db.ExecContext(ctx, query, args...)
	return err
}

func (a *ActorPostgres) DeleteActor(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, actorsTable)
	_, err := a.db.ExecContext(ctx, query, id)
	return err
}

func (a *ActorPostgres) GetActors(ctx context.Context) ([]model.Actor, error) {
	var items []model.Actor
	query := fmt.Sprintf(`SELECT a.id, a.name, a.gender, a.birth_date AS movies FROM %s a 
                                LEFT JOIN %s ma ON ma.actor_id = a.id 
                                LEFT JOIN %s m ON ma.movie_id = m.id GROUP BY a.id`,
		actorsTable, movieActorsTable, moviesTable)
	if err := a.db.Select(&items, query); err != nil {
		return nil, err
	}

	return items, nil
}
