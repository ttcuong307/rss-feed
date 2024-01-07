-- +goose Up
CREATE TABLE feed_follows (
  id varchar(255) NOT NULL,
  feed_id varchar(255) NOT NULL,
  user_id varchar(255) NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  FOREIGN KEY (feed_id) REFERENCES feeds (id) ON DELETE CASCADE,
  UNIQUE (feed_id, user_id)
) ENGINE=InnoDB CHARACTER SET utf8;


-- +goose Down
DELETE TABLE user_feed;