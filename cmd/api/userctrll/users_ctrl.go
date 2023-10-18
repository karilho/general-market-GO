package userctrll

import (
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

	// Again using intermediary structs (or a map) is useful for decoupling
	// the internal entities from what is exposed on the web:
	return ctx.JSON(map[string]interface{}{
		"id":   userID,
		"name": user.Name,
	})
}
