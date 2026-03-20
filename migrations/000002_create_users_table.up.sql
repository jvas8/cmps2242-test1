CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);