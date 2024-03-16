-- name: GetAllMovies :many
SELECT *
from movies;

-- name: GetMovie :one
SELECT *
FROM movies
WHERE ID = $1;


-- name: CreateMovie :one
INSERT INTO movies (title, director, release_at, ticket_price, updated_at)
values ($1, $2, $3, $4, $5)
RETURNING *;


-- name: UpdateMovie :one
UPDATE movies
SET title=$2,
    director=$3,
    release_at=$4,
    ticket_price=$5,
    updated_at=$6
WHERE id = $1
RETURNING *;

-- name: DeleteMovie :exec
DELETE
FROM movies
where id = $1;

