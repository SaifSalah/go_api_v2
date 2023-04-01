CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS comments (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug text,
    author text,
    body text
);