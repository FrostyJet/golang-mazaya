
-- +migrate Up
CREATE TABLE IF NOT EXISTS recipe
(
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    
    preparation_time INTEGER,
    preparation_time_unit VARCHAR(25),

    cooking_time INTEGER,
    cooking_time_unit VARCHAR(25),

    serves INTEGER,
    difficulty VARCHAR(25),

    description_text TEXT,
    description_html TEXT,
    poster TEXT,

    date_created TIMESTAMPTZ NOT NULL,
    date_updated TIMESTAMPTZ
);


-- +migrate Down
DROP TABLE IF EXISTS recipe;