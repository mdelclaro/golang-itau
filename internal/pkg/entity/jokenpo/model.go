package jokenpo

type Player string
type Option string

const (
	User     Player = "user"
	Computer Player = "computer"
	Empate   Player = "empate"

	Pedra   Option = "pedra"
	Papel   Option = "papel"
	Tesoura Option = "tesoura"
)

func (p Player) String() string {
	return string(p)
}

func (o Option) String() string {
	return string(o)
}
