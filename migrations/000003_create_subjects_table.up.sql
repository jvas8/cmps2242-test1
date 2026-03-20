CREATE TABLE IF NOT EXISTS subjects (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    description text
);