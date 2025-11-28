package sensors

import (
	"fmt"
	"rest-api-controller/internal/module/sensors/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAllSensors(ctx *fiber.Ctx) error {
	item, err := c.service.GetAll()
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) GetSensors(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id", "")
	if itemID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID parameter is required"})
	}
	item, err := c.service.GetItemByID(itemID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) CreateSensors(ctx *fiber.Ctx) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	body := new(model.CreateTemperatureSensorsDTO)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}

	err := validate.Struct(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid JSON body: %s", err.Error())})
	}
	data := body.ToModel()
	// робота із даними
	item, err := c.service.CreateItem(&data)
	if err != nil {
		return err
	}
	// обробка відповіді
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) UpdateSensors(ctx *fiber.Ctx) error {
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
	data := body.ToModel(itemID)
	// робота із даними
	item, err := c.service.UpdateItem(&data)
	if err != nil {
		return err
	}
	// обробка відповіді
	return ctx.Status(fiber.StatusOK).JSON(item)
}

func (c *Controller) DeleteSensors(ctx *fiber.Ctx) error {
	itemID := ctx.Params("id", "")
	if itemID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID parameter is required"})
	}
	err := c.service.DeleteItem(itemID)
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Item deleted successfully"})

}
