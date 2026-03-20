-- Insert Users
INSERT INTO users (name, email) VALUES 
('Alice Johnson', 'alice@example.com'),
('Bob Smith', 'bob@example.com');

-- Insert Subjects
INSERT INTO subjects (name, description) VALUES 
('Computer Science', 'Algorithms, Data Structures, and Go Programming.'),
('Mathematics', 'Calculus and Linear Algebra.');

-- Insert Study Groups
INSERT INTO study_groups (name, description, creator_id, subject_id) VALUES 
('Go Lang Masters', 'Preparing for the backend finals.', 1, 1),
('Calc 101', 'Weekly derivatives practice.', 2, 2);

-- Insert Group Members
INSERT INTO group_members (user_id, group_id) VALUES 
(1, 1), (2, 1), (2, 2);

-- Insert Study Sessions
INSERT INTO study_sessions (group_id, title, session_date, location, notes) VALUES 
(1, 'API Design Review', '2026-03-25 18:00:00', 'Library Room A', 'Bring your laptops.'),
(2, 'Derivatives Quiz Prep', '2026-03-26 15:00:00', 'Student Union', 'Review chapter 4.');