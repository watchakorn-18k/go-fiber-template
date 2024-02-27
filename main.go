package main

import (
	"log"
	"os"
	"template/configuration"
	ds "template/domain/datasources"
	repo "template/domain/repositories"
	gw "template/src/gateways"
	"template/src/middlewares"
	sv "template/src/services"

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
	app.Use(recover.New())
	app.Use(cors.New())

	mongodb := ds.NewMongoDB(10)

	userMongo := repo.NewUsersRepository(mongodb)

	sv0 := sv.NewUsersService(userMongo)

	gw.NewHTTPGateway(app, sv0)

	PORT := os.Getenv("DB_PORT_LOGIN")

	if PORT == "" {
		PORT = "8080"
	}

	app.Listen(":" + PORT)
}
