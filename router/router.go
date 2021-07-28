package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/phatthakarn-ji/website-sec-api/handler"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api/v1", logger.New())
	api.Get("/", handler.Hello)
	api.Get("/contact", handler.GetContact)
	api.Put("/contact", handler.UpdateContact)

}
