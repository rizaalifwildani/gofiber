-- Create the user_auths table
CREATE TABLE user_auths (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    password TEXT,
    token TEXT,
    expired_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

-- Create unique index on user_id
CREATE UNIQUE INDEX idx_user_auths_unique_user_id ON user_auths (user_id);
