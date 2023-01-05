package main

import (
	"usersapi/internal/application"
	"usersapi/internal/application/context"
	"usersapi/internal/infra/db"
)

func main() {
	ctx := &context.Context{}
	ctx.SetDBConnection(db.ConectToDatabase())

	app := application.NewApp(ctx)
	app.Run()
}
