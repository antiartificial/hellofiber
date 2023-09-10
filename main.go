package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

func main() {
	app := New()

	log.Fatal(app.Listen(":3000"))
}

func New() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/params/+", handler)

	app.Get("/test/*", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("*"))
	})

	app.Get("/park/+", func(c *fiber.Ctx) error {
		//return c.SendString(c.Params("+"))
		return c.SendString(c.Query("context"))
	})

	app.Get("/availability/:context", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.SendString(c.Params("context"))
	})

	app.Post("/availability", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		log.Println(p.Name)
		log.Println(p.Pass)

		return c.Status(fiber.StatusOK).JSON(p)
	})

	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	fmt.Println(string(data))

	return app
}

func handler(c *fiber.Ctx) error {
	return c.SendString(c.Params("key1"))
}
