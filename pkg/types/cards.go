package types

import "fmt"

var suitNames = [4]uint8{'s', 'h', 'd', 'c'}
var valueNames = [13]uint8{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

type Card int16
type Value uint8
type Suit uint8

func IsDistinct(cards ...Card) bool {
	var passed [52]bool
	for _, card := range cards {
		if passed[card] {
			return false
		}
		passed[card] = true
	}
	return true
}

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
	// In Format ValueSuit: 3s, Ac, Td, 9s
	if len(cardName) == 0 {
		panic("can not parse empty card")
	}
	if len(cardName) != 2 {
		panic(fmt.Sprintf("card format incorrect: %v", cardName))
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
