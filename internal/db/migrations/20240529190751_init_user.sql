-- +goose Up
-- +goose StatementBegin
CREATE TYPE ROLE AS ENUM ('admin', 'user');
CREATE TABLE USERS
(
    id uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    username VARCHAR(30) UNIQUE,
    password VARCHAR(72),
    role ROLE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE USERS;
DROP TYPE IF EXISTS ROLE;
-- +goose StatementEnd