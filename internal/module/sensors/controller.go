package sensors

import (
	"rest-api-controller/internal/module/sensors/service"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	service service.Service
}

func NewController(s service.Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) Register(app *fiber.App) {
	app.Post("/sensors", c.CreateSensors)
	app.Get("/sensors", c.GetAllSensors)
	app.Get("/sensors/:id", c.GetSensors)
	app.Patch("/sensors/:id", c.UpdateSensors)
	app.Delete("/sensors/:id", c.DeleteSensors)
}
