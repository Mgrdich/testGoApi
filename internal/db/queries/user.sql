-- name: GetByUsername :one
SELECT *
FROM USERS
WHERE USERNAME = $1;



-- name: GetByID :one
SELECT *
FROM USERS
WHERE ID= $1;

-- name: CreateUser :one
INSERT INTO USERS (USERNAME, PASSWORD, ROLE)
VALUES ($1, $2, $3)
RETURNING *;