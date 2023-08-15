package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type Ninja struct {
	Name   string
	Weapon string
}

func getNinja(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

var ninja Ninja

func createNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	ninja = Ninja{
		Name:   body.Name,
		Weapon: body.Weapon,
	}
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func FiberFunc() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})

	ninjaApp := app.Group("/ninja")
	ninjaApp.Get("", getNinja)
	ninjaApp.Post("", createNinja)
	app.Listen(":80")
}
