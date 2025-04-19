package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"kasir/config"
	"kasir/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func ReadUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}
	return c.JSON(user)
}

func ReadUsersPaginated(c *fiber.Ctx) error {
	var users []models.User

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	search := c.Query("search", "")
	sort := c.Query("sort", "")

	offset := (page - 1) * limit

	query := config.DB.Model(&models.User{})

	if search != "" {
		query = query.Where("username ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if sort != "" {
		parts := strings.Split(sort, "_")
		if len(parts) == 2 {
			col := parts[0]
			dir := strings.ToUpper(parts[1])
			if (col == "username" || col == "email" || col == "name") && (dir == "ASC" || dir == "DESC") {
				query = query.Order(fmt.Sprintf("%s %s", col, dir))
			}
		}
	}

	var total int64
	query.Count(&total)

	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	resp := map[string]interface{}{
		"data":       users,
		"page":       page,
		"limit":      limit,
		"total":      total,
		"totalPages": (total + int64(limit) - 1) / int64(limit),
	}

	return c.JSON(resp)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := config.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func DeleteAllUsers(c *fiber.Ctx) error {
	if err := config.DB.Where("1 = 1").Delete(&models.User{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
