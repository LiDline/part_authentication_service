DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id TEXT PRIMARY KEY,
  email TEXT NOT NULL
);

CREATE TABLE refresh_tokens (
  id SERIAL PRIMARY KEY,
  refresh_token TEXT,
  created_at BIGINT,
  ip TEXT NOT NULL,
  user_id TEXT NOT NULL,
  FOREIGN KEY (user_id)  REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO users (id, email)
VALUES ('57979158-bc47-490c-87fb-183d9b7a99d4', 
        'test.email@mail.com');


