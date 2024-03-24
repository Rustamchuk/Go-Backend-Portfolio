package postgres

import (
	"VK-Test_Ex/internal/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
	"log"
)

const (
	actorsTable      = "actors"
	moviesTable      = "movies"
	movieActorsTable = "movie_actors"
	usersTable       = "users"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password))
	log.Printf("Подключение к DB: host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode, cfg.Password)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Store struct {
	db *sql.DB
}

func (s *Store) GetMovies(ctx context.Context, sortBy string) ([]model.Movie, error) {
	var movies []model.Movie
	query := `SELECT id, title, description, release_date, rating FROM movies`

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

	rows, err := s.db.QueryContext(ctx, query)
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

func (s *Store) SearchMovies(ctx context.Context, titleFragment string) ([]model.Movie, error) {
	var movies []model.Movie
	query := `SELECT id, title, description, release_date, rating FROM movies WHERE title ILIKE '%' || $1 || '%'`

	rows, err := s.db.QueryContext(ctx, query, titleFragment)
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
