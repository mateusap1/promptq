CREATE TABLE user (
    id primary key,
    email varchar(256),
    password_hash char(32),
    salt bytea,
    email_verified boolean,
    confirm_token bytea,
    confirm_token_expires timestamp,
    reset_token bytea,
    reset_token_expires timestamp,
    created_at timestamp,
    updated_at timestamp,
)