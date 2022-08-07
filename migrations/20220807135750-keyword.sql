
-- +migrate Up
CREATE TABLE IF NOT EXISTS keyword
(
  id SERIAL NOT NULL PRIMARY KEY,
  text TEXT NOT NULL,

  date_created TIMESTAMPTZ NOT NULL,
  date_updated TIMESTAMPTZ
);

-- +migrate Down
DROP TABLE IF EXISTS keyword;