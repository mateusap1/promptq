CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR NOT NULL,
    password_hash VARCHAR NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE NOT NULL,
    validate_token VARCHAR NULL,
    validate_token_expires TIMESTAMP NULL,
    reset_token VARCHAR NULL,
    reset_token_expires TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT NOW,
    updated_at TIMESTAMP
);

CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INT NOT NULL REFERENCES users ON DELETE CASCADE,
    user_agent VARCHAR NOT NULL,
    ip_address VARCHAR NOT NULL,
    session_token VARCHAR NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP
);