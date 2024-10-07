-- databasesetup 

-- 1. Create the Database (Optional)

-- Note: Executing CREATE DATABASE requires appropriate privileges.

CREATE DATABASE autocorrect_db;
\c autocorrect_db;

-- 2. Create the misspellings Table


CREATE TABLE IF NOT EXISTS misspellings (
    id SERIAL PRIMARY KEY,                  -- Unique identifier for each record
    proper_word VARCHAR(100) NOT NULL,      -- Correctly spelled word
    misspelling VARCHAR(100) NOT NULL UNIQUE -- Common misspelling (unique to prevent duplicates)
);

-- 3. Insert Sample Data into misspellings Table

INSERT INTO misspellings (proper_word, misspelling) VALUES
    ('abandoned', 'abandonned'),
    ('aberration', 'aberation'),
    ('abilities', 'abilityes'),
    ('abilities', 'abilties'),
    ('ability', 'abilty'),
    ('abandon', 'abondon'),
    ('about', 'abbout'),
    ('about', 'abotu'),
    ('about a', 'abouta'),
    ('about it', 'aboutit'),
    ('about the', 'aboutthe'),
    ('absence', 'abscence'),
    ('abandoned', 'abondoned'),
    ('abandoning', 'abondoning'),
    ('abandons', 'abondons'),
    ('aborigine', 'aborigene'),
    ('accessories', 'accesories'),
    ('accident', 'accidant');

-- Sample list of misspellings pulled from Wikipedia list of mispellings https://en.wikipedia.org/wiki/Wikipedia:Lists_of_common_misspellings/For_machines

-- 4. Create Indexes 

-- Index on proper_word to speed up queries filtering by proper_word
CREATE INDEX IF NOT EXISTS idx_proper_word ON misspellings (proper_word);

-- Index on misspelling to speed up lookups by misspelling
CREATE INDEX IF NOT EXISTS idx_misspelling ON misspellings (misspelling);
