package main

import (
	"go-fiber-template/src/configuration"
	ds "go-fiber-template/src/domain/datasources"
	repo "go-fiber-template/src/domain/repositories"
	"go-fiber-template/src/gateways"
	"go-fiber-template/src/middlewares"
	sv "go-fiber-template/src/services"
	"log"
	"os"

	swagger "github.com/gofiber/contrib/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {

	// // // remove this before deploy ###################
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// /// ############################################

	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use(swagger.New(swagger.Config{
		BasePath: "/api/",
		FilePath: "./src/docs/swagger.yaml",
		Path:     "docs",
	}))
	app.Use(recover.New())
	app.Use(cors.New())

	mongodb := ds.NewMongoDB(10)

	userMongo := repo.NewUsersRepository(mongodb)

	sv0 := sv.NewUsersService(userMongo)
	sv1 := sv.NewIpService()

	gateways.NewHTTPGateway(app, sv0, sv1)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app.Listen("localhost:" + PORT)
}
