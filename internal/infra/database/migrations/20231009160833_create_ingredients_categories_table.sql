-- +goose Up
-- +goose StatementBegin
CREATE TABLE "ingredients_categories" (
  id BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL UNIQUE,
  "description" TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMPTZ
);
CREATE INDEX "ingredients_categories_deleted_at" ON "ingredients_categories" ("deleted_at");

ALTER TABLE ingredients ADD COLUMN IF NOT EXISTS category_id BIGINT REFERENCES ingredients_categories (id) ON DELETE SET NULL ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE ingredients DROP COLUMN IF EXISTS category_id;

DROP TABLE IF EXISTS "ingredients_categories";
-- +goose StatementEnd
