package poker

import (
	"strings"
)

type Board []Card

func ParseBoard(boardStr string) Board {
	cardsStr := strings.Split(boardStr, ",")
	if len(cardsStr) > 5 {
		panic("board must contain less then 5 cards")
	}
	if len(cardsStr) < 3 {
		panic("board must contain at least 3 cards")
	}
	var cards []Card
	for _, cardStr := range cardsStr {
		cards = append(cards, ParseCard(cardStr))
	}
	return cards
}
