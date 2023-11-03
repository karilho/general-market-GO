package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/karilho/general-market-GO/adapters/repo"
	"github.com/karilho/general-market-GO/adapters/repo/pg_repo"
	"github.com/karilho/general-market-GO/cmd/api/routes"
	"github.com/karilho/general-market-GO/cmd/api/userctrll"
	"github.com/karilho/general-market-GO/domain/users"
	"github.com/karilho/general-market-GO/migrationsSQL"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	dburl := os.Getenv("DATABASE_URL")
	// init dependencies
	var usersRepo repo.Users

	//migrate database
	pgrepo.MigrateDB(http.FS(migrationsSQL.MigrationsDir), dburl)

	//TODO -> Fix no nome do service, ficou @postgres por conta do meu service criado no k8s
	usersRepo, err := pgrepo.New(ctx, dburl)
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
