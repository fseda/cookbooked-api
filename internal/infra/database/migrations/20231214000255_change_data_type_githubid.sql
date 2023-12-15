-- +goose Up
-- +goose StatementBegin
  ALTER TABLE users ALTER COLUMN github_id TYPE bigint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  ALTER TABLE users ALTER COLUMN github_id TYPE varchar(255);
-- +goose StatementEnd
