-- +goose Up
-- +goose StatementBegin
 -- Create a script to cascade the results of deleting a user to the auth_tokens table
  CREATE OR REPLACE FUNCTION delete_auth_tokens_on_user_delete() RETURNS TRIGGER AS $$
  BEGIN
    DELETE FROM oauth_tokens WHERE user_id = OLD.id;
    RETURN OLD;
  END;
  $$ LANGUAGE plpgsql;

  CREATE TRIGGER delete_auth_tokens
  BEFORE DELETE ON users
  FOR EACH ROW EXECUTE PROCEDURE delete_auth_tokens_on_user_delete();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
  DROP TRIGGER IF EXISTS delete_auth_tokens ON users;
  DROP FUNCTION IF EXISTS delete_auth_tokens_on_user_delete();
-- +goose StatementEnd
