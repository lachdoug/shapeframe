package gui

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run(port string) {
	// Create the router.
	router := fiber.New()

	// Enable verbse logging.
	router.Use(logger.New())

	// Recover from panics.
	router.Use(recover.New(recover.Config{EnableStackTrace: true}))

	// Setup the routes.
	// routes(router)
	router.Get("/api/topology", topologyHandler)

	// Start listening.
	router.Listen(fmt.Sprintf(":%s", port))
}

func topologyHandler(context *fiber.Ctx) (err error) {
	context.Status(fiber.StatusOK).JSON(map[string]any{
		"Result": map[string]string{},
	})
	return
}
