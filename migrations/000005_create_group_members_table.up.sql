CREATE TABLE IF NOT EXISTS group_members (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    group_id bigint NOT NULL REFERENCES study_groups(id) ON DELETE CASCADE,
    joined_at timestamp(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);