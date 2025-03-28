package jokenpo

import (
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type HTTPHandler struct {
	service *Service
}

func NewHTTPHandler(app *fiber.App, s *Service) {
	h := &HTTPHandler{
		service: s,
	}

	app.Post("/jokenpo/:option", h.Play)
}

func (h *HTTPHandler) Play(c fiber.Ctx) error {
	playerOption, err := strconv.Atoi(c.Params("option"))
	if err != nil {
		return err
	}

	computerOption := rand.Intn(2)

	winner, options := h.service.Play(playerOption, computerOption)

	return c.JSON(map[string]string{
		"user":     options[0].String(),
		"computer": options[1].String(),
		"winner":   winner.String(),
	})
}
