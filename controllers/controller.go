package controllers

import (
	"amquizdua/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"math/rand"
)

//func Index(c *fiber.Ctx) error {

//	var quiz []models.Quizdua

//	models.DB.Db.Find(&quiz)

//	return c.Status(fiber.StatusOK).JSON(quiz)
//}

func Index(c *fiber.Ctx) error {
	var quiz []models.Quizdua

	var count int64
	if err := models.DB.Db.Model(&models.Quizdua{}).Count(&count).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to count records"})
	}

	indices := rand.Perm(int(count))[:2]

	for _, index := range indices {
		var quizEntry models.Quizdua
		if err := models.DB.Db.First(&quizEntry, index).Error; err == nil {
			quiz = append(quiz, quizEntry)
		}
	}

	return c.Status(fiber.StatusOK).JSON(quiz)
}

func Create(c *fiber.Ctx) error {

	quiz := new(models.Quizdua)

	if err := c.BodyParser(quiz); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	models.DB.Db.Create(&quiz)

	return c.Status(fiber.StatusCreated).JSON(quiz)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var quiz models.Quizdua

	result := models.DB.Db.Where("id = ?", id).First(&quiz)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Message": "Quiz not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": result.Error.Error(),
		})
	}

	return c.JSON(quiz)
}

func Update(c *fiber.Ctx) error {

	id := c.Params("id")
	var updatedData models.Quizdua

	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	if models.DB.Db.Where("id = ?", id).Updates(&updatedData).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Id tidak ditemukan.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(updatedData)
}
