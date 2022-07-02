package poker

type Board []Card

func parseCards(cardsStr string) []Card {
	var cards []Card
	for i := 0; i < len(cardsStr); i += 2 {
		card := ParseCard(cardsStr[i : i+2])
		cards = append(cards, card)
	}
	return cards
}

func ParseBoard(boardStr string) Board {
	cards := parseCards(boardStr)
	if len(cards) > 5 {
		panic("board must contain less then 5 cards")
	}
	if len(cards) < 3 {
		panic("board must contain at least 3 cards")
	}
	if !IsDistinct(cards...) {
		panic("board cards repeats")
	}
	return cards
}
