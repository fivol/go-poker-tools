package poker

import (
	"fmt"
	"strconv"
	"strings"
)

type Hand uint16

func NewHand(c1, c2 Card) Hand {
	if c1.Grater(c2) {
		c1, c2 = c2, c1
	}
	return Hand(c1*52 + c2)
}

func (h Hand) Cards() (Card, Card) {
	return Card(h / 52), Card(h % 52)
}

func (h Hand) ToString() string {
	c1, c2 := h.Cards()
	return c1.ToString() + c2.ToString()
}

func ParseHand(handStr string) Hand {
	c1 := ParseCard(handStr[:2])
	c2 := ParseCard(handStr[2:])
	if !IsDistinct(c1, c2) {
		panic("try parse hand with same cards")
	}
	return NewHand(c1, c2)
}

func ParseWightedHand(handStr string) (Hand, float32) {
	handWeightArr := strings.Split(handStr, ":")
	var weight float32 = 1
	if len(handStr) == 2 {
		w, err := strconv.ParseFloat(handWeightArr[1], 32)
		if err != nil {
			panic(fmt.Sprintf("Weight format not correct: %v", handStr))
		}
		weight = float32(w)
	}
	hand := ParseHand(handWeightArr[0])
	return hand, weight
}
