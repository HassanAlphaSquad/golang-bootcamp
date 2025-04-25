-- +goose UP
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    email TEXT NOT NULL UNIQUE
);

-- +goose DOWN
DROP TABLE users; 