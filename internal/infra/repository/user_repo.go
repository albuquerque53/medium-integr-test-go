package repository

import (
	"context"
	"usersapi/internal/domain/user"
	"usersapi/internal/infra/db"
)

type UserRepository struct {
	db *db.DBManager
}

func NewRepository(m *db.DBManager) *UserRepository {
	return &UserRepository{
		db: m,
	}
}

func (r *UserRepository) GetUser(ctx context.Context, id int) (*user.User, error) {
	row, err := r.db.QueryRow(ctx, "SELECT id, name, email, created_at, updated_at FROM users WHERE id=?", id)

	if err != nil {
		return nil, err
	}

	u := &user.User{}
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]*user.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, name, email, created_at, updated_at FROM users;")

	if err != nil {
		return nil, err
	}

	users := make([]*user.User, 0)

	for rows.Next() {
		user := &user.User{}

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) SaveUser(ctx context.Context, user *user.User) (int, error) {
	inserted, err := r.db.DB.ExecContext(ctx, "INSERT INTO users(name, email) VALUES(?, ?)", user.Name, user.Email)

	if err != nil {
		return 0, err
	}

	id, err := inserted.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), err
}
