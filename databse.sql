CREATE TABLE IF NOT EXISTS misspellings (
    id SERIAL PRIMARY KEY,
    proper_word VARCHAR(100) NOT NULL,
    misspelling VARCHAR(100) NOT NULL
);

