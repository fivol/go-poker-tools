package poker

import (
	"strings"
)

const rangeLen uint16 = 2704

type Range [rangeLen]float32

type RangeIterator struct {
	r    *Range
	hand Hand
}

func (iter *RangeIterator) Next() (h Hand, weight float32, end bool) {
	for ; uint16(iter.hand) < rangeLen; iter.hand++ {
		if iter.r[iter.hand] > 0 {
			h = iter.hand
			weight = iter.r[h]
			iter.hand++
			return
		}
	}
	end = true
	return
}

func NewRangeIterator(range_ *Range) (iter RangeIterator) {
	iter.r = range_
	return
}

func ParseRange(rangeStr string) (r Range) {
	handsStr := strings.Split(rangeStr, ",")
	for _, handStr := range handsStr {
		hand, weight := ParseWightedHand(handStr)
		r[hand] = weight
	}
	return
}

func (r *Range) GetIterator() RangeIterator {
	return NewRangeIterator(r)
}

func (r *Range) RemoveHands(hands []Hand) {
	for _, hand := range hands {
		r[hand] = 0
	}
}

func (r *Range) RemoveCards(cards []Card) {
	var hands []Hand
	for _, c1 := range cards {
		for c2 := Card(0); c2 < 52; c2++ {
			if c1 != c2 {
				hands = append(hands, NewHand(c1, c2))
			}
		}
	}
	r.RemoveHands(hands)
}
