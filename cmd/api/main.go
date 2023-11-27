package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/karilho/general-market-GO/adapters/repo/pg_repo"
	"github.com/karilho/general-market-GO/cmd/api/controllers"
	"github.com/karilho/general-market-GO/cmd/api/routes"
	"github.com/karilho/general-market-GO/domain/buyers"
	"github.com/karilho/general-market-GO/domain/users"
	"github.com/karilho/general-market-GO/migrationsSQL"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	dburl := os.Getenv("DATABASE_URL")

	pgrepo.MigrateDB(http.FS(migrationsSQL.MigrationsDir), dburl)

	repositories, err := pgrepo.New(ctx, dburl)
	if err != nil {
		log.Fatal(err)
	}

	usersService := users.NewUserService(repositories)
	buyerService := buyers.NewBuyerService(repositories)

	controllersInit := []controllers.Controller{
		controllers.NewUserController(usersService),
		controllers.NewBuyerController(buyerService),
		//mais1

	}

	app := fiber.New()
	routes.InitRoutes(app, controllersInit)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}

}
