CREATE TABLE movies
(
    id           uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    title        varchar(30),
    director     varchar(30),
    release_at   date,
    ticket_price NUMERIC(6, 3),
    created_at   TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE person
(
    id         uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    first_name varchar(30),
    last_name  varchar(40)
);
