-- Create the member_occupations table
CREATE TABLE member_occupations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    member_id UUID REFERENCES members(id) ON DELETE CASCADE ON UPDATE CASCADE,
    company VARCHAR(50) NOT NULL,
    department VARCHAR(50) NOT NULL,
    address TEXT NOT NULL,
    postal_code VARCHAR(6),
    phone VARCHAR(15),
    fax VARCHAR(20),
    email VARCHAR(50),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

-- Create unique index on branch_id and member_id
CREATE UNIQUE INDEX idx_member_occupations_unique_ids ON member_occupations (member_id);
CREATE INDEX idx_member_occupations_company ON member_occupations(company);
CREATE INDEX idx_member_occupations_department ON member_occupations(department);
CREATE INDEX idx_member_occupations_address ON member_occupations(address);
CREATE INDEX idx_member_occupations_postal_code ON member_occupations(postal_code);
CREATE INDEX idx_member_occupations_phone ON member_occupations(phone);
CREATE INDEX idx_member_occupations_fax ON member_occupations(fax);
CREATE INDEX idx_member_occupations_email ON member_occupations(email);