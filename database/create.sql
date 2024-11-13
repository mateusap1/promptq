CREATE TABLE user (
    id int primary key,
    email varchar,
    password_hash varchar,
    salt bytea,
    email_verified boolean,
    confirm_token bytea,
    confirm_token_expires timestamp,
    reset_token bytea,
    reset_token_expires timestamp,
    created_at timestamp,
    updated_at timestamp,
);