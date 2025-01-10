package controllers

import (
	"myapp/config"
	"myapp/models"

	"github.com/gofiber/fiber/v2"
)

func GetCategory(c *fiber.Ctx) error {
	var categories []models.CategoryView
	err := config.DB.Model(&models.Category{}).Where("deleted = ? AND status = ?", false, "active").Find(&categories).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Error fetching categories",
		})
	}
	return c.JSON(fiber.Map{
		"categories": categories,
	})
}
