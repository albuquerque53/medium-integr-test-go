package server

import (
	"net/http"
	"usersapi/internal/application/context"
	"usersapi/internal/application/writer"
	"usersapi/internal/domain/user"
	"usersapi/internal/infra/repository"

	"github.com/labstack/echo/v4"
)

func SetupServer(ctx *context.Context) *echo.Echo {
	e := echo.New()

	// Rota de listagem de usuários
	e.GET("/list", func(c echo.Context) error {
		ctx := c.(*context.Context)

		r := repository.NewRepository(ctx.GetDBConnection())

		u, err := user.NewUserEntity(r).ListUsers()

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.String(http.StatusOK, writer.ToJSON(u))
	})

	// Rota de criação de usuários
	e.POST("/new", func(c echo.Context) error {
		ctx := c.(*context.Context)

		u, _ := ctx.GetDTO()

		r := repository.NewRepository(ctx.GetDBConnection())

		created, err := user.NewUserEntity(r).NewUser(u)

		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.String(http.StatusCreated, writer.ToJSON(created))
	})

	// Middleware para injeção de dependências via Context
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx.Context = c
			return next(ctx)
		}
	})

	return e
}
