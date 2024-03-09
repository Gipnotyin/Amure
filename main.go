package main

import (
	"Amure/config"
	"Amure/postgres"
	"Amure/server"
	colors "Amure/utils"
	"context"
)

func main() {
	cfg := config.MustConfig()
	ctx := context.Background()

	colors.PrintLog(colors.ColorBlue, cfg)
	pg := postgres.NewPostgres(ctx, cfg.DBConfig)
	client, err := pg.Connection()
	if err != nil {
		colors.PrintLog(colors.ColorRed, "Failed connection to database")
		return
	}

	if err := pg.MigrationsUp(); err != nil && err.Error() != "no change" {
		colors.PrintLog(colors.ColorRed, "Failed migrations Up: ", err)
		return
	}
	colors.PrintLog(colors.ColorPurple, "Migrations update success!")

	runner := server.NewServer(cfg)
	if err := runner.Run(pg, client, ctx); err != nil {
		colors.PrintLog(colors.ColorRed, "Server is not run: ", err)
		return
	}
	colors.PrintLog(colors.ColorPurple, "Server is running complete success")
}
