package routes

import (
	"golang-fiber/routes/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Router struct {
	DB *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{DB: db}
}

func (r *Router) StaticFile(app *fiber.App) {
	// Serve local path
	app.Static("/images", "./public/images", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    false,
		Index:     "index.html",
		MaxAge:    3600,
	})
}

func (r *Router) MainRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
}

func (r *Router) ApiRouter(app *fiber.App) {
	apiRoute := app.Group("/api", logger.New())
	api.PermissionRoute(apiRoute, r.DB)
}
