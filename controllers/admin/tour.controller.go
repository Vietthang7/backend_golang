package controllers

import (
	"myapp/config"
	"myapp/models"

	"github.com/gofiber/fiber/v2"
)

// GetTours lấy danh sách các tour
func GetTours(c *fiber.Ctx) error {
	var tours []models.TourView

	// Truy vấn cơ sở dữ liệu và lấy dữ liệu tour
	err := config.DB.Model(&models.Tour{}).Select("id", "schedule", "title").Where("deleted = ? AND status = ?", false, "active").Find(&tours).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error fetching tours",
		})
	}

	return c.JSON(fiber.Map{
		"pageTitle": "Danh sách tours",
		"tours":     tours,
	})
}
