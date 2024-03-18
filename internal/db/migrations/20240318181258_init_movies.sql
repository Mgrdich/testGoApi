-- +goose Up
-- +goose StatementBegin
CREATE TABLE MOVIES
(
    id           uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    title        varchar(30),
    director     varchar(30),
    release_at   date,
    ticket_price NUMERIC(6, 3),
    created_at   TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE MOVIES;
-- +goose StatementEnd
