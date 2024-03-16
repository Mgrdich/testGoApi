-- name: GetAllMovies :many
SELECT * from movies;

-- name: GetMovie :one
SELECT * from movies where id = $1;


--name: CreateMovie :one

