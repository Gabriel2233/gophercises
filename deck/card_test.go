package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{
		Suit: Diamond,
		Rank: Ace,
	})
	fmt.Println(Card{
		Suit: Spade,
		Rank: Queen,
	})
	fmt.Println(Card{
		Suit: Club,
		Rank: Ten,
	})
	fmt.Println(Card{
		Suit: Heart,
		Rank: Two,
	})
	fmt.Println(Card{
		Suit: Joker,
	})

	// Output:
	// Ace of Diamonds
	// Queen of Spades
	// Ten of Clubs
	// Two of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()

	got := len(cards)
	want := 52

	if got != want {
		t.Errorf("want %d got %d", want, got)
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)

	got := cards[0]
	want := Card{Suit: Spade, Rank: Ace}

	if got != want {
		t.Error("expected Ace of Spades to be the first card. Received: ", got)
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))

	got := cards[0]
	want := Card{Suit: Spade, Rank: Ace}

	if got != want {
		t.Error("expected Ace of Spades to be the first card. Received: ", got)
	}
}

func TestShuffle(t *testing.T) {
	shuffleRand = rand.New(rand.NewSource(0))

	orig := New()
	first := orig[40]
	second := orig[35]

	cards := New(Shuffle)

	if cards[0] != first {
		t.Errorf("expecetd first card to be %d, received %d", first, cards)
	}
	if cards[1] != second {
		t.Errorf("expecetd second card to be %d, received %d", second, cards)
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0

	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}

	if count != 3 {
		t.Error("Expected 3 Jokers. Got: ", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))
	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Expected all twos and threes to be filtered out")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))

	got := len(cards)
	want := 13 * 4 * 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
