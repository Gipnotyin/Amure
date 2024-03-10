package user

import (
	"Amure/config"
	"Amure/models"
	colors "Amure/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type services struct {
	app    fiber.Router
	db     User
	config *config.Config
}

type Services interface {
	RegisterRouter()
	Ping(ctx *fiber.Ctx) error
	CreateUser(ctx *fiber.Ctx) error
	SelectUser(ctx *fiber.Ctx) error
	CreateStatusCode(ctx *fiber.Ctx) error
	SelectUsers(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
}

func NewServices(app fiber.Router, db User, config *config.Config) Services {
	return &services{
		app:    app,
		db:     db,
		config: config,
	}
}

func (routes *services) RegisterRouter() {
	routes.app.Get("/ping", routes.Ping)
	routes.app.Post("/create", routes.CreateUser)
	routes.app.Get("/get", routes.SelectUser)
	routes.app.Get("/select", routes.SelectUsers)
	routes.app.Post("/update", routes.UpdateUser)
}

func (routes *services) Ping(ctx *fiber.Ctx) error {
	return ctx.SendString("ok")
}

func (routes *services) SelectUser(ctx *fiber.Ctx) error {
	request := new(models.SelectUserIn)
	response := new(models.SelectUserOut)
	err := ctx.BodyParser(request)
	if err != nil {
		response.Status = http.StatusBadRequest
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(response)
	}

	user, err := routes.db.SelectUser(request.Login)
	if user == nil {
		response.Status = http.StatusNoContent
		ctx.Status(http.StatusNoContent)
		return ctx.JSON(response)
	}

	if err != nil {
		response.Status = http.StatusBadRequest
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(response)
	}

	response.Status = http.StatusOK
	response.Data = user
	return ctx.JSON(response)
}

func (routes *services) SelectUsers(ctx *fiber.Ctx) error {
	request := new(models.SelectUsersIn)
	response := new(models.Response)

	err := ctx.BodyParser(request)
	if err != nil {
		response.Status = http.StatusForbidden
		response.Err = err.Error()
		ctx.Status(http.StatusForbidden)
		return ctx.JSON(response)
	}

	users, err := routes.db.SelectUsers(request.Sort, request.Field)
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Err = err.Error()
		ctx.Status(http.StatusBadRequest)
		return ctx.JSON(response)
	}

	ctx.Status(http.StatusOK)
	response.Status = http.StatusOK
	response.Data = users
	return ctx.JSON(response)
}

func (routes *services) CreateUser(ctx *fiber.Ctx) error {
	request := new(models.UserIn)
	err := ctx.BodyParser(request)

	if err != nil {
		return fmt.Errorf("can not parse body err: %w", err)
	}

	request.HashPassword, err = routes.createHashPassword(request.HashPassword)
	if err != nil {
		return err
	}

	userID, err := routes.db.InsertUser(request)
	if err != nil || userID == nil {
		return fmt.Errorf("error insert into user %v", err)
	}

	response := struct {
		Status int
		UserID *string
	}{
		Status: http.StatusOK,
		UserID: userID,
	}
	ctx.Status(http.StatusOK)
	colors.PrintLog(colors.ColorPurple, "User create is success! His user id is: ", *userID)
	return ctx.JSON(response)
}

func (routes *services) UpdateUser(ctx *fiber.Ctx) error {
	request := new(models.UserIn)
	response := new(models.Response)
	err := ctx.BodyParser(request)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		response.Status = http.StatusBadRequest
		response.Err = err.Error()
		return fmt.Errorf("can not parse body err: %w", err)
	}

	request.HashPassword, err = routes.createHashPassword(request.HashPassword)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Err = err.Error()
		return err
	}

	if err := routes.db.UpdateUser(request); err != nil {
		ctx.Status(http.StatusForbidden)
		response.Status = http.StatusForbidden
		response.Err = err.Error()
		return err
	}

	ctx.Status(http.StatusNoContent)
	response.Status = http.StatusNoContent
	return ctx.JSON(response)
}

func (routes *services) CreateStatusCode(ctx *fiber.Ctx) error {
	response := ctx.Response()
	colors.PrintLog(colors.ColorCyan, response.Body(), response.StatusCode())

	return ctx.JSON(response)
}
