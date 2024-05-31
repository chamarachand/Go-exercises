package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	app := fiber.New()

	// Data
	users := []User{
		{ID: 1, Name: "Ashan Silva", Email: "ashan@gmail.com"},
		{ID: 2, Name: "Pramod Perera", Email: "pramod@gmail.com"},
		{ID: 3, Name: "Asanka Fernando", Email: "asanka@gmail.com"},
	}

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(users)
	})

	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for _, user := range users {
			if strconv.Itoa(user.ID) == id {
				return c.JSON(user)
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	})

	app.Listen(":3000")
	fmt.Println("Listening on port 3000...")
}
