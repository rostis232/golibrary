-- Creating table for types
CREATE TABLE type (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

-- Creating table for difficulties
CREATE TABLE difficulty (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

-- Creating table for languages
CREATE TABLE language (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL
);

-- Creating table for library items
CREATE TABLE library_item (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    short_desc TEXT,
    description TEXT,
    url VARCHAR(255),
    type INT REFERENCES type(id),
    difficulty INT REFERENCES difficulty(id),
    language INTEGER[]
);