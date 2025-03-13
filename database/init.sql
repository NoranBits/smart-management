-----------------------------
-- database initialization --
-----------------------------

-- Ensure the database exists (Optional, if running outside of Docker Compose)
-- CREATE DATABASE IF NOT EXISTS pgDatabase;

-- Use the target database
\c pgDatabase;

-- Create ENUM for user roles
DO $$ BEGIN
    CREATE TYPE user_role AS ENUM ('admin', 'manager', 'waiter', 'chef', 'customer');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;
-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL DEFAULT 'customer',
    active BOOLEAN DEFAULT TRUE NOT NULL
);
-- Index for optimized email lookups
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Trigger to auto-update on 'updated_at' field
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();

-----------------------------
-- Insert Default Users
-----------------------------

INSERT INTO users (name, email, password, role, active)
VALUES 
    ('Admin User', 'admin@restaurant.com', '$2a$10$hashedpasswordexample', 'admin', TRUE)
ON CONFLICT (email) DO NOTHING;

INSERT INTO users (name, email, password, role, active)
VALUES 
    ('Manager', 'manager@restaurant.com', '$2a$10$hashedpasswordexample', 'manager', TRUE)
ON CONFLICT (email) DO NOTHING;

INSERT INTO users (name, email, password, role, active)
VALUES 
    ('Waiter John', 'waiter@restaurant.com', '$2a$10$hashedpasswordexample', 'waiter', TRUE)
ON CONFLICT (email) DO NOTHING;

INSERT INTO users (name, email, password, role, active)
VALUES 
    ('Chef Mike', 'chef@restaurant.com', '$2a$10$hashedpasswordexample', 'chef', TRUE)
ON CONFLICT (email) DO NOTHING;

INSERT INTO users (name, email, password, role, active)
VALUES 
    ('Customer Alice', 'customer@restaurant.com', '$2a$10$hashedpasswordexample', 'customer', TRUE)
ON CONFLICT (email) DO NOTHING;


