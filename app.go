package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDB() {
	dsn := "root:123456@tcp(localhost:3306)/tour_web?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể kết nối với cơ sở dữ liệu: ", err)
	}
	fmt.Println("Kết nối cơ sở dữ liệu thành công.")
}
func main() {
	connectDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("ok")
		return c.SendString("Hello, World!") // Trả về một thông báo cho client
	})

	log.Fatal(app.Listen(":3000"))
}
