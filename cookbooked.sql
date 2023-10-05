CREATE TABLE users (
	id            SERIAL PRIMARY KEY,
	username      VARCHAR(255) UNIQUE NOT NULL,
	email 	 	    VARCHAR(255) UNIQUE NOT NULL,
	password_hash TEXT NOT NULL,
);

CREATE TABLE recipes (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	description TEXT,
	body text,
	link varchar(500),
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
	user_id INT,

	CONSTRAINT FK_recipes_users FOREIGN KEY (user_id)
		REFERENCES users (id)
		ON DELETE SET NULL
	CONSTRAINT user_recipes_title_must_be_unique UNIQUE (user_id, title)
);

CREATE TABLE tags (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100),
	is_system_tag BOOLEAN DEFAULT FALSE NOT NULL,
	user_id INT,
	
	CONSTRAINT user_tags_must_be_unique UNIQUE (user_id, name, is_system_tag),
	CONSTRAINT FK_tags_users FOREIGN KEY (user_id)
		REFERENCES users (id)
		ON DELETE SET NULL
);

CREATE TABLE ingredients (
	id SERIAL PRIMARY KEY, 
	name VARCHAR(100) NOT NULL, 
	icon TEXT, 
	is_system_ingredient BOOLEAN DEFAULT FALSE NOT NULL, 
	user_id INT, 
	CONSTRAINT user_ingredients_must_be_unique UNIQUE (user_id, name, is_system_ingredient), 
	CONSTRAINT FK_ingredients_users FOREIGN KEY (user_id) REFERENCES users (id)
		ON DELETE SET NULL
);

CREATE TABLE units (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL
	is_system_unit BOOLEAN DEFAULT FALSE NOT NULL,
	user_id INT,

	CONSTRAINT user_units_must_be_unique UNIQUE (user_id, name, is_system_unit)
	CONSTRAINT FK_units_users FOREIGN KEY (user_id)
		REFERENCES users (id)
		ON DELETE CASCADE;
);

CREATE TABLE recipe_tags (
	id SERIAL PRIMARY KEY,
	recipe_id INT NOT NULL,
	tag_id INT NOT NULL,

	CONSTRAINT recipe_tags_must_be_unique UNIQUE (recipe_id, tag_id),
	CONSTRAINT FK_recipe_tags_recipes FOREIGN KEY (recipe_id)
		REFERENCES recipes (id)
		ON DELETE CASCADE,
	CONSTRAINT FK_recipe_tags_tags FOREIGN KEY (tag_id)
		REFERENCES tags (id)
		ON DELETE CASCADE
);

CREATE TABLE recipe_ingredients (
	id SERIAL PRIMARY KEY,
	quantity INT NOT NULL,

	recipe_id INT NOT NULL,
	ingredient_id INT NOT NULL,
	unit_id INT NOT NULL,

	CONSTRAINT recipe_ingredients_must_be_unique UNIQUE (recipe_id, ingredient_id)
	CONSTRAINT FK_recipe_ingredients_recipe FOREIGN KEY (recipe_id)
		REFERENCES recipes (id)
		ON DELETE CASCADE,
	
	CONSTRAINT FK_recipe_ingredients_ingredient FOREIGN KEY (ingredient_id)
		REFERENCES ingredient (id)
		ON DELETE RESTRICT,
	
	CONSTRAINT FK_recipe_ingredients_unit FOREIGN KEY (unit_id)
		REFERENCES units (id)
		ON DELETE RESTRICT;
);

CREATE OR REPLACE FUNCTION prevent_update_of_default_system_values_function()
  RETURNS TRIGGER AS $$
	DECLARE
		table_name text;
		is_system_column text;
		name_column text;
	BEGIN
		-- Get the name of the table associated with the trigger
		table_name := TG_TABLE_NAME;

		-- Determine the corresponding column names
		CASE table_name
			WHEN 'tags' THEN
				is_system_column := 'is_system_tag';
				name_column := 'name';
			WHEN 'ingredients' THEN
				is_system_column := 'is_system_ingredient';
				name_column := 'name';
			WHEN 'units' THEN
				is_system_column := 'is_system_unit';
				name_column := 'name';
			-- Add more cases for other tables if needed
		END CASE;

		IF OLD.is_system_column = TRUE THEN
			IF NEW.name_column IS DISTINCT FROM OLD.name_column THEN
				RAISE EXCEPTION 'Cannot update "%" column of (%) system default values', name_column, table_name;
			END IF;
		END IF;

		IF NEW.is_system_column IS DISTINCT FROM OLD.is_system_column THEN
			RAISE EXCEPTION 'Cannot update "%" column of (%)', is_system_column, table_name;
		END IF;
		
		RETURN NEW;
	END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_update_of_default_system_values
  BEFORE UPDATE ON tags, ingredients, units
  FOR EACH ROW
  EXECUTE FUNCTION prevent_update_of_default_system_values_function();
	