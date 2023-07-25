package userHandlers

import (
	"github.com/RugeFX/go-fiber-1.git/database"
	"github.com/RugeFX/go-fiber-1.git/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Gets all users
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "failed",
			"error":  "Users not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   users,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	db := database.DB
	var user models.User

	db.First(&user, c.Params("id"))

	if !user.Username.Valid {
		return c.Status(404).JSON(fiber.Map{
			"status": "failed",
			"error":  "User not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	// newPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	// user.Password = string(newPass)

	result := db.Create(&user)

	if err := result.Error; err != nil {
		return c.Status(200).JSON(fiber.Map{
			"status": "failed",
			"error":  err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}
