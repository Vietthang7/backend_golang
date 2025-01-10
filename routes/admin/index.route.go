package routes

import "github.com/gofiber/fiber/v2"

// SetupRoutes thiết lập các route chính cho ứng dụng
func SetupRoutes(app *fiber.App) {
	// Gọi SetupTourRoutes để thêm các route liên quan đến tour
	SetupTourRoutes(app)
}
