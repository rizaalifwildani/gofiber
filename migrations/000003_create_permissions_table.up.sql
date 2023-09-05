-- Create the permissions table
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(25) UNIQUE NOT NULL,
    display_name VARCHAR(35) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);
CREATE UNIQUE INDEX idx_permissions_name ON permissions(name);
CREATE INDEX idx_permissions_display_name ON permissions(display_name);