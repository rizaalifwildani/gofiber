-- Create the user_branches table
CREATE TABLE user_branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    branch_id UUID REFERENCES branches(id) ON DELETE CASCADE,
    status VARCHAR(10) DEFAULT 'pending' CHECK (status IN ('active', 'inactive', 'pending')),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

-- Create unique index on branch_id and user_id
CREATE UNIQUE INDEX idx_user_branches_unique_ids ON user_branches (user_id, branch_id);
CREATE INDEX idx_user_branches_status ON user_branches (status);
