package main

import (
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/piyushpatel2005/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "server listen address")
	flag.Parse()
	app := fiber.New()

	apiv1 := app.Group("/api/v1")
	app.Get("/foo", handleFoo)

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"message": "Hello there",
	})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"user": "john.smith",
	})
}
