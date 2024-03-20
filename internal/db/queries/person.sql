-- name: GetAllPerson :many
SELECT *
FROM PERSON;

-- name: GetPerson :one
SELECT *
FROM PERSON
WHERE ID = $1;


-- name: CreatePerson :one
INSERT INTO PERSON (first_name, last_name)
VALUES ($1, $2)
RETURNING *;