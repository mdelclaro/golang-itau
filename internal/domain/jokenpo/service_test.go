package jokenpo_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"golang-jokenpo/internal/domain/jokenpo"
	jokenpo_entity "golang-jokenpo/internal/pkg/entity/jokenpo"
)

func TestPlay(t *testing.T) {
	service := jokenpo.NewService()

	tests := []struct {
		name           string
		playerOption   int
		computerOption int
		expectedWinner jokenpo_entity.Player
	}{
		{
			name:           "Empate",
			playerOption:   0,
			computerOption: 0,
			expectedWinner: jokenpo_entity.Empate,
		},
		{
			name:           "Empate",
			playerOption:   1,
			computerOption: 1,
			expectedWinner: jokenpo_entity.Empate,
		},
		{
			name:           "Empate",
			playerOption:   2,
			computerOption: 2,
			expectedWinner: jokenpo_entity.Empate,
		},
		{
			name:           "User Pedra perde de Papel",
			playerOption:   0,
			computerOption: 1,
			expectedWinner: jokenpo_entity.Computer,
		},
		{
			name:           "User Pedra ganha de Tesoura",
			playerOption:   0,
			computerOption: 2,
			expectedWinner: jokenpo_entity.User,
		},
		{
			name:           "User Papel ganha de Pedra",
			playerOption:   1,
			computerOption: 0,
			expectedWinner: jokenpo_entity.User,
		},
		{
			name:           "User Papel perde de Tesoura",
			playerOption:   1,
			computerOption: 2,
			expectedWinner: jokenpo_entity.Computer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, moves := service.Play(tt.playerOption, tt.computerOption)

			fmt.Println(tt)
			fmt.Println(result, moves)

			assert.Equal(t, tt.expectedWinner, result)
			assert.Equal(t, service.Options[tt.playerOption], moves[0])
			assert.Equal(t, service.Options[tt.computerOption], moves[1])
		})
	}
}
