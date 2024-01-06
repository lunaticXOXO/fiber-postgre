package controller

import (

	"github.com/gofiber/fiber/v2"
	"github.com/fiber-postgre/model"

)


func AddUsertype(c *fiber.Ctx) error{
	var types model.Usertype

	if err := c.BodyParser(&types); err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := model.DB.Create(&types).Error; err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "success",
		"data" : types,
	})
}