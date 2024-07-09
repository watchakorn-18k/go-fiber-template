package gateways

import (
	service "go-fiber-template/src/services"

	"github.com/gofiber/fiber/v2"
)

type HTTPGateway struct {
	UserService service.IUsersService
	IPService   service.IIpService
}

func NewHTTPGateway(app *fiber.App, users service.IUsersService, ip service.IIpService) {
	gateway := &HTTPGateway{
		UserService: users,
		IPService:   ip,
	}

	RouteUsers(*gateway, app)
	RouteIP(*gateway, app)
}
