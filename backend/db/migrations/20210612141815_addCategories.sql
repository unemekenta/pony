
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE categories (
  id SERIAL PRIMARY KEY NOT NULL,
  content varchar(255) NOT NULL,
  created_at TIMESTAMP default CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE categories;
