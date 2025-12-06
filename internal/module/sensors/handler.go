package sensors

import (
	"fmt"
	"rest-api-controller/internal/module/sensors/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAllTempSensors(ctx *fiber.Ctx) error {
	item, err := c.service.GetAllTemperature()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) GetTempSensors(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id", "")
	if itemID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID parameter is required"})
	}
	item, err := c.service.GetTemperatureItemByID(itemID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) CreateTempSensors(ctx *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	body := new(model.CreateTemperatureSensorsDTO)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}

	err := validate.Struct(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}
	// робота із даними
	item, err := c.service.CreateTemperatureItem(body)
	if err != nil {
		return err
	}
	// обробка відповіді
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) UpdateTempSensors(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id", "")
	if itemID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID parameter is required"})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	body := new(model.UpdateTemperatureSensorsDTO)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}

	err := validate.Struct(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}
	// робота із даними
	item, err := c.service.UpdateTemperatureItem(itemID, body)
	if err != nil {
		return err
	}
	// обробка відповіді
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) DeleteTempSensors(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id", "")
	if itemID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID parameter is required"})
	}
	err := c.service.DeleteTemperatureItem(itemID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item deleted successfully"})

}

func (c *Controller) GetAllHumiditySensors(ctx *fiber.Ctx) error {
	item, err := c.service.GetAllHumidity()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) GetHumiditySensors(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id", "")
	if itemID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID parameter is required"})
	}
	item, err := c.service.GetHumidityItemByID(itemID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) CreateHumiditySensors(ctx *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	body := new(model.CreateHumiditySensorsDTO)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}

	err := validate.Struct(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}
	// робота із даними
	item, err := c.service.CreateHumidityItem(body)
	if err != nil {
		return err
	}
	// обробка відповіді
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) UpdateHumiditySensors(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id", "")
	if itemID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID parameter is required"})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	body := new(model.UpdateHumiditySensorsDTO)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}

	err := validate.Struct(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}
	// робота із даними
	item, err := c.service.UpdateHumidityItem(itemID, body)
	if err != nil {
		return err
	}
	// обробка відповіді
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) DeleteHumiditySensors(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id", "")
	if itemID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID parameter is required"})
	}
	err := c.service.DeleteHumidityItem(itemID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item deleted successfully"})

}
