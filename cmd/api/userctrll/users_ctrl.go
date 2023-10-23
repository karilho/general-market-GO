package userctrll

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/general-market-GO/domain"
	"github.com/karilho/general-market-GO/domain/users"
)

// direct injection of users service
type Controller struct {
	usersService users.Service
}

func NewController(usersService users.Service) Controller {
	return Controller{
		usersService: usersService,
	}
}

func (c Controller) UpsertUser(ctx *fiber.Ctx) error {
	// Intermediary structure so I don't expose my internal
	// user representation to the outside world:
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
		// in this case the internal name for the ID attribute is just ID not `UserID`:
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

func (c Controller) GetUser(ctx *fiber.Ctx) error {
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
