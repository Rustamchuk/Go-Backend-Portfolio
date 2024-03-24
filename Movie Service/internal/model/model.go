package model

import "time"

type User struct {
	Id       int    `json:"-"`
	Role     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Actor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
}

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      float64   `json:"rating"`
	Actors      []Actor   `json:"actors,omitempty"`
}

type SignInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateActor struct {
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	BirthDate time.Time `json:"birth_date"`
}

type UpdateMovie struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"release_date"`
	Rating      float64   `json:"rating"`
	Actors      []Actor   `json:"actors,omitempty"`
}

type MovieActor struct {
	MovieID int `json:"movie_id"`
	ActorID int `json:"actor_id"`
}
