-- +goose Up
CREATE TABLE posts (
  id varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  title text NOT NULL,
  url varchar(255) NOT NULL,
  description text DEFAULT NULL,
  published_at datetime DEFAULT NULL,
  feed_id varchar(255) NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (url),
  FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE
) ENGINE=InnoDB CHARACTER SET utf8;

-- +goose Down
DROP TABLE posts;