-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM ('user', 'admin');
ALTER TABLE users ADD COLUMN role user_role DEFAULT 'user' NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN role;
DROP TYPE user_role;
-- +goose StatementEnd
