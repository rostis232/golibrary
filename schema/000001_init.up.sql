-- Creating table for types
CREATE TABLE type (
    id SERIAL PRIMARY KEY,
    title_en VARCHAR(255) NOT NULL,
    title_uk VARCHAR(255) NOT NULL
);

-- Creating table for difficulties
CREATE TABLE difficulty (
    id SERIAL PRIMARY KEY,
    title_en VARCHAR(255) NOT NULL,
    title_uk VARCHAR(255) NOT NULL
);

-- Creating table for languages
CREATE TABLE language (
    id SERIAL PRIMARY KEY,
    title_en VARCHAR(255) NOT NULL,
    title_uk VARCHAR(255) NOT NULL
);

-- Creating table for library items
CREATE TABLE library_item (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    short_desc_en TEXT,
    short_desc_uk TEXT,
    description_en TEXT,
    description_uk TEXT,
    url VARCHAR(255),
    type INT REFERENCES type(id),
    difficulty INT REFERENCES difficulty(id),
    language INTEGER[]
);