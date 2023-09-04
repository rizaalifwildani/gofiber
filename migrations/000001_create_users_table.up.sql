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