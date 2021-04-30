package blackjack

import "github.com/Gabriel2233/gophercises/deck"

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type State int8

type Game struct {
	deck   []deck.Card
	state  State
	player []deck.Card
	dealer []deck.Card
}

type Move func(Game) Game

func Hit(gs Game) Game {
	return gs
}

func Stand(gs Game) Game {
	return gs
}
