-- +goose Up
-- +goose StatementBegin

  ALTER TABLE users ADD COLUMN "name" VARCHAR(255);
  ALTER TABLE users ADD COLUMN "bio" TEXT;
  ALTER TABLE users ADD COLUMN "avatar_url" VARCHAR(255);
  ALTER TABLE users ADD COLUMN "location" VARCHAR(255);


-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

  ALTER TABLE users DROP COLUMN "name";
  ALTER TABLE users DROP COLUMN "bio";
  ALTER TABLE users DROP COLUMN "avatar_url";
  ALTER TABLE users DROP COLUMN "location";

-- +goose StatementEnd
