package poker

import (
	"fmt"
	"strconv"
	"strings"
)

type Hand []Card

type Weight float32

type WeightedHand struct {
	Hand   Hand
	Weight Weight
}

func (h Hand) ToString() string {
	return h[0].ToString() + h[1].ToString()
}

func (h WeightedHand) ToString() string {
	return fmt.Sprintf("%s:%f", h.Hand.ToString(), h.Weight)
}

func ParseHand(handStr string) Hand {
	card1str := handStr[:2]
	card2str := handStr[2:]
	card1 := ParseCard(card1str)
	card2 := ParseCard(card2str)
	return []Card{card1, card2}
}

func ParseWightedHand(handStr string) WeightedHand {
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
	return WeightedHand{hand, Weight(weight)}
}
