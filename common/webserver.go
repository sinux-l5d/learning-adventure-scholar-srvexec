package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func execFactory(fn Handler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var j ToHandle

		if err := c.BodyParser(&j); err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
			// return err
		}

		status, out := fn(j)

		c.JSON(map[string]interface{}{
			"status": status.String(),
			"output": out,
		})
		return nil
	}
}

func Webserver(fn Handler) *fiber.App {
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: recover.ConfigDefault.StackTraceHandler,
	}))

	app.Use(logger.New())

	app.Post("/exec", execFactory(fn))

	return app
}
