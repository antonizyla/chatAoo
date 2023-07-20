CREATE TABLE IF NOT EXISTS users_chat (
        chat_id UUID REFERENCES chat(id),
        user_id UUID REFERENCES users(id),
        PRIMARY KEY (chat_id, user_id)
    );

