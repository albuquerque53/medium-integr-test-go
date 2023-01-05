package user

import (
	"context"
)

type Repository interface {
	GetUsers(ctx context.Context) ([]*User, error)
	SaveUser(ctx context.Context, user *User) (int, error)
	GetUser(ctx context.Context, id int) (*User, error)
}
