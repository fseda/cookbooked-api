-- +goose Up
-- +goose StatementBegin
  ALTER TABLE oauth_tokens ALTER COLUMN expires_at DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  ALTER TABLE oauth_tokens ALTER COLUMN expires_at SET NOT NULL;
-- +goose StatementEnd
