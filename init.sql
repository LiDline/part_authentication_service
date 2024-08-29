DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
  guuid UUID PRIMARY KEY,
  password TEXT NOT NULL
);

CREATE TABLE refresh_tokens (
  refresh_token TEXT PRIMARY KEY,
  created_at TIMESTAMP,
  guuid UUID NOT NULL,
  FOREIGN KEY (guuid)  REFERENCES users (guuid) ON DELETE CASCADE
);

INSERT INTO users (guuid, password)
VALUES ('57979158-bc47-490c-87fb-183d9b7a99d4', 'YWRtaW4=');


