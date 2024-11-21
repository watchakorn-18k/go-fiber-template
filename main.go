package main

import (
	"fmt"
	"go-fiber-template/src/configuration"
	ds "go-fiber-template/src/domain/datasources"
	repo "go-fiber-template/src/domain/repositories"
	"go-fiber-template/src/gateways"
	"go-fiber-template/src/infrastructure/httpclients"
	"go-fiber-template/src/middlewares"
	sv "go-fiber-template/src/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {
	app := fiber.New(configuration.NewFiberConfiguration())
	middlewares.Logger(app)
	app.Use(middlewares.ScalarMiddleware(middlewares.Config{
		PathURL:   "/api/docs",
		SpecURL:   "./src/docs/swagger.yaml",
		PageTitle: "Fiber API documentation",
	}))
	app.Use(recover.New())
	app.Use(cors.New())

	mongodb := ds.NewMongoDB(10)
	ipHC := httpclients.NewIPHttpClient()
	userMongo := repo.NewUsersRepository(mongodb)
	userSV := sv.NewUsersService(userMongo)
	ipSV := sv.NewIpService(ipHC)

	gateways.NewHTTPGateway(app, userSV, ipSV)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
