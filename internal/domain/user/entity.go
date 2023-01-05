package user

import "context"

type Entity struct {
	r Repository
}

func NewUserEntity(r Repository) *Entity {
	return &Entity{r}
}

func (e *Entity) ListUsers() ([]*User, error) {
	ctx := context.Background()

	return e.r.GetUsers(ctx)
}

func (e *Entity) NewUser(user *User) (*User, error) {
	ctx := context.Background()

	id, err := e.r.SaveUser(ctx, user)

	if err != nil {
		return nil, err
	}

	u, err := e.r.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}

	return u, err
}
