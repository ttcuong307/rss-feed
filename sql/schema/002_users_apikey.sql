-- +goose Up
ALTER TABLE users ADD COLUMN api_key varchar(64) NOT NULL;

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;