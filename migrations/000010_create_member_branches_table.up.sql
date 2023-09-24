-- Create the member_branches table
CREATE TABLE member_branches (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    member_id UUID REFERENCES members(id) ON DELETE CASCADE,
    branch_id UUID REFERENCES branches(id) ON DELETE CASCADE,
    status VARCHAR(10) DEFAULT 'pending' CHECK (status IN ('active', 'inactive', 'pending')),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

-- Create unique index on branch_id and member_id
CREATE UNIQUE INDEX idx_member_branches_unique_ids ON member_branches (member_id, branch_id);
CREATE INDEX idx_member_branches_status ON member_branches (status);
