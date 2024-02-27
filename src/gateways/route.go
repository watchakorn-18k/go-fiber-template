package gateways

import "github.com/gofiber/fiber/v2"

func GatewayUsers(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/users")

	api.Post("/users/", gateway.CreatNewUserAccount)
	api.Get("/all_users", gateway.GetAllUserData)
}
