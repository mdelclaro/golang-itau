package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"golang-jokenpo/internal/config"
	"golang-jokenpo/internal/domain/jokenpo"
)

func main() {
	app := fiber.New()

	service := jokenpo.NewService()

	jokenpo.NewHTTPHandler(app, service)

	log.Fatal(app.Listen(config.GetEnv("APP_PORT")))
}
