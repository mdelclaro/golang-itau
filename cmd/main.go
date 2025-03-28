package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"golang-itau/internal/config"
	"golang-itau/internal/domain/jokenpo"
)

func main() {
	app := fiber.New()

	service := jokenpo.NewService()

	jokenpo.NewHTTPHandler(app, service)

	log.Fatal(app.Listen(config.GetEnv("APP_PORT")))
}
