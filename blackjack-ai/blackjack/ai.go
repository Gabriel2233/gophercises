package blackjack

import (
	"fmt"

	"github.com/Gabriel2233/gophercises/deck"
)

type AI interface {
	Bet() int
	Play(hand, dealer []deck.Card) Move
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type HumanAI struct{}

func (ai *HumanAI) Bet() int {
	return 1
}

func (ai *HumanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		var input string

		fmt.Println("Player: ", hand)
		fmt.Println("Dealer: ", dealer)
		fmt.Println("What are you going to do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)

		switch input {
		case "h":
			return Hit
		case "s":
			return Stand
		default:
			fmt.Println("Invalid option: ", input)
		}
	}
}

func (ai *HumanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("== FINAL HANDS ==")
	fmt.Println("Player: ", hand)
	fmt.Println("Dealer: ", dealer)
}
