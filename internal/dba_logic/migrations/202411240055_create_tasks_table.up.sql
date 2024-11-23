CREATE TABLE IF NOT EXISTS "Task" (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status BOOLEAN NOT NULL,
    added_time TIMESTAMP NOT NULL,
    images Bytea,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "User" (id)
    );