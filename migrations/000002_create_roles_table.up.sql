-- Create the roles table
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(25) UNIQUE NOT NULL,
    display_name VARCHAR(35) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);
CREATE UNIQUE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_roles_display_name ON roles(display_name);