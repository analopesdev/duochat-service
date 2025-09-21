package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, u *User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO users (id, nickname, avatar, token_version, updated_at) VALUES ($1, $2, $3, $4, $5)", u.ID, u.Nickname, u.Avatar, u.TokenVersion, u.UpdatedAt)
	return err
}

func (r *Repository) FindAll(ctx context.Context) ([]*User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, nickname, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Nickname, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (*User, error) {
	u := &User{}
	err := r.db.QueryRow(ctx, "SELECT id, nickname, updated_at FROM users WHERE id = $1", id).Scan(&u.ID, &u.Nickname, &u.UpdatedAt)
	return u, err
}

func (r *Repository) GetByNickname(ctx context.Context, nickname string) (*User, error) {
	const q = `
		SELECT id, nickname, avatar, token_version, updated_at
		FROM users
		WHERE nickname = $1
	`
	u := new(User)
	if err := r.db.
		QueryRow(ctx, q, nickname).
		Scan(&u.ID, &u.Nickname, &u.Avatar, &u.TokenVersion, &u.UpdatedAt); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return u, nil
}
