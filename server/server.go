package server

import (
	"Amure/config"
	"Amure/postgres"
	"Amure/user"
	"context"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v5"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	config *config.Config
}

type Server interface {
	Run(psg postgres.Database, client *pgx.Conn, ctx context.Context) error
}

func NewServer(config *config.Config) Server {
	return &server{
		config: config,
	}
}

func (c *server) Run(psg postgres.Database, client *pgx.Conn, ctx context.Context) error {
	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/user")

	dbUser := user.NewUser(client, psg, ctx)

	routes := user.NewServices(api, dbUser, c.config)
	routes.RegisterRouter()

	return app.Listen(c.config.Address)
}
