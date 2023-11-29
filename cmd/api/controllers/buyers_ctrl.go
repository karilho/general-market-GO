package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/general-market-GO/domain"
	"github.com/karilho/general-market-GO/domain/buyers"
	"time"
)

type BuyerController struct {
	buyerService buyers.Service
}

func NewBuyerController(buyerService buyers.Service) BuyerController {
	return BuyerController{
		buyerService: buyerService,
	}
}

func (c BuyerController) RegisterRoutes(app *fiber.App) {
	app.Post("/createBuyer", c.UpsertBuyer)
	app.Get("/getBuyer/:buyer_id", c.GetBuyer)
}

func (c BuyerController) UpsertBuyer(ctx *fiber.Ctx) error {

	var inputUserData struct {
		CurrentType   string `json:"current_type"`
		Username      string `json:"username"`
		Email         string `json:"email"`
		PasswordHash  string `json:"password_hash"`
		FullName      string `json:"full_name"`
		PhoneNumber   string `json:"phone_number"`
		StreetAddress string `json:"street_address"`
		PlaceNumber   string `json:"place_number"`
		City          string `json:"city"`
		StateProvince string `json:"state_province"`
		PostalCode    string `json:"postal_code"`
	}

	err := json.Unmarshal(ctx.Body(), &inputUserData)
	if err != nil {
		return domain.BadRequestErr("unable to parse payload as JSON", map[string]interface{}{
			"payload": string(ctx.Body()),
			"error":   err.Error(),
		})
	}

	UserDataID, err := c.buyerService.UpsertUserData(ctx.Context(), domain.UserData{
		CurrentType:      inputUserData.CurrentType,
		Username:         inputUserData.Username,
		Email:            inputUserData.Email,
		PasswordHash:     inputUserData.PasswordHash,
		FullName:         inputUserData.FullName,
		PhoneNumber:      inputUserData.PhoneNumber,
		RegistrationDate: time.Now(),
		StreetAddress:    inputUserData.StreetAddress,
		PlaceNumber:      inputUserData.PlaceNumber,
		City:             inputUserData.City,
		StateProvince:    inputUserData.StateProvince,
		PostalCode:       inputUserData.PostalCode,
	})
	if err != nil {
		return err
	}

	BuyerId, err := c.buyerService.UpsertBuyer(ctx.Context(), domain.Buyers{
		UserDataID:   UserDataID,
		HasPurchased: false,
	})
	if err != nil {
		return err
	}

	return ctx.JSON(map[string]interface{}{
		"status":   "success",
		"buyer_id": BuyerId,
	})
}

func (c BuyerController) GetBuyer(ctx *fiber.Ctx) error {
	buyerID, err := ctx.ParamsInt("buyer_id")
	if err != nil {
		return domain.BadRequestErr("the input buyer id is not a valid integer", map[string]interface{}{
			"received_id": ctx.Params(":buyer_id"),
		})
	}

	buyer, err := c.buyerService.GetBuyer(ctx.Context(), buyerID)
	if err != nil {
		return err
	}

	return ctx.JSON(map[string]interface{}{
		"status": "success",
		"buyer":  buyer,
	})
}
