	
CREATE TABLE IF NOT EXISTS room_users (
    id uuid PRIMARY KEY,
    room_id uuid NOT NULL,
    user_id uuid NOT NULL,
    is_admin BOOLEAN NOT NULL,
    joined_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);
  
  