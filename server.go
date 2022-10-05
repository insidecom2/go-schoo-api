package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/:value", func(c *fiber.Ctx) error {
		primes := []string{"2", "3", "5", "7", "11", "13"}
		var arrays [1]any
		arrays[0] = "55555"
		
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"success": fiber.StatusOK,
			"message": "Successfully",
			"data": primes,
			"data2": arrays,
		  })
	})

	app.Listen(":3000")
}