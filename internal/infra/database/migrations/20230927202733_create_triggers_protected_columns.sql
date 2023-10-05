-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION prevent_update_of_default_system_values_function_tags()
  RETURNS TRIGGER AS $$
	BEGIN
		IF OLD.is_system_tag = TRUE THEN
			IF NEW.name IS DISTINCT FROM OLD.name THEN
				RAISE EXCEPTION 'Cannot update "name" column of (tags) system default values';
			END IF;
		END IF;

		IF NEW.is_system_tag IS DISTINCT FROM OLD.is_system_tag THEN
			RAISE EXCEPTION 'Cannot update "is_system_tag" column of (tags)';
		END IF;
		
		RETURN NEW;
	END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_update_of_default_system_values_function_ingredients()
  RETURNS TRIGGER AS $$
	BEGIN
		IF OLD.is_system_ingredient = TRUE THEN
			IF NEW.name IS DISTINCT FROM OLD.name THEN
				RAISE EXCEPTION 'Cannot update "name" column of (ingredients) system default values';
			END IF;
		END IF;

		IF NEW.is_system_ingredient IS DISTINCT FROM OLD.is_system_ingredient THEN
			RAISE EXCEPTION 'Cannot update "is_system_ingredient" column of (ingredients)';
		END IF;
		
		RETURN NEW;
	END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION prevent_update_of_default_system_values_function_units()
  RETURNS TRIGGER AS $$
	BEGIN
		IF OLD.is_system_unit = TRUE THEN
			IF NEW.name IS DISTINCT FROM OLD.name THEN
				RAISE EXCEPTION 'Cannot update "name" column of (units) system default values';
			END IF;
		END IF;

		IF NEW.is_system_unit IS DISTINCT FROM OLD.is_system_unit THEN
			RAISE EXCEPTION 'Cannot update "is_system_unit" column of (units)';
		END IF;
		
		RETURN NEW;
	END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER prevent_update_of_default_system_values_tags
  BEFORE UPDATE ON tags
  FOR EACH ROW
  EXECUTE FUNCTION prevent_update_of_default_system_values_function_tags();

CREATE OR REPLACE TRIGGER prevent_update_of_default_system_values_ingredients
  BEFORE UPDATE ON ingredients
  FOR EACH ROW
  EXECUTE FUNCTION prevent_update_of_default_system_values_function_ingredients();

CREATE OR REPLACE TRIGGER prevent_update_of_default_system_values_units
  BEFORE UPDATE ON ingredients
  FOR EACH ROW
  EXECUTE FUNCTION prevent_update_of_default_system_values_function_units();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS prevent_update_of_default_system_values_function_tags CASCADE;
DROP FUNCTION IF EXISTS prevent_update_of_default_system_values_function_ingredients CASCADE;
DROP FUNCTION IF EXISTS prevent_update_of_default_system_values_function_units CASCADE;

DROP TRIGGER IF EXISTS prevent_update_of_default_system_values_tags ON tags CASCADE;
DROP TRIGGER IF EXISTS prevent_update_of_default_system_values_ingredients ON ingredients CASCADE;
DROP TRIGGER IF EXISTS prevent_update_of_default_system_values_units ON units CASCADE;
-- +goose StatementEnd
