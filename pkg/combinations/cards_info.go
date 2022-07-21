package combinations

import (
	"go-poker-tools/pkg/generics"
	"go-poker-tools/pkg/types"
)

type cardsInfo struct {
	cards []types.Card
	//chart  [4][13]bool
	suits  [4]uint8
	values [13]uint8

	// 3 6 8 6 -> [2:0, 3:3, 4:0, 5:0, 6:2, 7:0, 8:1]
	valueOrder        [13]uint8
	graterValuesCount [13]uint8
	stairsUpLen       [13]uint8
	stairsDownLen     [13]uint8
	maxValueSuits     [4]uint8
	maxValue          uint8
	minValue          uint8
	//maxFreeValue      [4]uint8
	maxOrderLen       uint8
	maxOrderIdx       uint8
	maxSameSuitsCount uint8
	maxSameValues     uint8
}

func newCardsInfo(cards []types.Card) (ci cardsInfo) {
	ci.cards = cards
	if !types.IsDistinct(cards...) {
		panic("hand and board intersects, can not extract winner")
	}
	ci.minValue = 12
	for _, card := range cards {
		//ci.chart[card.Suit()][card.Value()] = true
		ci.suits[card.Suit()]++
		ci.maxSameSuitsCount = generics.Max(ci.maxSameSuitsCount, ci.suits[card.Suit()])
		ci.values[card.Value()]++
		ci.maxValueSuits[card.Suit()] = generics.Max(ci.maxValueSuits[card.Suit()], card.Value())
		ci.maxValue = generics.Max(ci.maxValue, card.Value())
		ci.minValue = generics.Min(ci.minValue, card.Value())
		ci.maxSameValues = generics.Max(ci.maxSameValues, ci.values[card.Value()])
	}
	//for suit := 0; suit < 4; suit++ {
	//	for value := 12; value >= 0; value-- {
	//		if !ci.chart[suit][value] {
	//			ci.maxFreeValue[suit] = uint8(value)
	//			break
	//		}
	//	}
	//}
	order := uint8(1)
	maxOrder := uint8(0)
	for i := 12; i >= 0; i-- {
		if ci.values[i] > 0 {
			ci.valueOrder[i] = order
			order++
			maxOrder++
		} else {
			maxOrder = 0
		}
		ci.stairsUpLen[i] = maxOrder
		ci.graterValuesCount[i] = order - 1
		if maxOrder > ci.maxOrderLen {
			ci.maxOrderLen = maxOrder
			ci.maxOrderIdx = uint8(i)
		}
	}
	if ci.values[12] > 0 {
		maxOrder++
		ci.stairsUpLen[12] = maxOrder
		if maxOrder > ci.maxOrderLen {
			ci.maxOrderLen = maxOrder
			ci.maxOrderIdx = 12
		}
	}

	maxOrder = 0
	for i := -1; i <= 12; i++ {
		val := (i + 13) % 13
		if ci.values[val] > 0 {
			maxOrder++
		} else {
			maxOrder = 0
		}
		ci.stairsDownLen[val] = maxOrder
	}
	return
}
