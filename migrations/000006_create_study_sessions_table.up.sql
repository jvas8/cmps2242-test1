CREATE TABLE IF NOT EXISTS study_sessions (
    id bigserial PRIMARY KEY,
    group_id bigint NOT NULL REFERENCES study_groups(id) ON DELETE CASCADE,
    title varchar(255) NOT NULL,
    session_date timestamp(0) WITH TIME ZONE NOT NULL,
    location varchar(255),
    notes text,
    created_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);