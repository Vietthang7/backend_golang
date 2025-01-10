package routes

import (
	controllers "myapp/controllers/admin"

	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(app *fiber.App) {
	app.Get("api/admin/categories", controllers.GetCategory)
	app.Post("api/admin/create/category", controllers.CreateCategory)
}
