package room

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

func (r *Repository) Create(ctx context.Context, room *Room) error {
	_, err := r.db.Exec(ctx, "INSERT INTO rooms (name, description, is_private, password, max_users, created_by) VALUES ($1, $2, $3, $4, $5, $6)", room.Name, room.Description, room.IsPrivate, room.Password, room.MaxUsers, room.CreatedBy)
	return err
}

func (r *Repository) FindAll(ctx context.Context) ([]*Room, error) {
	rows, err := r.db.Query(ctx,
		"SELECT rooms.id, rooms.name, rooms.description, rooms.is_private, rooms.password, rooms.max_users, rooms.created_by, rooms.created_at, rooms.updated_at FROM rooms LEFT JOIN room_users ON rooms.id = room_users.room_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rooms := []*Room{}

	for rows.Next() {
		var room Room
		err := rows.Scan(&room.ID, &room.Name, &room.Description, &room.IsPrivate, &room.Password, &room.MaxUsers, &room.CreatedBy, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*Room, error) {
	room := &Room{}
	err := r.db.QueryRow(ctx, "SELECT id, name, description, is_private, password, max_users, created_by, created_at, updated_at FROM rooms WHERE id = $1", id).Scan(&room.ID, &room.Name, &room.Description, &room.IsPrivate, &room.Password, &room.MaxUsers, &room.CreatedBy, &room.CreatedAt, &room.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID, createdBy uuid.UUID) error {
	_, err := r.db.Exec(ctx, "DELETE FROM rooms WHERE id = $1 AND created_by = $2", id, createdBy)
	return err
}
