-- +goose Up
-- +goose StatementBegin
-- SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'users' AND table_type = 'BASE TABLE'

-- CREATE TABLE "users" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"username" varchar(255) NOT NULL UNIQUE,"email" varchar(255) NOT NULL UNIQUE,"password_hash" text NOT NULL,"role" text NOT NULL DEFAULT 'user',PRIMARY KEY ("id"))

-- CREATE INDEX IF NOT EXISTS "idx_users_deleted_at" ON "users" ("deleted_at")

-- SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'ingredients' AND table_type = 'BASE TABLE'

-- CREATE TABLE "ingredients" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"name" varchar(100) NOT NULL,"icon" varchar(5),"is_system_ingredient" boolean NOT NULL DEFAULT false,"user_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_users_ingredients" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE)

-- CREATE UNIQUE INDEX IF NOT EXISTS "user_ingredients_must_be_unique" ON "ingredients" ("name","is_system_ingredient","user_id")
-- CREATE INDEX IF NOT EXISTS "idx_ingredients_deleted_at" ON "ingredients" ("deleted_at")

-- SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'recipes' AND table_type = 'BASE TABLE'

-- CREATE TABLE "recipes" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"title" varchar(255) NOT NULL,"description" text NOT NULL,"body" text NOT NULL,"link" varchar(500) NOT NULL,"user_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_users_recipes" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE)

-- CREATE UNIQUE INDEX IF NOT EXISTS "unique_user_id_title" ON "recipes" ("title","user_id")
-- CREATE INDEX IF NOT EXISTS "idx_recipes_deleted_at" ON "recipes" ("deleted_at")

--  CREATE TABLE "units" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"name" varchar(50) NOT NULL,"symbol" varchar(10) UNIQUE,"type" text,"system" text,"is_system_unit" boolean NOT NULL DEFAULT false,"user_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_users_units" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE)
-- CREATE UNIQUE INDEX IF NOT EXISTS "user_units_must_be_unique" ON "units" ("name","is_system_unit","user_id")

-- CREATE INDEX IF NOT EXISTS "idx_units_deleted_at" ON "units" ("deleted_at")

-- SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'recipe_ingredients' AND table_type = 'BASE TABLE'

-- CREATE TABLE "recipe_ingredients" ("id" bigserial NOT NULL,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"quantity" decimal NOT NULL,"recipe_id" bigint,"ingredient_id" bigint,"unit_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_ingredients_recipe_ingredients" FOREIGN KEY ("ingredient_id") REFERENCES "ingredients"("id") ON DELETE RESTRICT,CONSTRAINT "fk_units_recipe_ingredients" FOREIGN KEY ("unit_id") REFERENCES "units"("id") ON DELETE RESTRICT,CONSTRAINT "fk_recipes_recipe_ingredients" FOREIGN KEY ("recipe_id") REFERENCES "recipes"("id") ON DELETE CASCADE)

-- CREATE UNIQUE INDEX IF NOT EXISTS "recipe_ingredients_must_be_unique" ON "recipe_ingredients" ("recipe_id","ingredient_id")
-- CREATE INDEX IF NOT EXISTS "idx_recipe_ingredients_deleted_at" ON "recipe_ingredients" ("deleted_at")


-- SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'tags' AND table_type = 'BASE TABLE'
-- CREATE TABLE "tags" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"name" varchar(100) NOT NULL,"is_system_tag" boolean NOT NULL DEFAULT false,"user_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_users_tags" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE)
-- CREATE UNIQUE INDEX IF NOT EXISTS "user_tags_must_be_unique" ON "tags" ("name","is_system_tag","user_id")

-- CREATE INDEX IF NOT EXISTS "idx_tags_deleted_at" ON "tags" ("deleted_at")
-- SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'recipe_tags' AND table_type = 'BASE TABLE'

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

