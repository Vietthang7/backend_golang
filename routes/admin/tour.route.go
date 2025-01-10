package routes

import (
	controllers "myapp/controllers/admin"

	"github.com/gofiber/fiber/v2"
)

// SetupTourRoutes thiết lập các route liên quan đến tour
func SetupTourRoutes(app *fiber.App) {
	// Route lấy danh sách các tour
	app.Get("/admin/tours", controllers.GetTours)
}
