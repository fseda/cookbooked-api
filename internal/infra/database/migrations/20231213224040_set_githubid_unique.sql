-- +goose Up
-- +goose StatementBegin
  ALTER TABLE users ADD CONSTRAINT users_githubid_unique UNIQUE (github_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  ALTER TABLE users DROP CONSTRAINT users_githubid_unique;
-- +goose StatementEnd
