CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE compromised_passwords (
    id SERIAL PRIMARY KEY,
    hashed_password TEXT NOT NULL
);

CREATE INDEX hash_prefix_idx ON compromised_passwords (substring(hashed_password FROM 1 FOR 5));
