-- Create the users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(25) UNIQUE NOT NULL,
    phone VARCHAR(15),
    email VARCHAR(50),
    first_name VARCHAR(20),
    last_name VARCHAR(30),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP
);
CREATE UNIQUE INDEX idx_users_username ON users(username);
CREATE INDEX idx_phone ON users(phone);
CREATE INDEX idx_users_searchs ON users(email, first_name, last_name);