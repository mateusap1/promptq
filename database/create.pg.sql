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
    updated_at TIMESTAMP
);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users ON DELETE CASCADE,
    user_agent VARCHAR NOT NULL,
    ip_address VARCHAR NOT NULL,
    session_token VARCHAR NOT NULL,
    active BOOLEAN DEFAULT TRUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    expires_at TIMESTAMP
);

CREATE TABLE threads (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users ON DELETE CASCADE,
    tid VARCHAR NOT NULL,
    tname VARCHAR NOT NULL,
    deleted BOOLEAN DEFAULT false NOT NULL,
    deleted_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE prompts (
    id SERIAL PRIMARY KEY,
    thread_id INT NOT NULL REFERENCES threads ON DELETE CASCADE,
    ai BOOLEAN DEFAULT false NOT NULL,
    content VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);