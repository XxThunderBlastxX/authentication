package controller

import "github.com/gofiber/fiber/v2"

func InitialController() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jsonData := fiber.Map{"Name": "Authentication", "Created By": "Koustav Mondal <ThunderBlast>", "Status": "Running", "Version": "0.0.1"}

		return ctx.Status(200).JSON(jsonData)
	}
}
