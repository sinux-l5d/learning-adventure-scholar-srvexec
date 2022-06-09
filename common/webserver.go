package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func execFactory(fn Executor) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var j ToExecute

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

func Webserver(fn Executor) *fiber.App {
	app := fiber.New()

	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: recover.ConfigDefault.StackTraceHandler,
	}))

	app.Post("/exec", execFactory(fn))

	return app
}
