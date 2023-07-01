package api

import (
	"golang-fiber/app/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PermissionRoute(router fiber.Router, db *gorm.DB) {
	var (
		permissionRepo repository.PermissionRepo = repository.NewPermissionRepo(db)
	)
	route := router.Group("/permission")
	route.Get("/", func(c *fiber.Ctx) error {
		data, _ := permissionRepo.GetAll()
		return c.JSON(data)
	})
}
