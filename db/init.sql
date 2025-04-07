CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE compromised_passwords (
    id SERIAL PRIMARY KEY,
    hashed_password TEXT NOT NULL
);
