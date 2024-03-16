// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: movies.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (title, director, release_at, ticket_price, updated_at)
values ($1, $2, $3, $4, $5)
RETURNING id, title, director, release_at, ticket_price, created_at, updated_at
`

type CreateMovieParams struct {
	Title       pgtype.Text
	Director    pgtype.Text
	ReleaseAt   pgtype.Date
	TicketPrice pgtype.Numeric
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRow(ctx, createMovie,
		arg.Title,
		arg.Director,
		arg.ReleaseAt,
		arg.TicketPrice,
		arg.UpdatedAt,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Director,
		&i.ReleaseAt,
		&i.TicketPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE
FROM movies
where id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteMovie, id)
	return err
}

const getAllMovies = `-- name: GetAllMovies :many
SELECT id, title, director, release_at, ticket_price, created_at, updated_at
from movies
`

func (q *Queries) GetAllMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.db.Query(ctx, getAllMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Director,
			&i.ReleaseAt,
			&i.TicketPrice,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMovie = `-- name: GetMovie :one
SELECT id, title, director, release_at, ticket_price, created_at, updated_at
FROM movies
WHERE ID = $1
`

func (q *Queries) GetMovie(ctx context.Context, id pgtype.UUID) (Movie, error) {
	row := q.db.QueryRow(ctx, getMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Director,
		&i.ReleaseAt,
		&i.TicketPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateMovie = `-- name: UpdateMovie :one
UPDATE movies
SET title=$2,
    director=$3,
    release_at=$4,
    ticket_price=$5,
    updated_at=$6
WHERE id = $1
RETURNING id, title, director, release_at, ticket_price, created_at, updated_at
`

type UpdateMovieParams struct {
	ID          pgtype.UUID
	Title       pgtype.Text
	Director    pgtype.Text
	ReleaseAt   pgtype.Date
	TicketPrice pgtype.Numeric
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) (Movie, error) {
	row := q.db.QueryRow(ctx, updateMovie,
		arg.ID,
		arg.Title,
		arg.Director,
		arg.ReleaseAt,
		arg.TicketPrice,
		arg.UpdatedAt,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Director,
		&i.ReleaseAt,
		&i.TicketPrice,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
