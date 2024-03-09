package postgres

import (
	"Amure/config"
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database interface {
	Connection() (*pgx.Conn, error)
	MigrationsUp(url ...string) error
	Close() error
}

type Postgres struct {
	Config config.PostgresQL
	client *pgx.Conn
	ctx    context.Context
	url    string
}

func NewPostgres(ctx context.Context, database config.PostgresQL) Database {
	return &Postgres{
		Config: database,
		ctx:    ctx,
	}
}

func (p *Postgres) Connection() (*pgx.Conn, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", p.Config.User, p.Config.Password, p.Config.Host, p.Config.Port, p.Config.DbName)
	client, err := pgx.Connect(p.ctx, url)
	if err != nil {
		return nil, err
	}

	p.client = client
	p.url = url

	return client, nil
}

func (p *Postgres) MigrationsUp(url ...string) error {
	var sourceURL string
	if url == nil {
		sourceURL = "file://postgres/migrations/up"
	} else {
		sourceURL = url[0]
	}
	m, err := migrate.New(sourceURL, p.url)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Close() error {
	err := p.client.Close(p.ctx)

	if err != nil {
		return err
	}
	return nil
}
