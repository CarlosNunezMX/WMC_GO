package router

import (
	"github.com/CarlosNunezMX/WMC_GO/addons"
	"github.com/CarlosNunezMX/WMC_GO/controllers"
	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	apiGroup := app.Group("/api")

	apiGroup.Get("/files", controllers.Find)
	apiGroup.Get("/file", controllers.Video)

	addonsGroup := app.Group("/addons")

	addonsGroup.Get("/online", addons.GetSource)
	addonsGroup.Post("/online/temp", addons.TempStoreOnlineMedia)
	addonsGroup.Get("/online/temp", addons.GetTemps)
	addonsGroup.Get("/proxy", addons.Proxy)

}
