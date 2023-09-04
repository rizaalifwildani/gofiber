-- Create the user_roles table
CREATE TABLE user_roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

-- Create unique index on role_id and user_id
CREATE UNIQUE INDEX idx_user_roles_unique_ids ON user_roles (user_id, role_id);
