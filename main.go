package main

import (
	"log"
	// Import package config
	"myapp/config"
	routes "myapp/routes/admin"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// func getTours(c *fiber.Ctx) error {
// 	var tours []models.TourView
// 	// khai báo slice : nhớ là muốn lưu thì phải dùng con trỏ

// 	config.DB.Model(&models.Tour{}).Select([]string{"id", "schedule", "title"}).Where("deleted = ? AND status = ?", false, "active").Find(&tours)
// 	return c.JSON(fiber.Map{
// 		"pageTitle": "Danh sách tours",
// 		"tours":     tours,
// 	})
// }

func main() {
	config.ConnectDB()

	// Lấy giá trị cổng từ biến môi trường
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	// Khởi tạo Fiber
	app := fiber.New()
	// app.Get("/tours", getTours)
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + port))
}
