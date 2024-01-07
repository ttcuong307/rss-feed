-- +goose Up
CREATE TABLE users(
  id varchar(255) NOT NULL,
  name varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB CHARACTER SET utf8;

-- +goose Down
DROP TABLE users;