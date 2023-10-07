-- +goose Up
-- +goose StatementBegin
DO $$ BEGIN
  CREATE TYPE unit_type AS ENUM ('weight', 'volume', 'temperature');
EXCEPTION
  WHEN duplicate_object THEN null;
END $$;

DO $$ BEGIN
  CREATE TYPE unit_system AS ENUM ('metric', 'imperial');
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
  ('farenheit', '°F', true, 1, 'temperature', 'imperial');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM units 
WHERE name IN 
('milliliter', 'liter', 'deciliter', 'teaspoon', 'tablespoon', 'fluid ounce', 'pint', 'quart', 'gallon', 'miligram', 'gram', 'kilogram', 'pound', 'ounce', 'celsius', 'farenheit') 
AND is_system_unit = true;

ALTER TABLE units DROP COLUMN IF EXISTS "type";
ALTER TABLE units DROP COLUMN IF EXISTS "system";

DROP TYPE IF EXISTS unit_type;
DROP TYPE IF EXISTS unit_system;
-- +goose StatementEnd
