package router

import (
	"github.com/CarlosNunezMX/WMC_GO/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	apiGroup := app.Group("/api")

	apiGroup.Get("/files", controllers.Find)
	apiGroup.Get("/file", controllers.Video)
}
