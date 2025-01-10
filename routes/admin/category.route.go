package routes

import (
	controllers "myapp/controllers/admin"

	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(app *fiber.App) {
	app.Get("/admin/categories", controllers.GetCategory)
}
