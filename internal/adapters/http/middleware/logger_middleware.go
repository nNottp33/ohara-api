package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	latency := time.Since(start)
	ms := float64(latency) / float64(time.Millisecond)

	fmt.Printf(
		"%s | %d | %.3f ms | %s | %s | %s | %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		c.Response().StatusCode(),
		ms,
		c.IP(),
		c.Method(),
		c.Path(),
		c.Get("User-Agent"),
	)
	return err
}
