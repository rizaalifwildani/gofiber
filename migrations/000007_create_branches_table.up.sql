-- Create the branches table
CREATE TABLE branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(35) UNIQUE NOT NULL,
    code VARCHAR(15) UNIQUE NOT NULL,
    address TEXT,
    description TEXT,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_branches_name ON branches(name);
CREATE UNIQUE INDEX idx_branches_code ON branches(code);