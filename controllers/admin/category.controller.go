package controllers

import (
	"myapp/config"
	"myapp/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}
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

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category

	// Parse JSON request body
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request body"})
	}
	// validate data
	err := validate.Struct(category)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	// Lưu danh mục vào cơ sở dữ liệu
	if err := config.DB.Create(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating category"})
	}
	var categoryView models.CategoryView
	if err := mapstructure.Decode(category, &categoryView); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error converting category to CategoryView"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":  "Category created successfully",
		"category": categoryView,
	})
}
