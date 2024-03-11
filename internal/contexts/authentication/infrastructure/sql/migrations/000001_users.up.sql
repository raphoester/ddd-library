CREATE TABLE users(
    id uuid PRIMARY KEY,
    email_address TEXT NOT NULL,
    hashed_password TEXT NOT NULL,
    salt TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    is_active BOOLEAN NOT NULL,
    role TEXT NOT NULL
)
