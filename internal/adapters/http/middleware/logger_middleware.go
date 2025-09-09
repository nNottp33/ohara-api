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
	latency := time.Since(start)

	now := carbon.Now().ToDateTimeString()

	fmt.Printf(
		"%s | %d | %.3f ms | %s | %s | %s | %s\n",
		now,
		c.Response().StatusCode(),
		float64(latency)/float64(time.Millisecond),
		c.IP(),
		c.Method(),
		c.Path(),
		utils.ExtractBrowserUA(c.Get("User-Agent")),
	)

	return err
}
