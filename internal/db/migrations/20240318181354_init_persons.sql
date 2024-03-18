-- +goose Up
-- +goose StatementBegin
CREATE TABLE PERSON
(
    id         uuid PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    first_name varchar(30),
    last_name  varchar(40)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE PERSON;
-- +goose StatementEnd
