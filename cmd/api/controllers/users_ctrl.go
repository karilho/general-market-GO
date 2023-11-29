package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/general-market-GO/domain"
	"github.com/karilho/general-market-GO/domain/users"
)

type UserController struct {
	usersService users.Service
}

func NewUserController(usersService users.Service) UserController {
	return UserController{
		usersService: usersService,
	}
}

func (c UserController) RegisterRoutes(app *fiber.App) {
	app.Post("/create", c.UpsertUser)
	app.Get("/getUser/:userId", c.GetUser)
	app.Get("/healthcheck", c.HealthCheck)
}

func (c UserController) UpsertUser(ctx *fiber.Ctx) error {
	var user struct {
		UserID int    `json:"user_id"`
		Name   string `json:"name"`
	}
	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return domain.BadRequestErr("unable to parse payload as JSON", map[string]interface{}{
			"payload": string(ctx.Body()),
			"error":   err.Error(),
		})
	}

	userID, err := c.usersService.UpsertUser(ctx.Context(), domain.User{
		ID:   user.UserID,
		Name: user.Name,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(map[string]interface{}{
		"status":  "success",
		"user_id": userID,
	})
}

func (c UserController) GetUser(ctx *fiber.Ctx) error {
	userID, err := ctx.ParamsInt("id")
	if err != nil {
		return domain.BadRequestErr("the input user id is not a valid integer", map[string]interface{}{
			"received_id": ctx.Params(":id"),
		})
	}

	user, err := c.usersService.GetUser(ctx.Context(), userID)
	if err != nil {
		return err
	}

	return ctx.JSON(map[string]interface{}{
		"id":   userID,
		"name": user.Name,
	})
}

func (c UserController) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]interface{}{
		"status": "ok",
	})
}
