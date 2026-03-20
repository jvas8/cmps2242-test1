CREATE TABLE IF NOT EXISTS study_groups (
    id bigserial PRIMARY KEY,
    name varchar(255) NOT NULL,
    description text,
    creator_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    subject_id bigint NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);