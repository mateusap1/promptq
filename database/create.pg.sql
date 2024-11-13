CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR NOT NULL,
    password_hash VARCHAR NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE NOT NULL,
    validate_token VARCHAR NULL,
    validate_token_expires TIMESTAMP NULL,
    reset_token VARCHAR NULL,
    reset_token_expires TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);