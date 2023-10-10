-- +goose Up
-- +goose StatementBegin
ALTER TABLE ingredients_categories RENAME COLUMN name TO category;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE ingredients_categories RENAME COLUMN category TO name;
-- +goose StatementEnd
