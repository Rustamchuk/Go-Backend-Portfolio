package postgres

import (
	"VK-Test_Ex/internal/model"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type MoviePostgres struct {
	db *sqlx.DB
}

func NewMoviePostgres(db *sqlx.DB) *MoviePostgres {
	return &MoviePostgres{db: db}
}

func (m *MoviePostgres) CreateMovie(ctx context.Context, movie *model.Movie) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`INSERT INTO %s (title, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING id`, moviesTable)

	row := tx.QueryRowContext(ctx, query, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating)
	err = row.Scan(&movie.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, actor := range movie.Actors {
		query = fmt.Sprintf(`INSERT INTO %s (movie_id, actor_id) VALUES ($1, $2)`, movieActorsTable)
		_, err = m.db.ExecContext(ctx, query, movie.ID, actor.ID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (m *MoviePostgres) UpdateMovie(ctx context.Context, movie *model.Movie) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if movie.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, movie.Title)
		argId++
	}

	if movie.Description != "" {
		setValues = append(setValues, fmt.Sprintf("descriprion=$%d", argId))
		args = append(args, movie.Description)
		argId++
	}

	setValues = append(setValues, fmt.Sprintf("rating=$%d", argId))
	args = append(args, movie.Rating)
	argId++

	if movie.Actors != nil {
		query := fmt.Sprintf(`DELETE FROM %s WHERE movie_id = $1`, movieActorsTable)
		_, err := m.db.ExecContext(ctx, query, movie.ID)
		for _, actor := range movie.Actors {
			query = fmt.Sprintf(`INSERT INTO %s (movie_id, actor_id) VALUES ($1, $2)`, movieActorsTable)
			_, err = m.db.ExecContext(ctx, query, movie.ID, actor.ID)
			if err != nil {
				return err
			}
		}
	}

	if movie.ReleaseDate.IsZero() {
		setValues = append(setValues, fmt.Sprintf("release_date=$%d", argId))
		args = append(args, movie.ReleaseDate)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id = $%d`, moviesTable, setQuery, argId)
	args = append(args, movie.ID)
	_, err := m.db.ExecContext(ctx, query, args...)
	return err
}

func (m *MoviePostgres) DeleteMovie(ctx context.Context, id int64) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, moviesTable)
	_, err := m.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	query = fmt.Sprintf(`DELETE FROM %s WHERE movie_id = $1`, movieActorsTable)
	_, err = m.db.ExecContext(ctx, query, id)
	return err
}

func (m *MoviePostgres) GetMovies(ctx context.Context, sortBy string) ([]model.Movie, error) {
	var movies []model.Movie
	query := fmt.Sprintf(`SELECT id, title, description, release_date, rating FROM %s`, moviesTable)

	switch sortBy {
	case "title":
		query += " ORDER BY title"
	case "rating":
		query += " ORDER BY rating DESC"
	case "release_date":
		query += " ORDER BY release_date DESC"
	default:
		query += " ORDER BY rating DESC"
	}

	rows, err := m.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m model.Movie
		if err = rows.Scan(&m.ID, &m.Title, &m.Description, &m.ReleaseDate, &m.Rating); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

func (m *MoviePostgres) SearchMovies(ctx context.Context, titleFragment string) ([]model.Movie, error) {
	var movies []model.Movie
	query := fmt.Sprintf(`SELECT id, title, description, release_date, rating FROM %s 
                                WHERE title ILIKE '%' || $1 || '%'`, moviesTable)

	rows, err := m.db.QueryContext(ctx, query, titleFragment)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m model.Movie
		if err := rows.Scan(&m.ID, &m.Title, &m.Description, &m.ReleaseDate, &m.Rating); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}
