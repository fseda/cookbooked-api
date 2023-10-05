-- +goose Up
-- +goose StatementBegin
DO $$ BEGIN
    CREATE TYPE user_role AS ENUM ('user', 'admin');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;
ALTER TABLE users ADD COLUMN role user_role DEFAULT 'user' NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN role;
DROP TYPE user_role;
-- +goose StatementEnd
