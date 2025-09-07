package room_user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, roomUser *RoomUser) error {
	_, err := r.db.Exec(ctx, "INSERT INTO room_users (room_id, user_id, is_admin) VALUES ($1, $2, $3)", roomUser.RoomID, roomUser.UserID, roomUser.IsAdmin)
	return err
}

func (r *Repository) FindAll(ctx context.Context) ([]*RoomUser, error) {
	rows, err := r.db.Query(ctx, "SELECT id, room_id, user_id, is_admin, joined_at, is_active FROM room_users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roomUsers := []*RoomUser{}

	for rows.Next() {
		var roomUser RoomUser
		err := rows.Scan(&roomUser.ID, &roomUser.RoomID, &roomUser.UserID, &roomUser.IsAdmin, &roomUser.JoinedAt, &roomUser.IsActive)
		if err != nil {
			return nil, err
		}
		roomUsers = append(roomUsers, &roomUser)
	}
	return roomUsers, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, "DELETE FROM room_users WHERE id = $1", id)
	return err
}
