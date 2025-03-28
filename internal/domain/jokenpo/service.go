package jokenpo

import (
	"golang-jokenpo/internal/pkg/entity/jokenpo"
)

type Service struct {
	Options []jokenpo.Option
}

func NewService() *Service {
	return &Service{
		Options: []jokenpo.Option{jokenpo.Pedra, jokenpo.Papel, jokenpo.Tesoura},
	}
}

// 0 - Pedra
// 1 - Papel
// 2 - Tesoura
func (s *Service) Play(playerOption, computerOption int) (jokenpo.Player, []jokenpo.Option) {
	computerMove := s.Options[computerOption]
	playerMove := s.Options[playerOption]

	if playerMove == computerMove {
		return jokenpo.Empate, []jokenpo.Option{playerMove, computerMove}
	}

	if playerMove == jokenpo.Pedra && computerMove == jokenpo.Tesoura ||
		playerMove == jokenpo.Tesoura && computerMove == jokenpo.Papel ||
		playerMove == jokenpo.Papel && computerMove == jokenpo.Pedra {
		return jokenpo.User, []jokenpo.Option{playerMove, computerMove}
	}

	return jokenpo.Computer, []jokenpo.Option{playerMove, computerMove}
}
