package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/piyushpatel2005/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "John",
		LastName:  "Smith",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"user": "john.smith",
	})
}
