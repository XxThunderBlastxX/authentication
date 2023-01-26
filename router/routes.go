package router

import (
	"github.com/XxThuderBlastxX/authentication/controller"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controller.InitialController())
}
