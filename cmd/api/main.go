package main

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	godotenv.Load("config.env")

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
	}

	CreateS3Bucket()

	app := fiber.New()
	routes.InitRoutes(app, controllersInit)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}

func CreateS3Bucket() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err.Error())
	}

	s3Client := s3.New(sess)

	bucketName := "my-new-bucket-testt"

	_, err = s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err == nil {
		log.Println("Bucket already exists, skipping creation")
		return
	}

	_, err = s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("Failed to create bucket: %s", err.Error())
	}

	err = s3Client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Fatalf("Failed to wait for bucket to exist: %s", err.Error())
	}

	log.Println("Bucket created successfully")

}
