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
	for i := iter.hand; uint16(i) < rangeLen; i++ {
		iter.hand++
		if iter.r[iter.hand] > 0 {
			h = iter.hand
			weight = iter.r[h]
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
