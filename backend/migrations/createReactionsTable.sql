CREATE TABLE IF NOT EXISTS reactions (
    reaction_id SERIAL PRIMARY KEY, 
    message_id UUID NOT NULL references messages(id) ON DELETE CASCADE,
    user_id UUID NOT NULL references users(id) ON DELETE CASCADE,
    emoji TEXT NOT NULL
)
