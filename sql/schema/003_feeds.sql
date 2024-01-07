-- +goose Up
CREATE TABLE feeds (
  id varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  url varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  user_id varchar(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (url),
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
) ENGINE=InnoDB CHARACTER SET utf8;

-- +goose Down
DELETE TABLE feeds;