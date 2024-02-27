package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(ctx *fiber.App) {
	ctx.Use(logger.New(logger.Config{
		TimeFormat: "15:04:05",
		TimeZone:   "Asia/Bangkok",
	}))
}
