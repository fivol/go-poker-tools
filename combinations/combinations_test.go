package combinations

import (
	"github.com/stretchr/testify/assert"
	"go-poker-equity/poker"
	"sort"
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
	for _, testCase := range table {
		assert.True(t, testCase.big.GraterThen(testCase.small), "Combinations compare failed")
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
			findHighComb,
			"AsAdAc",
			"2h3d",
			true,
			Combination{name: High, values: [5]uint8{12, 12, 12, 1, 0}},
		},
		{
			findHighComb,
			"2c3h9d",
			"2h3d",
			true,
			Combination{name: High, values: [5]uint8{7, 1, 1, 0, 0}},
		},
		{
			findHighComb,
			"2c3h9d",
			"2hTd",
			true,
			Combination{name: High, values: [5]uint8{8, 7, 1, 0, 0}},
		},
		{
			findPairComb,
			"AsAdAc",
			"2h3d",
			true,
			Combination{name: Pair, values: [5]uint8{12, 12, 12, 1, 0}},
		},
		{
			findPairComb,
			"Ts2dAc",
			"2h3d",
			true,
			Combination{name: Pair, values: [5]uint8{0, 0, 12, 8, 1}},
		},
		{
			findSetComb,
			"AsAdAc",
			"2h3d",
			true,
			Combination{name: Set, values: [5]uint8{12, 12, 12, 1, 0}},
		},
		{
			findSetComb,
			"AsAd3c",
			"Ah3d",
			true,
			Combination{name: Set, values: [5]uint8{12, 12, 12, 1, 1}},
		},
		{
			findQuadsComb,
			"AsAdAc",
			"2h3d",
			false,
			Combination{name: Set, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			findQuadsComb,
			"AsAdAc",
			"2hAh",
			true,
			Combination{name: Quads, values: [5]uint8{12, 12, 12, 12, 0}},
		},
		{
			findTwoPairsComb,
			"AsAdAc",
			"2h2c",
			true,
			Combination{name: TwoPairs, values: [5]uint8{12, 12, 0, 0, 12}},
		},
		{
			findTwoPairsComb,
			"AsAdAc",
			"2h3c",
			false,
			Combination{name: TwoPairs, values: [5]uint8{12, 0, 0, 0, 0}},
		},
		{
			findTwoPairsComb,
			"9s7cTd9c",
			"6h6d",
			true,
			Combination{name: TwoPairs, values: [5]uint8{7, 7, 4, 4, 8}},
		},
		{
			findStraightComb,
			"4c5c6c7c8c",
			"2c3c",
			true,
			Combination{name: Straight, values: [5]uint8{6, 0, 0, 0, 0}},
		},
		{
			findStraightComb,
			"4c5c6cTc8c",
			"2c3c",
			true,
			Combination{name: Straight, values: [5]uint8{4, 0, 0, 0, 0}},
		},
		{
			findStraightComb,
			"4c5cTcTd8c",
			"2c3c",
			false,
			Combination{name: Straight, values: [5]uint8{4, 0, 0, 0, 0}},
		},
		{
			findStraightComb,
			"4c5cTcTd8c",
			"2c3c",
			false,
			Combination{name: Flush, values: [5]uint8{8, 6, 3, 2, 1}},
		},
		{
			findStraightFlushComb,
			"4c5cTcTd8c",
			"2c3c",
			false,
			Combination{name: Flush, values: [5]uint8{8, 6, 3, 2, 1}},
		},
		{
			findStraightFlushComb,
			"4c5c6c7d8c",
			"2c3c",
			true,
			Combination{name: StraightFlush, values: [5]uint8{4, 0, 0, 0, 0}},
		},
		{
			findFullHouseComb,
			"6d5c6cAdAc",
			"2c6s",
			true,
			Combination{name: FullHouse, values: [5]uint8{4, 12, 0, 0, 0}},
		},
		{
			findFullHouseComb,
			"6d5c6cAdAc",
			"2c5s",
			false,
			Combination{name: FullHouse, values: [5]uint8{4, 12, 0, 0, 0}},
		},
	}
	for _, testCase := range table {
		board := poker.ParseBoard(testCase.board)
		hand := poker.ParseHand(testCase.hand)
		selector := newCombinationsSelector(board, hand)
		selector.calcCardsEntry()
		combination, found := testCase.extractor(&selector)
		assert.Equal(t, testCase.found, found, "Combination found not match")
		if testCase.found {
			assert.Equal(t, testCase.combination, combination, "Combination not match")
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
		board   string
		hands   []string
		winners []int
	}{
		{
			"As9h8c4c3h",
			[]string{"Ac2s", "9s2s"},
			[]int{0},
		},
		{
			"7s9h8c4c3h",
			[]string{"9c8s", "5s6s"},
			[]int{1},
		},
		{
			"4s5s6s7s8s",
			[]string{"Td9s", "3h2s", "Ac2h"},
			[]int{0},
		},
		{
			"4s5s6s7s8s",
			[]string{"Td9c", "3h2s", "Ac2h"},
			[]int{0, 1, 2},
		},
		{
			"2s8d5dThQc",
			[]string{"TsTc", "2c2h"},
			[]int{0},
		},
		{
			"2s8d5dThQc",
			[]string{"AsAc", "2c2h"},
			[]int{1},
		},
		{
			"2s8d5dThQc",
			[]string{"As3h", "Kc3d", "Ac4h"},
			[]int{0, 2},
		},
		{
			"2s8d5dThQc",
			[]string{"As3h", "Kc3d", "Ac6h"},
			[]int{2},
		},
		{
			"2s8d5dThQc",
			[]string{"2cKs", "Ts2h"},
			[]int{1},
		},
		{
			"2s8d5dThQc",
			[]string{"2cKs", "As3h"},
			[]int{0},
		},
	}
	for _, testCase := range table {
		board := poker.ParseBoard(testCase.board)
		var hands []poker.Hand
		for _, hand := range testCase.hands {
			hands = append(hands, poker.ParseHand(hand))
		}
		winners := DetermineWinners(board, hands)
		sort.Ints(winners)
		sort.Ints(testCase.winners)
		assert.Equal(t, testCase.winners, winners, "Winners sets does not match")
	}

}
