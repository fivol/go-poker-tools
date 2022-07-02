package combinations

import (
	"github.com/stretchr/testify/assert"
	"go-poker-equity/poker"
	"testing"
)

func TestCombinationsCompare(t *testing.T) {
	table := []struct {
		small Combination
		big   Combination
	}{
		{
			Combination{name: Flush, values: [5]uint8{5, 0, 0, 0, 0}},
			Combination{name: Flush, values: [5]uint8{8, 0, 0, 0, 0}},
		},
		{
			Combination{name: TwoPairs, values: [5]uint8{0, 0, 0, 0, 0}},
			Combination{name: Set, values: [5]uint8{0, 0, 0, 0, 0}},
		},
		{
			Combination{name: High, values: [5]uint8{13, 0, 0, 0, 0}},
			Combination{name: Pair, values: [5]uint8{0, 0, 0, 0, 0}},
		},
		{
			Combination{name: High, values: [5]uint8{13, 12, 11, 10, 0}},
			Combination{name: High, values: [5]uint8{13, 12, 11, 10, 1}},
		},
	}
	for _, suitCase := range table {
		assert.True(t, suitCase.big.GraterThen(suitCase.small), "Combinations compare failed")
	}
}

func TestExtractors(t *testing.T) {
	table := []struct {
		extractor   CombinationExtractor
		board       string
		hand        string
		found       bool
		combination Combination
	}{
		{
			FindHighComb,
			"AsAdAc",
			"2h3d",
			true,
			Combination{name: High, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			FindHighComb,
			"2c3h9d",
			"2h3d",
			true,
			Combination{name: High, values: [5]uint8{7, 0, 0, 0, 0}},
		},
		{
			FindHighComb,
			"2c3h9d",
			"2hTd",
			true,
			Combination{name: High, values: [5]uint8{8, 0, 0, 0, 0}},
		},
		{
			FindPairComb,
			"AsAdAc",
			"2h3d",
			true,
			Combination{name: Pair, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			FindSetComb,
			"AsAdAc",
			"2h3d",
			true,
			Combination{name: Set, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			FindSetComb,
			"AsAd3c",
			"Ah3d",
			true,
			Combination{name: Set, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			FindQuadsComb,
			"AsAdAc",
			"2h3d",
			false,
			Combination{name: Set, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			FindQuadsComb,
			"AsAdAc",
			"2hAh",
			true,
			Combination{name: Quads, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			FindTwoPairsComb,
			"AsAdAc",
			"2h2c",
			true,
			Combination{name: TwoPairs, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			FindTwoPairsComb,
			"AsAdAc",
			"2h3c",
			false,
			Combination{name: TwoPairs, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			FindTwoPairsComb,
			"9s7cTd9c",
			"6h6d",
			true,
			Combination{name: TwoPairs, values: [5]uint8{7, 4, 0, 0, 0}},
		},
		{
			FindStraightComb,
			"4c5c6c7c8c",
			"2c3c",
			true,
			Combination{name: Straight, values: [5]uint8{6, 0, 0, 0, 0}},
		},
		{
			FindStraightComb,
			"4c5c6cTc8c",
			"2c3c",
			true,
			Combination{name: Straight, values: [5]uint8{4, 0, 0, 0, 0}},
		},
		{
			FindStraightComb,
			"4c5cTcTd8c",
			"2c3c",
			false,
			Combination{name: Straight, values: [5]uint8{4, 0, 0, 0, 0}},
		},
		{
			FindStraightComb,
			"4c5cTcTd8c",
			"2c3c",
			false,
			Combination{name: Flush, values: [5]uint8{8, 6, 3, 2, 1}},
		},
		{
			FindStraightFlushComb,
			"4c5cTcTd8c",
			"2c3c",
			false,
			Combination{name: Flush, values: [5]uint8{8, 6, 3, 2, 1}},
		},
		{
			FindStraightFlushComb,
			"4c5c6c7d8c",
			"2c3c",
			true,
			Combination{name: StraightFlush, values: [5]uint8{4, 0, 0, 0, 0}},
		},
	}
	for _, suitCase := range table {
		board := poker.ParseBoard(suitCase.board)
		hand := poker.ParseHand(suitCase.hand)
		selector := newCombinationsSelector(board, hand)
		selector.calcCardsEntry()
		combination, found := suitCase.extractor(&selector)
		assert.Equal(t, suitCase.found, found, "Combination found not match")
		if suitCase.found {
			assert.Equal(t, suitCase.combination, combination, "Combination not match")
		}
	}
	assert.Panics(t, func() {
		newCombinationsSelector(poker.ParseBoard("Ts3h5c"), poker.ParseHand("5cTd"))
	}, "Must panic that cards intersects")
	assert.Panics(t, func() {
		newCombinationsSelector(poker.ParseBoard("5c6c7c8c9c"), poker.ParseHand("7c9c"))
	}, "Must panic that cards intersects")
}

func TestWinners(t *testing.T) {
	table := []struct {
		board       string
		hand        string
		found       bool
		combination Combination
	}{{}}
	DetermineWinners()
}
