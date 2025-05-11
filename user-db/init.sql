CREATE TABLE IF NOT EXISTS users_info (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
);

-- Таблица результатов упрощения текста
CREATE TABLE IF NOT EXISTS simplification_results (
    id SERIAL PRIMARY KEY,
    original TEXT NOT NULL,
    simplified TEXT NOT NULL,
    user_login TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_login) REFERENCES users_info(login) ON DELETE CASCADE
);