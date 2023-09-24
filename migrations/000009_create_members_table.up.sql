-- Create the members table
CREATE TABLE members (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    phone VARCHAR(15) UNIQUE NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    first_name VARCHAR(20),
    last_name VARCHAR(30),
    identity_number VARCHAR(15) UNIQUE NOT NULL,
    place_of_birth VARCHAR(30) NOT NULL,
    birthdate DATE NOT NULL,
    gender VARCHAR(10) NOT NULL CHECK (gender IN ('male', 'female')),
    nationality VARCHAR(30) NOT NULL,
    address TEXT NOT NULL,
    postal_code VARCHAR(6) NOT NULL,
    home_phone VARCHAR(15) NOT NULL,
    office_phone VARCHAR(15) NOT NULL,
    education TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);

CREATE UNIQUE INDEX idx_members_phone ON members(phone);
CREATE UNIQUE INDEX idx_members_email ON members(email);
CREATE INDEX idx_members_first_name ON members(first_name);
CREATE INDEX idx_members_last_name ON members(last_name);
CREATE UNIQUE INDEX idx_members_identity_number ON members (identity_number);
CREATE INDEX idx_members_place_of_birth ON members (place_of_birth);
CREATE INDEX idx_members_birthdate ON members (birthdate);
CREATE INDEX idx_members_gender ON members (gender);
CREATE INDEX idx_members_nationality ON members (nationality);
CREATE INDEX idx_members_address ON members (address);
CREATE INDEX idx_members_postal_code ON members (postal_code);
CREATE INDEX idx_members_home_phone ON members (home_phone);
CREATE INDEX idx_members_office_phone ON members (office_phone);
CREATE INDEX idx_members_education ON members (education);
