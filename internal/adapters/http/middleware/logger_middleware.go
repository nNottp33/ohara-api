package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/nNottp33/ohara-api/pkg/utils"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	latency := float64(time.Since(start)) / float64(time.Millisecond)

	fmt.Printf(
		"%s | %s | %d | %.2f ms | %s | %s | %s | %s\n",
		carbon.Now().ToDateTimeString(),
		c.Locals("requestid"),
		c.Response().StatusCode(),
		latency,
		c.IP(),
		c.Method(),
		c.Path(),
		utils.ExtractBrowserUA(c.Get("User-Agent")),
	)

	return err
}
