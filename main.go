package main

import (
	"log"
	"myapp/config"
	"myapp/models" // Import package config
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func getTours(c *fiber.Ctx) error {
	// Lấy danh sách các tour từ cơ sở dữ liệu với điều kiện "deleted = false" và "status = 'active'"
	var tours []models.Tour
	if err := config.DB.Where("deleted = ? AND status = ?", false, "active").Find(&tours).Error; err != nil {
		log.Println("Lỗi khi truy vấn cơ sở dữ liệu:", err)
		return c.Status(500).SendString("Lỗi khi truy vấn cơ sở dữ liệu")
	}

	// Trả về dữ liệu tour dưới dạng JSON
	return c.JSON(fiber.Map{
		"pageTitle": "Danh sách tour5",
		"tours":     tours,
	})
}

func main() {
	// Kết nối đến Database
	config.ConnectDB()

	// Lấy giá trị cổng từ biến môi trường
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	// Khởi tạo Fiber
	app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	fmt.Println("ok")
	// 	log.Println("ok")
	// 	return c.SendString("Hello, World!") // Trả về một thông báo cho client
	// })
	app.Get("/tours", getTours)

	// Lắng nghe kết nối tại cổng được chỉ định
	log.Fatal(app.Listen(":" + port))
}
