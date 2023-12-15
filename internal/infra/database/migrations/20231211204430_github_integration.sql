-- +goose Up
-- +goose StatementBegin
  ALTER TABLE IF EXISTS users ADD COLUMN IF NOT EXISTS github_id VARCHAR(255);

  CREATE TABLE oauth_tokens (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    access_token VARCHAR(255) NOT NULL,
    refresh_token VARCHAR(255),
    expires_at TIMESTAMP NOT NULL,
    refresh_token_expires_at TIMESTAMP,
    device_code VARCHAR(255),

    CONSTRAINT FK_user_tokens FOREIGN KEY (user_id) REFERENCES users (id)
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  ALTER TABLE users DROP COLUNM github_id;

  DROP TABLE oauth_tokens;
-- +goose StatementEnd
