
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE categories_of_comments (
  id SERIAL PRIMARY KEY NOT NULL,
  comment_id int NOT NULL REFERENCES comments(id) ON DELETE CASCADE,
  category_id int NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
  created_at TIMESTAMP default CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE categories_of_comments;
