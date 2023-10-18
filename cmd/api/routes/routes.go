package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/general-market-GO/cmd/api/userctrll"
)

func InitRoutes(app *fiber.App, userController userctrll.Controller) {

	//User CRUD Routes
	//app.Post("/createUser", userController.CreateUser)
	app.Get("/getUserById/:userId", userController.GetUser)
	//app.Get("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	//app.Put("/updateUser/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	//app.Delete("/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	//Login Route
	//app.Post("/login", userController.LoginUser)
}
