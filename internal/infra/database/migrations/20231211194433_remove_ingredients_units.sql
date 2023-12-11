-- +goose Up
-- +goose StatementBegin
  DROP TABLE IF EXISTS recipe_ingredients CASCADE;
  DROP TABLE IF EXISTS ingredients_categories CASCADE;
  DROP TABLE IF EXISTS ingredients CASCADE;
  DROP TABLE IF EXISTS units CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

  -- Ingredients ----------------
    SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'ingredients' AND table_type = 'BASE TABLE'

    CREATE TABLE "ingredients" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"name" varchar(100) NOT NULL,"icon" varchar(5),"is_system_ingredient" boolean NOT NULL DEFAULT false,"user_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_users_ingredients" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE)

    CREATE UNIQUE INDEX IF NOT EXISTS "user_ingredients_must_be_unique" ON "ingredients" ("name","is_system_ingredient","user_id")
    CREATE INDEX IF NOT EXISTS "idx_ingredients_deleted_at" ON "ingredients" ("deleted_at")

    INSERT INTO ingredients_categories (name, description) 
    VALUES 
      ('vegetable', 'A plant or part of a plant used as food'),
      ('fruit', 'The mature ovary of a flowering plant'),
      ('protein', 'A source of amino acids'),
      ('dairy', 'Food produced from the milk of mammals'),
      ('spice', 'A substance used to flavor food'),
      ('grain', 'A small, hard, dry seed, with or without an attached hull or fruit layer, harvested for human or animal consumption'),
      ('other', 'Other types of ingredients'),
      ('condiment', 'A substance, such as a sauce or seasoning, added to food to enhance its flavor'),
      ('legume', 'The fruit or seed of any of various bean or pea plants');


    INSERT INTO ingredients (name, is_system_ingredient, user_id, category_id)
    VALUES
    ('steak', true, 1, 3), -- protein
    ('chicken breast', true, 1, 3), -- protein
    ('chicken thigh', true, 1, 3), -- protein
    ('skinless chicken thigh', true, 1, 3), -- protein
    ('skinless boneless chicken thigh', true, 1, 3), -- protein
    ('white rice', true, 1, 6), -- grain
    ('black beans', true, 1, 9), -- legume
    ('carrot', true, 1, 1), -- vegetable
    ('onion', true, 1, 1), -- vegetable
    ('salt', true, 1, 5), -- spice
    ('black pepper', true, 1, 5), -- spice
    ('smoked paprika', true, 1, 5), -- spice
    ('sweet paprika', true, 1, 5), -- spice
    ('spicy paprika', true, 1, 5), -- spice
    ('white sugar', true, 1, 7), -- other
    ('orange peel', true, 1, 2), -- fruit
    ('lemon peel', true, 1, 2), -- fruit
    ('mango', true, 1, 2), -- fruit
    ('apple', true, 1, 2), -- fruit
    ('green apple', true, 1, 2), -- fruit
    ('blueberry', true, 1, 2), -- fruit
    ('blackberry', true, 1, 2), -- fruit
    ('raspberry', true, 1, 2), -- fruit
    ('strawberry', true, 1, 2), -- fruit
    ('banana', true, 1, 2), -- fruit
    ('garlic', true, 1, 1), -- vegetable
    ('bell pepper', true, 1, 1), -- vegetable
    ('spinach', true, 1, 1), -- vegetable
    ('zucchini', true, 1, 1), -- vegetable
    ('broccoli', true, 1, 1), -- vegetable
    ('cauliflower', true, 1, 1), -- vegetable
    ('mushroom', true, 1, 1), -- vegetable
    ('ginger', true, 1, 5), -- spice
    ('turmeric', true, 1, 5), -- spice
    ('chili powder', true, 1, 5), -- spice
    ('cumin', true, 1, 5), -- spice
    ('nutmeg', true, 1, 5), -- spice
    ('brown sugar', true, 1, 7), -- other
    ('honey', true, 1, 7), -- other
    ('maple syrup', true, 1, 7), -- other
    ('olive oil', true, 1, 8), -- condiment
    ('soy sauce', true, 1, 8), -- condiment
    ('tomato sauce', true, 1, 8), -- condiment
    ('mustard', true, 1, 8), -- condiment
    ('mayonnaise', true, 1, 8), -- condiment
    ('vinegar', true, 1, 8), -- condiment
    ('peach', true, 1, 2), -- fruit
    ('pear', true, 1, 2), -- fruit
    ('grape', true, 1, 2), -- fruit
    ('pineapple', true, 1, 2), -- fruit
    ('watermelon', true, 1, 2), -- fruit
    ('almond', true, 1, 7), -- other
    ('walnut', true, 1, 7), -- other
    ('peanut', true, 1, 7), -- other
    ('chia seed', true, 1, 7), -- other
    ('flax seed', true, 1, 7), -- other
    ('celery', true, 1, 1), -- vegetable
    ('parsley', true, 1, 1), -- vegetable
    ('basil', true, 1, 1), -- vegetable
    ('rosemary', true, 1, 5), -- spice
    ('cayenne pepper', true, 1, 5), -- spice
    ('tarragon', true, 1, 5), -- spice
    ('dill', true, 1, 5), -- spice
    ('capers', true, 1, 8), -- condiment
    ('thyme', true, 1, 5), -- spice
    ('red wine vinegar', true, 1, 8), -- condiment
    ('balsamic vinegar', true, 1, 8), -- condiment
    ('sesame oil', true, 1, 8), -- condiment
    ('worcestershire sauce', true, 1, 8), -- condiment
    ('barbecue sauce', true, 1, 8), -- condiment
    ('lentils', true, 1, 9), -- legume
    ('tofu', true, 1, 3), -- protein
    ('tempeh', true, 1, 3), -- protein
    ('coconut oil', true, 1, 8), -- condiment
    ('saffron', true, 1, 5), -- spice
    ('agave syrup', true, 1, 7), -- other
    ('coconut sugar', true, 1, 7), -- other
    ('cashews', true, 1, 7), -- other
    ('sesame seeds', true, 1, 7), -- other
    ('sunflower seeds', true, 1, 7), -- other
    ('pumpkin seeds', true, 1, 7), -- other
    ('avocado', true, 1, 1), -- vegetable
    ('sriracha', true, 1, 8), -- condiment
    ('hoisin sauce', true, 1, 8), -- condiment
    ('asparagus', true, 1, 1), -- vegetable
    ('eggplant', true, 1, 1); -- vegetable
  ----------------------------------------------------------------

  -- Units ----------------
    CREATE TABLE "units" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"name" varchar(50) NOT NULL,"symbol" varchar(10) UNIQUE,"type" text,"system" text,"is_system_unit" boolean NOT NULL DEFAULT false,"user_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_users_units" FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE SET NULL ON UPDATE CASCADE);
    CREATE UNIQUE INDEX IF NOT EXISTS "user_units_must_be_unique" ON "units" ("name","is_system_unit","user_id");

    CREATE INDEX IF NOT EXISTS "idx_units_deleted_at" ON "units" ("deleted_at");

    DO $$ BEGIN
      CREATE TYPE unit_type AS ENUM ('weight', 'volume', 'temperature', 'unit');
    EXCEPTION
      WHEN duplicate_object THEN null;
    END $$;

    DO $$ BEGIN
      CREATE TYPE unit_system AS ENUM ('metric', 'imperial', 'unit');
    EXCEPTION
      WHEN duplicate_object THEN null;
    END $$;

    ALTER TABLE units ADD COLUMN IF NOT EXISTS type unit_type;
    ALTER TABLE units ADD COLUMN IF NOT EXISTS system unit_system;

    INSERT INTO units (name, symbol, is_system_unit, user_id, type, system)
    VALUES
    ('milliliter', 'ml', true, 1, 'volume', 'metric'),
    ('liter', 'l', true, 1, 'volume', 'metric'),
    ('deciliter', 'dl', true, 1, 'volume', 'metric'),
    ('teaspoon', 'tsp', true, 1, 'volume', 'imperial'),
    ('tablespoon', 'tbs', true, 1, 'volume', 'imperial'),
    ('fluid ounce', 'fl oz', true, 1, 'volume', 'imperial'),
    ('pint', 'pt', true, 1, 'volume', 'imperial'),
    ('quart', 'qt', true, 1, 'volume', 'imperial'),
    ('gallon', 'gal', true, 1, 'volume', 'imperial'),
    ('miligram', 'mg', true, 1, 'weight', 'metric'),
    ('gram', 'g', true, 1, 'weight', 'metric'),
    ('kilogram', 'kg', true, 1, 'weight', 'metric'),
    ('pound', 'lb', true, 1, 'weight', 'imperial'),
    ('ounce', 'oz', true, 1, 'weight', 'imperial'),
    ('celsius', '°C', true, 1, 'temperature', 'metric'),
    ('farenheit', '°F', true, 1, 'temperature', 'imperial')
    ('unit', '', true, 1, 'unit', 'unit');
  ----------------------------------------------------------------

  -- Recipe Ingredients --------------------------------
    SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'recipe_ingredients' AND table_type = 'BASE TABLE'

    CREATE TABLE "recipe_ingredients" ("id" bigserial NOT NULL,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"quantity" decimal NOT NULL,"recipe_id" bigint,"ingredient_id" bigint,"unit_id" bigint,PRIMARY KEY ("id"),CONSTRAINT "fk_ingredients_recipe_ingredients" FOREIGN KEY ("ingredient_id") REFERENCES "ingredients"("id") ON DELETE RESTRICT,CONSTRAINT "fk_units_recipe_ingredients" FOREIGN KEY ("unit_id") REFERENCES "units"("id") ON DELETE RESTRICT,CONSTRAINT "fk_recipes_recipe_ingredients" FOREIGN KEY ("recipe_id") REFERENCES "recipes"("id") ON DELETE CASCADE)

    CREATE UNIQUE INDEX IF NOT EXISTS "recipe_ingredients_must_be_unique" ON "recipe_ingredients" ("recipe_id","ingredient_id")
    CREATE INDEX IF NOT EXISTS "idx_recipe_ingredients_deleted_at" ON "recipe_ingredients" ("deleted_at")
  ----------------------------------------------------------------

  -- All --------------------------------------------------------
    ALTER TABLE ingredients DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
    ALTER TABLE units DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
    ALTER TABLE recipe_ingredients DROP COLUMN created_at, DROP COLUMN updated_at, DROP COLUMN deleted_at;
  ----------------------------------------------------------------
-- +goose StatementEnd
