package main

import (
	"log"
	"mp1/query"
	"mp1/view"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func main() {
	err := query.Init()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New(fiber.Config{
		Views: jet.New("./templates", ".jet.html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/view/")
	})

	view.Init(app.Group("/view/").(*fiber.Group))

	// TODO customize which port to start process on
	app.Listen(":8080")

}
