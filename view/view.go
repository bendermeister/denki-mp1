package view

import (
	"mp1/query"
	"mp1/query/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Init(g *fiber.Group) {
	g.Get("/", getBody)
}

func parseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func getTable(c *fiber.Ctx) error {
	ui, err := parseInt(c.FormValue("hasUI", "asdf"))
	if err != nil {
		return err
	}
	min, err := parseInt(c.FormValue("min", "asdf"))
	if err != nil {
		return err
	}
	max, err := parseInt(c.FormValue("max", "adsf"))
	if err != nil {
		return err
	}

	var projects []db.Project

	switch ui {
	case 0:
		projects, err = query.GetUI(min, max)
	case 1:
		projects, err = query.GetNoUI()
	case 2:
		projects, err = query.GetAll()
	}
	if err != nil {
		return err
	}

	return c.Render("table", fiber.Map{
		"projects": projects,
	})
}

func getBody(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}
