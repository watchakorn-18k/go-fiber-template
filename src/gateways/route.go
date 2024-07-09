package gateways

import "github.com/gofiber/fiber/v2"

func RouteUsers(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/user")

	api.Post("/add_user", gateway.CreateNewUserAccount)
	api.Get("/users", gateway.GetAllUserData)
	api.Put("/update_user", gateway.UpdateUserData)
	api.Delete("/delete_user/:user_id", gateway.DeleteUser)
	api.Get("/get_user", gateway.GetUser)
}

func RouteIP(gateway HTTPGateway, app *fiber.App) {
	api := app.Group("/api/ip")
	api.Get("/check_ip", gateway.GetIp)
}
