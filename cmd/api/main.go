package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/general-market-GO/adapters/repo"
	"github.com/karilho/general-market-GO/adapters/repo/pg_repo"
	"github.com/karilho/general-market-GO/cmd/api/routes"
	"github.com/karilho/general-market-GO/cmd/api/userctrll"
	"github.com/karilho/general-market-GO/domain/users"
	"log"
)

func main() {
	ctx := context.Background()

	// init dependencies
	var usersRepo repo.Users
	//TODO -> Fix no nome do service, ficou @postgres por conta do meu service criado no k8s
	usersRepo, err := pgrepo.New(ctx, "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	usersService := users.NewService(usersRepo)
	userController := userctrll.NewController(usersService)

	// init API
	app := fiber.New()
	// init routes
	//Why not working when i put init routes on the same package as main? only when i make a new package?
	routes.InitRoutes(app, userController)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

}
