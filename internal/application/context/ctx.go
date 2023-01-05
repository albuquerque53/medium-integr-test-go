package context

import (
	"database/sql"
	"usersapi/internal/domain/user"
	"usersapi/internal/infra/db"

	"github.com/labstack/echo/v4"
)

type ContextInterface interface {
	SetDBConnection(conn *sql.DB)
	GetEchoContext() *echo.Context
	GetDBConnection() *db.DBManager
}

type Context struct {
	echo.Context
	db *sql.DB
}

func (c *Context) SetDBConnection(conn *sql.DB) {
	c.db = conn
}

func (c *Context) GetEchoContext() *echo.Context {
	return &c.Context
}

func (c *Context) GetDBConnection() *db.DBManager {
	return &db.DBManager{
		DB: c.db,
	}
}

func (c *Context) GetDTO() (*user.User, error) {
	u := &user.User{}

	if err := c.Bind(u); err != nil {
		return nil, err
	}

	return u, nil
}
