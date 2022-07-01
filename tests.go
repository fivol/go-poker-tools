package main

import "go-poker-equity/poker"

func test() {
	card1 := poker.ParseCard("AS")
	card2 := poker.ParseCard("2S")
	card3 := poker.ParseCard("2H")
	println(poker.Suit(card1) == poker.Suit(card2))
	println(poker.Suit(card1) == poker.Suit(card2))
	println(poker.Value(card2) == poker.Value(card3))
	println(poker.Value(card1) != poker.Value(card2))

	//println(ToString(card1) == "AS")
	//println(ToString(card2) == "2S")
	//println(ToString(card3) == "2H")
}
