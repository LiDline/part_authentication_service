DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id TEXT PRIMARY KEY,
  email TEXT NOT NULL
);

CREATE TABLE refresh_tokens (
  refresh_token TEXT PRIMARY KEY,
  created_at TIMESTAMP,
  ip TEXT NOT NULL,
  id TEXT NOT NULL,
  FOREIGN KEY (id)  REFERENCES users (id) ON DELETE CASCADE
);

INSERT INTO users (id, email)
VALUES ('57979158-bc47-490c-87fb-183d9b7a99d4', 
        'test.email@mail.com');


