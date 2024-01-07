-- +goose Up
ALTER TABLE feeds ADD COLUMN last_fetched_at datetime NULL;

-- +goose Down
ALTER TABLE feeds REMOVE COLUMN last_fetched_at;
