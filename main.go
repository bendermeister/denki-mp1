package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"mp1/query"
	"mp1/view"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func loadData(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	reader := csv.NewReader(file)
	reader.Comma = ';'

	for row, err := reader.Read(); err != io.EOF; row, err = reader.Read() {
		log.Print(row)
		name := strings.TrimSpace(row[0])
		url := strings.TrimSpace(row[1])
		hasUI, err := strconv.ParseBool(row[2])
		if err != nil {
			return err
		}
		var points int64
		if !hasUI {
			points = 0
		} else {
			points, err = strconv.ParseInt(row[3], 10, 64)
			if err != nil {
				return err
			}
		}
		err = query.Insert(name, url, hasUI, points)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := query.Init()
	if err != nil {
		log.Fatal(err)
	}

	dataPath := flag.String("data", "", "[required] path to the data file in csv format")
	port := flag.String("port", "8080", "port on which to run")
	flag.Parse()

	err = loadData(*dataPath)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Print("[ERROR]: ", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		},
		Views: jet.New("./templates", ".jet.html"),
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/view/")
	})

	app.Static("/img", "./img")

	view.Init(app.Group("/view/").(*fiber.Group))

	// TODO customize which port to start process on
	app.Listen(fmt.Sprintf(":%s", *port))

}
