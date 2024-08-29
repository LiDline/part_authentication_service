DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id TEXT PRIMARY KEY,
  password TEXT NOT NULL
);

CREATE TABLE refresh_tokens (
  refresh_token TEXT PRIMARY KEY,
  created_at TIMESTAMP,
  id TEXT NOT NULL,
  FOREIGN KEY (id)  REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO users (id, password)
VALUES ('57979158-bc47-490c-87fb-183d9b7a99d4', '$2a$10$0MMD66Q0EPVJLzJ8G04apuKYa/aATC/73t4K5G1wjkdNzrYPnvrKa');


