CREATE TABLE IF NOT EXISTS user_chat(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    chat_id INT NOT NULL REFERENCES chats(id),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);