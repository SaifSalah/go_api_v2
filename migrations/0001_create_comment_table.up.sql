CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS comments (
    id uuid,
    slug text,
    author text,
    body text
);