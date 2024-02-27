package gateways

import (
	service "template/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	UserService service.IUsersService
}

func NewHTTPGateway(app *fiber.App, users service.IUsersService) {
	gateway := &HTTPGateway{
		UserService: users,
	}

	GatewayUsers(*gateway, app)
}
