package tests

import "go-poker-equity/poker"

func test() {
	card1 := poker.ParseCard("AS")
	card2 := poker.ParseCard("2S")
	card3 := poker.ParseCard("2H")
	println(card1.Suit() == card2.Suit())
	println(card1.Suit() == card2.Suit())
	println(card2.Value() == card3.Value())
	println(card1.Value() != card2.Value())

	//println(ToString(card1) == "AS")
	//println(ToString(card2) == "2S")
	//println(ToString(card3) == "2H")
}
