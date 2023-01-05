package application

import (
	"usersapi/internal/application/context"
	"usersapi/internal/application/server"
)

type App struct {
	ctx *context.Context
}

func NewApp(ctx *context.Context) *App {
	return &App{
		ctx: ctx,
	}
}

func (a *App) Run() {
	e := server.SetupServer(a.ctx)

	e.Logger.Fatal(e.Start(":2001"))
}
