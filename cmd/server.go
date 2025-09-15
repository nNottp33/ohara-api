package main

import (
	"strings"
	"time"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nNottp33/ohara-api/internal/adapters/http/middleware"
	"github.com/nNottp33/ohara-api/internal/config/env"
)

func main() {
	app := fiber.New()
	appEnv := env.Get[env.AppConfig]("App")
	isProd := appEnv.AppEnv == "production"

	app.Use(
		requestid.New(
			requestid.Config{
				Header:    "X-Custom-Header",
				Generator: middleware.GenerateRequestIdMiddleware,
			},
		),
	)

	app.Use(middleware.LoggerMiddleware)

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

	app.Use(helmet.New())

	app.Use(compress.New(compress.Config{}))

	app.Use(
		limiter.New(
			limiter.Config{
				Next: func(c *fiber.Ctx) bool {
					return c.IP() == "127.0.0.1"
				},
				Max:        20,
				Expiration: 30 * time.Second,
				LimitReached: func(c *fiber.Ctx) error {
					return c.Status(fiber.StatusTooManyRequests).JSON(
						fiber.Map{
							"error": "Rate limit exceeded for this user and IP",
						},
					)
				},
				// KeyGenerator:          func(c *fiber.Ctx) string {
				//     return c.Get("x-forwarded-for")
				// },
				// Storage: myCustomStorage{}, // Can use redis storage
			},
		),
	)

	if !isProd {
		app.Use(
			swagger.New(
				swagger.Config{
					BasePath: "/",
					FilePath: "./docs/swagger.json",
					Path:     "swagger",
					Title:    "Swagger API Docs",
				},
			),
		)
	}

	app.Get(
		"/", func(c *fiber.Ctx) error {
			return c.Status(200).JSON(
				fiber.Map{
					"message": "Hello, World!",
				},
			)
		},
	)

	err := app.Listen(":" + appEnv.AppPort)
	if err != nil {
		panic(err)
	}
}
