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
	app.Post("/sensors/temperature", c.CreateTempSensors)
	app.Get("/sensors/temperature", c.GetAllTempSensors)
	app.Get("/sensors/temperature/:id", c.GetTempSensors)
	app.Patch("/sensors/temperature/:id", c.UpdateTempSensors)
	app.Delete("/sensors/temperature/:id", c.DeleteTempSensors)

	app.Post("/sensors/humidity", c.CreateHumiditySensors)
	app.Get("/sensors/humidity", c.GetAllHumiditySensors)
	app.Get("/sensors/humidity/:id", c.GetHumiditySensors)
	app.Patch("/sensors/humidity/:id", c.UpdateHumiditySensors)
	app.Delete("/sensors/humidity/:id", c.DeleteHumiditySensors)
}
