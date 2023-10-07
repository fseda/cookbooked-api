-- +goose Up
-- +goose StatementBegin
ALTER TABLE units ADD COLUMN IF NOT EXISTS symbol VARCHAR(10) UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE units DROP COLUMN IF EXISTS symbol;
-- +goose StatementEnd
