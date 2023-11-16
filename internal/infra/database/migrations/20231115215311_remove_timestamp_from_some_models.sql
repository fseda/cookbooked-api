-- +goose Up
-- +goose StatementBegin
ALTER TABLE ingredients DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
ALTER TABLE ingredients_categories DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
ALTER TABLE tags DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
ALTER TABLE units DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
ALTER TABLE recipe_ingredients DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
ALTER TABLE recipe_tags DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
