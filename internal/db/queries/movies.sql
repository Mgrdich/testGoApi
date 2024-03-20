-- name: GetAllMovies :many
SELECT *
FROM MOVIES;

-- name: GetMovie :one
SELECT *
FROM MOVIES
WHERE ID = $1;


-- name: CreateMovie :one
INSERT INTO MOVIES (TITLE, DIRECTOR, RELEASE_AT, TICKET_PRICE)
VALUES ($1, $2, $3, $4)
RETURNING *;


-- name: UpdateMovie :one
UPDATE movies
SET TITLE=$2,
    DIRECTOR=$3,
    RELEASE_AT=$4,
    TICKET_PRICE=$5,
    UPDATED_AT=$6
WHERE id = $1
RETURNING *;

-- name: DeleteMovie :exec
DELETE
FROM MOVIES
WHERE ID = $1;