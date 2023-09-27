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


CREATE OR REPLACE FUNCTION prevent_update_of_protected_columns(
	protect_column TEXT
)
	RETURNS TRIGGER AS $$
	BEGIN
		IF NEW.protected_column IS DISTINCT FROM OLD.protected_column THEN
			RAISE EXCEPTION 'Cannot update protected column: %', protect_column;
		END IF;

		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_update_of_is_system_tag
	BEFORE UPDATE ON tags
	FOR EACH ROW
	EXECUTE FUNCTION prevent_update_of_protected_columns('is_system_tag');

CREATE TRIGGER prevent_update_of_is_system_ingredient
	BEFORE UPDATE ON ingredients
	FOR EACH ROW
	EXECUTE FUNCTION prevent_update_of_protected_columns('is_system_ingredient');

CREATE TRIGGER prevent_update_of_is_system_unit
	BEFORE UPDATE ON units
	FOR EACH ROW
	EXECUTE FUNCTION prevent_update_of_protected_columns('is_system_unit');


CREATE OR REPLACE FUNCTION prevent_update_of_default_system_name_column(
	name_column TEXT,
	is_system_column TEXT
)
	RETURNS TRIGGER AS $$
	BEGIN
		IF NEW.is_system_column = TRUE THEN
			IF NEW.name_column IS DISTINCT FROM OLD.name_column THEN
				RAISE EXCEPTION 'Cannot update "%" column of (%) system default values', name_column, is_system_column;
			END IF;
		END IF;
		
		RETURN NEW;
	END;
	$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_update_of_name_of_system_tags
	BEFORE UPDATE ON tags
	FOR EACH ROW
	EXECUTE FUNCTION prevent_update_of_default_system_name_column("name", "is_system_tag");

CREATE TRIGGER prevent_update_of_name_of_system_ingredients
	BEFORE UPDATE ON ingredients
	FOR EACH ROW
	EXECUTE FUNCTION prevent_update_of_default_system_name_column("name", "is_system_ingredient");

CREATE TRIGGER prevent_update_of_name_of_system_units
	BEFORE UPDATE ON units
	FOR EACH ROW
	EXECUTE FUNCTION prevent_update_of_default_system_name_column("name", "is_system_unit");

