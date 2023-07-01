package main

import (
	"golang-fiber/config"
	"golang-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB = config.SetupDatabaseConnection()
	appRouter          = routes.NewRouter(db)
)

func main() {
	appConfig := config.FiberConfig()

	app := fiber.New(appConfig)

	defer config.CloseDatabaseConnection(db)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	appRouter.StaticFile(app)
	appRouter.MainRoutes(app)
	appRouter.ApiRouter(app)

	err := app.Listen(":8000")
	if err != nil {
		panic(err.Error())
	}
}
