package combinations

import "go-poker-equity/poker"

type Selector struct {
	cards  []poker.Card
	suits  [4]uint8
	values [13]uint8
}

func newCombinationsSelector(board poker.Board, hand poker.Hand) Selector {
	cards := hand
	for _, card := range board {
		cards = append(cards, card)
	}
	return Selector{cards: cards}
}

func (c *Selector) calcCardsEntry() {
	for _, card := range c.cards {
		c.suits[poker.Suit(card)]++
		c.values[poker.Value(card)]++
	}
}

type CombinationExtractor func(c *Selector) (Combination, bool)

func extractCombinations(board poker.Board, hand poker.Hand) []Combination {
	selector := newCombinationsSelector(board, hand)
	selector.calcCardsEntry()

	var combinations []Combination
	extractors := []CombinationExtractor{FindCombHigh}

	for _, extractor := range extractors {
		combination, found := extractor(&selector)
		if found {
			combinations = append(combinations, combination)
		}
	}
	return combinations
}

func selectHighestCombination(combinations []Combination) Combination {
	best := combinations[0]
	for _, c := range combinations {
		if c.GraterThen(best) {
			best = c
		}
	}
	return best
}

func DetermineWinners(board poker.Board, hands []poker.Hand) []poker.Hand {
	var winners []poker.Hand
	var handsCombos []Combination
	for _, hand := range hands {
		combinations := extractCombinations(board, hand)
		highestComb := selectHighestCombination(combinations)
		handsCombos = append(handsCombos, highestComb)
	}
	highestComb := selectHighestCombination(handsCombos)
	for i := 0; i < len(hands); i++ {
		if highestComb == handsCombos[i] {
			winners = append(winners, hands[i])
		}
	}
	return winners
}
