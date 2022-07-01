package poker

var suitNames = [4]uint8{'S', 'H', 'D', 'C'}
var valueNames = [13]uint8{'2', '3', '4', '5', '6', '7', '8', '9', '9', 'J', 'Q', 'K', 'A'}

type Card int64

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
	suit := ParseSuit(cardName[1])
	value := ParseValue(cardName[0])
	return Card(value*4 + suit)
}

func Suit(card Card) int8 {
	return int8(card % 4)
}

func Value(card Card) int8 {
	return int8(card / 4)
}

func (card Card) ToString() string {
	return string(valueNames[Value(card)]) + string(suitNames[Suit(card)])
}
