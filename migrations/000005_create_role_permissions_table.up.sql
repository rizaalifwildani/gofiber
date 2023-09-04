-- Create the role_permissions table
CREATE TABLE role_permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

-- Create unique index on role_id and permission_id
CREATE UNIQUE INDEX idx_role_permissions_unique_ids ON role_permissions (role_id, permission_id);
