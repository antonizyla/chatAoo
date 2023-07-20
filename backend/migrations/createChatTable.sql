CREATE TABLE IF NOT EXISTS chat (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        description TEXT NOT NULL,
        name TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

