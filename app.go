package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Lấy giá trị từ biến môi trường
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")

	// Chuỗi DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Kết nối Database
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối với cơ sở dữ liệu: ", err)
	}
	fmt.Println("Kết nối cơ sở dữ liệu thành công.")
}
func main() {
	connectDB()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("ok")
		return c.SendString("Hello, World!") // Trả về một thông báo cho client
	})

	log.Fatal(app.Listen(":" + port))
}
