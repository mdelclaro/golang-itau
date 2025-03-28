package jokenpo_integration_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"

	"golang-jokenpo/internal/domain/jokenpo"
)

func TestJokenpoEndpoint(t *testing.T) {
	app := fiber.New()

	service := jokenpo.NewService()
	jokenpo.NewHTTPHandler(app, service)

	tests := []struct {
		name         string
		option       string
		expectedCode int
	}{
		{
			name:         "Player chooses Pedra (0)",
			option:       "0",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Player chooses Papel (1)",
			option:       "1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Player chooses Tesoura (2)",
			option:       "2",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid option",
			option:       "3",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(
				http.MethodPost,
				"/jokenpo/"+tt.option,
				nil,
			)

			res, err := app.Test(req, fiber.TestConfig{Timeout: -1})
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, res.StatusCode)

			if res.StatusCode == http.StatusOK {
				var response map[string]string
				err := json.NewDecoder(res.Body).Decode(&response)
				assert.NoError(t, err)

				assert.Contains(t, []string{"pedra", "papel", "tesoura"}, response["user"])
				assert.Contains(t, []string{"pedra", "papel", "tesoura"}, response["computer"])
				assert.Contains(t, []string{"user", "computer", "empate"}, response["winner"])
			}
		})
	}
}
