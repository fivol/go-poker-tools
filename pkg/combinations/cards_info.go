package combinations

import (
	"go-poker-tools/pkg/types"
)

type cardsInfo struct {
	cards  []types.Card
	suits  [4]uint8
	values [13]uint8

	// 3 6 8 6 -> [2:0, 3:3, 4:0, 5:0, 6:2, 7:0, 8:1]
	valueOrder        [13]int
	maxOrderLen       uint8
	maxSameSuitsCount uint8
	maxSameValues     uint8
}

func newCardsInfo(cards []types.Card) (ci cardsInfo) {
	ci.cards = cards
	if !types.IsDistinct(cards...) {
		panic("hand and board intersects, can not extract winner")
	}
	for _, card := range cards {
		ci.suits[card.Suit()]++
		ci.maxSameSuitsCount = Max(maxSameSuitsCount, ci.suits[card.Suit()])
		ci.values[card.Value()]++
	}

	return
}
