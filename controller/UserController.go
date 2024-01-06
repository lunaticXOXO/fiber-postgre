package controller

import (

	"github.com/gofiber/fiber/v2"
	"github.com/fiber-postgre/model"

)

func Register(c *fiber.Ctx) error {
	var user model.Users

	if err := c.BodyParser(&user); err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := model.DB.Create(&user).Error;err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "register success",
		"data" : user,
	})

}

func Login(c *fiber.Ctx) error {
	var user model.Users
	if err := c.BodyParser(&user);err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := model.DB.Begin().Where("username=? AND password=?",user.Username,user.Password).First(&user).Error; err == nil{
			get_token,err := GenerateToken(c,user)
			if err == nil{
				return c.Status(fiber.StatusOK).JSON(fiber.Map{
					"message" : "login success",
					"token" : get_token,
				})
			}
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message" : "user unauthorized",
	})


}

func Logout(c *fiber.Ctx) error{
	ResetToken(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "logout success",
	})
}

func GetUsers(c *fiber.Ctx) error {

	var users []model.Users
	if err := model.DB.Find(&users).Error; err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data" : users,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	var user model.Users

	if err := model.DB.Where("username","admin").Delete(&user).Error; err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "delete success",
	})
}