package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Ping struct {
	Times int `json:"times"`
}

type Pong struct {
	Pongs []string `json:"pongs"`
}

func CraftResponse(pong *Pong, times int) {
	for i := 0; i < times; i++ {
		pong.Pongs = append(pong.Pongs, "pong")
	}
}

func initServer() fiber.App {
	app := fiber.New()
	app.Post("/ping", func(c *fiber.Ctx) error {
		body := Ping{}
		if err := c.BodyParser(&body); err != nil {
			return err
		}
		var pong = Pong{}
		CraftResponse(&pong, body.Times)
		return c.JSON(pong)
	})
	return *app
}

func main() {
	port := 8080
	app := initServer()
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Error starting server")
	}
}
