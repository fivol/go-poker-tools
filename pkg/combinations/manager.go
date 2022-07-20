package combinations

import "go-poker-tools/pkg/types"

func newCombinationsSelector(board types.Board, hand types.Hand) Selector {
	c1, c2 := hand.Cards()
	cards := make([]types.Card, len(board)+2)
	copy(cards, board)
	cards[len(board)] = c1
	cards[len(board)+1] = c2
}

type CombinationExtractor func(c *Selector) bool

type ExtractorWithName struct {
	extractor CombinationExtractor
	name      Comb
}

var combinations = []ExtractorWithName{
	{findStraightFlush, "straight_flush"},
}

func getCombinations(board types.Board, hand types.Hand) []Comb {
	selector := newCombinationsSelector(board, hand)
	selector.calcCardsEntry()
	var combs []Comb
	for combData := range combinations {
		combData
	}
	return combs
}
