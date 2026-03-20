-- Deleting in reverse order of foreign key constraints
DELETE FROM study_sessions;
DELETE FROM group_members;
DELETE FROM study_groups;
DELETE FROM subjects;
DELETE FROM users;