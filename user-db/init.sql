CREATE TABLE IF NOT EXISTS users_info (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);
