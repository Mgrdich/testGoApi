CREATE TABLE movies
(
    id           uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title        varchar(30),
    director     varchar(30),
    release_at   date,
    ticket_price NUMERIC(3, 3),
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP
);