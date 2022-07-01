package common

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func execFactory(fn Handler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var j ToHandle

		if err := c.BodyParser(&j); err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"error": err.Error()})
			// return err
		}

		out, status := fn(j)

		c.Status(status.HttpCode())

		c.JSON(map[string]interface{}{
			"status": status.String(),
			"output": out,
		})
		return nil
	}
}

func Webserver(fn Handler) *fiber.App {
	app := fiber.New()

	// recover from errors
	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: recover.ConfigDefault.StackTraceHandler,
	}))

	// Logger
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		c.Next()
		LogInfo("status=%d time=%s ip=%s method=%s path=%s", c.Response().StatusCode(), time.Since(start).Round(time.Millisecond), c.IP(), c.Method(), c.Path())
		return nil
	})

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// Execute the code
	app.Post("/exec", execFactory(fn))

	return app
}
