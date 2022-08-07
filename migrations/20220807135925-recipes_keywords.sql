
-- +migrate Up
CREATE TABLE IF NOT EXISTS recipes_keywords
(
  recipe_id INTEGER NOT NULL,
  keyword_id INTEGER NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS recipes_keywords;