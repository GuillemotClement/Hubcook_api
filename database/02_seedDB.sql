BEGIN;

INSERT INTO role (title) 
VALUES 
('admin'),
('user'),
('moderator'),
('writter');

INSERT INTO category (title)
VALUES
('fast food'),
('vegan'),
('entree'),
('desert'),
('soupe');

COMMIT;