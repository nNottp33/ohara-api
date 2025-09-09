package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nNottp33/ohara-api/internal/adapters/http/middleware"
	"github.com/nNottp33/ohara-api/internal/config/env"
)

func main() {
	app := fiber.New()
	appConfig := env.Get[env.AppConfig]("App")

	app.Use(middleware.LoggerMiddleware)

	isProd := appConfig.AppEnv == "production"
	origins := map[bool]string{
		true:  "http://localhost:3033",
		false: "*",
	}[isProd]

	app.Use(
		cors.New(
			cors.Config{
				AllowOrigins: origins,
				AllowHeaders: "Origin, Content-Type, Accept, Authorization",
				AllowMethods: strings.Join(
					[]string{
						fiber.MethodGet,
						fiber.MethodPost,
						fiber.MethodHead,
						fiber.MethodPut,
						fiber.MethodDelete,
						fiber.MethodPatch,
					}, ",",
				),
			},
		),
	)

	app.Get(
		"/", func(c *fiber.Ctx) error {
			return c.Status(200).JSON(
				fiber.Map{
					"message": "Hello, World!",
				},
			)
		},
	)

	err := app.Listen(":" + appConfig.AppPort)
	if err != nil {
		panic(err)
	}
}
