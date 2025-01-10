package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // Xuất khẩu biến DB

// ConnectDB thực hiện kết nối đến database
func ConnectDB() {
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
	var errOpen error
	DB, errOpen = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errOpen != nil {
		log.Fatal("Không thể kết  nối với cơ sở dữ liệu: ", errOpen)
	}
	fmt.Println("Kết nối cơ sở dữ liệu thành công.")
}
