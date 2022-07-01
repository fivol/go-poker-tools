package poker

import (
	"strings"
)

type Range []WeightedHand

func ParseRange(rangeStr string) Range {
	handsStr := strings.Split(rangeStr, ",")
	var hands []WeightedHand
	for _, handStr := range handsStr {
		hands = append(hands, ParseWightedHand(handStr))
	}
	return hands
}

func RemoveHands(r Range, hands []Hand) Range {
	var rangeHands []WeightedHand;
	for _, hand := range r {

	}
}
