
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE images (
  id SERIAL PRIMARY KEY NOT NULL,
  url varchar NOT NULL,
  user_id int NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  created_at TIMESTAMP default CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE images;
