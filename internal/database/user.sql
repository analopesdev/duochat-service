CREATE TABLE IF NOT EXISTS users (
  id         uuid PRIMARY KEY,
  nickname   VARCHAR(255) UNIQUE NOT NULL,
  avatar     VARCHAR(255) NOT NULL,
  token_version INT NOT NULL DEFAULT 1,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);