package models

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID          uuid.UUID
	Title       string
	Director    string
	ReleaseDate time.Time
	TicketPrice float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateMovieParam struct {
	Title       string
	Director    string
	ReleaseDate time.Time
	TicketPrice float64
}

type UpdateMovieParam struct {
	Title       string
	Director    string
	ReleaseDate time.Time
	TicketPrice float64
}
