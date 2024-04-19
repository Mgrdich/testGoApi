-- +goose Up
-- +goose StatementBegin
ALTER TABLE PERSON ADD created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE PERSON DROP created_at
-- +goose StatementEnd
