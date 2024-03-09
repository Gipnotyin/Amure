package user

import (
	"Amure/models"
	"Amure/postgres"
	colors "Amure/utils"
	"context"
	"github.com/jackc/pgx/v5"
)

type user struct {
	client *pgx.Conn
	db     postgres.Database
	ctx    context.Context
}

type User interface {
	SelectUser(login string) (*models.UserOut, error)
	InsertUser(*models.UserIn) (*string, error)
}

func NewUser(cl *pgx.Conn, db postgres.Database, ctx context.Context) User {
	return &user{
		client: cl,
		db:     db,
		ctx:    ctx,
	}
}

func (p *user) SelectUser(login string) (*models.UserOut, error) {
	args := pgx.NamedArgs{
		"login": login,
	}

	row, err := p.client.Query(p.ctx, SelectUser, args)
	if err != nil {
		colors.PrintLog(colors.ColorRed, "query select doesn't complete")
		return nil, err
	}

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.UserOut])
	if err != nil {
		colors.PrintLog(colors.ColorRed, "CollectOneRow doesn't complete")
		return nil, err
	}

	colors.PrintLog(colors.ColorGreen, "Select user is success")
	return &user, nil
}

func (p *user) InsertUser(user *models.UserIn) (*string, error) {
	var userID *string

	tr, err := p.client.Begin(p.ctx)
	if err != nil {
		colors.PrintLog(colors.ColorRed, "client doesn't create")
		return nil, err
	}
	defer func() {
		err = tr.Commit(p.ctx)
	}()

	args := pgx.NamedArgs{
		"name":          user.Name,
		"last_name":     user.LastName,
		"email":         user.Email,
		"login":         user.Login,
		"hash_password": user.HashPassword,
		"phone":         user.Phone,
	}

	err = tr.QueryRow(p.ctx, InsertUser, args).Scan(&userID)
	if err != nil {
		colors.PrintLog(colors.ColorRed, "Query row doesn't complete")
		return nil, err
	}

	if err := tr.Commit(p.ctx); err != nil {
		colors.PrintLog(colors.ColorRed, "Commit is not success")
		return nil, err
	}

	colors.PrintLog(colors.ColorGreen, "Insert user is success")
	return userID, nil
}
