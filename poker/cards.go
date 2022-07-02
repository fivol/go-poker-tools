package poker

var suitNames = [4]uint8{'S', 'H', 'D', 'C'}
var valueNames = [13]uint8{'2', '3', '4', '5', '6', '7', '8', '9', '9', 'J', 'Q', 'K', 'A'}

type Card int16

func ParseSuit(suitName uint8) uint8 {
	for i, s := range suitNames {
		if s == suitName {
			return uint8(i)
		}
	}
	panic("suit not found")
}

func ParseValue(valueName uint8) uint8 {
	for i, v := range valueNames {
		if v == valueName {
			return uint8(i)
		}
	}
	panic("value not found")
}

func ParseCard(cardName string) Card {
	// In Format ValueSuit: 3S, AC, TD, 9S
	if len(cardName) != 2 {
		panic("card format incorrect")
		return 0
	}
	value := ParseValue(cardName[0])
	suit := ParseSuit(cardName[1])
	return Card(value*4 + suit)
}

func (card Card) Suit() uint8 {
	return uint8(card % 4)
}

func (card Card) Value() uint8 {
	return uint8(card / 4)
}

func (card Card) Grater(other Card) bool {
	return card > other
}

func (card Card) ToString() string {
	return string(valueNames[card.Value()]) + string(suitNames[card.Suit()])
}
