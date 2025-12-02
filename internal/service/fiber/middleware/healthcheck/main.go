package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (c *Middleware) Register(app *fiber.App) {
	app.Use("/healthcheck", healthcheck.New(healthcheck.Config{
		LivenessProbe:    func(c *fiber.Ctx) bool { return true },
		LivenessEndpoint: "/live",
		ReadinessProbe: func(c *fiber.Ctx) bool {
			return true
		},
		ReadinessEndpoint: "/ready",
	}))
}
