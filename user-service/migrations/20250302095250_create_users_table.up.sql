-- обновить id -> uuid
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(64) NOT NULL UNIQUE,
    password TEXT NOT NULL
)