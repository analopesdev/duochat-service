CREATE TABLE IF NOT EXISTS rooms (
  id         uuid PRIMARY KEY,
  name       VARCHAR(255) NOT NULL,
  description VARCHAR(255) NOT NULL,
  is_private BOOLEAN NOT NULL,
  password VARCHAR(255),
  max_users INT NOT NULL,
  created_by uuid NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);