package view

import (
	"math"
	"mp1/query"
	"mp1/query/db"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Init(g *fiber.Group) {
	g.Get("/", getBody)
	g.Get("/table", getTable)
	g.Get("/stack", getStack)
	g.Get("/thanks", getThanks)
}

func round(f float64) float64 {
	return math.Round(f*100) / 100
}

func parseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func getStack(c *fiber.Ctx) error {
	return c.Render("stack", fiber.Map{})
}

func getThanks(c *fiber.Ctx) error {
	return c.Render("thanks", fiber.Map{})
}

func getTable(c *fiber.Ctx) error {
	ui, err := parseInt(c.FormValue("hasUI", "hasUI"))
	if err != nil {
		return err
	}
	min, err := parseInt(c.FormValue("min", "min"))
	if err != nil {
		return err
	}
	max, err := parseInt(c.FormValue("max", "max"))
	if err != nil {
		return err
	}

	pCount, err := query.Count()
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
	sCount := len(projects)

	avgPoints := int64(0)
	for i, _ := range projects {
		avgPoints += projects[i].Points
	}

	return c.Render("table", fiber.Map{
		"rCount":    round(float64(sCount) / float64(pCount) * 100.0),
		"pCount":    pCount,
		"sCount":    sCount,
		"projects":  projects,
		"avgPoints": round(float64(avgPoints) / float64(sCount)),
	})
}

func getBody(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}
