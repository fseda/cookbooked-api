-- +goose Up
-- +goose StatementBegin
ALTER TABLE units ADD COLUMN symbol VARCHAR(10) UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE units DROP COLUMN symbol;
-- +goose StatementEnd
