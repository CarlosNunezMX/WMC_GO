package main

import (
	"github.com/CarlosNunezMX/WMC_GO/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var port string = ":3000"

func main() {

	app := fiber.New()
	app.Name("WMC_GO")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Use(logger.New())

	router.Router(app)
	app.Handler()
	app.Listen(port)
}
