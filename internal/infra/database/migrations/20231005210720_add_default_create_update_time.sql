-- +goose Up
-- +goose StatementBegin
DO $$ 
DECLARE 
   table_name text;
BEGIN 
   FOR table_name IN (SELECT unnest(ARRAY['users', 'units', 'tags', 'recipes', 'recipe_tags', 'recipe_ingredients', 'ingredients']))
   LOOP
      EXECUTE format('ALTER TABLE %I ALTER COLUMN created_at SET DEFAULT now()', table_name);
      EXECUTE format('ALTER TABLE %I ALTER COLUMN updated_at SET DEFAULT now()', table_name);
   END LOOP;
END $$;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DO $$ 
DECLARE 
   table_name text;
BEGIN 
   FOR table_name IN (SELECT unnest(ARRAY['users', 'units', 'tags', 'recipes', 'recipe_tags', 'recipe_ingredients', 'ingredients']))
   LOOP
      EXECUTE format('ALTER TABLE %I ALTER COLUMN created_at DROP DEFAULT', table_name);
      EXECUTE format('ALTER TABLE %I ALTER COLUMN updated_at DROP DEFAULT', table_name);
   END LOOP;
END $$;
-- +goose StatementEnd
