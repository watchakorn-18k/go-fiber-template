package gateways

import "github.com/gofiber/fiber/v2"

func GatewayUsers(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/users")

	api.Post("/add_users", gateway.CreateNewUserAccount)
	api.Get("/users", gateway.GetAllUserData)
}
