package types

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCardParsing(t *testing.T) {
	cardsToCheck := []string{
		"As", "Td", "8h", "Jc", "Js", "2d", "Ts", "Ah", "4s", "8d", "Qc",
	}
	for _, card := range cardsToCheck {
		assert.Equal(t, card, ParseCard(card).ToString(),
			"Cards not equal", card, "!=", ParseCard(card).ToString())
	}
}

func TestHandParsing(t *testing.T) {
	handsToCheck := []string{
		"AsAd", "Td4h", "8hQc", "Jc3c", "JsJh", "2d7s", "TsAs", "Ah3c", "4s5c", "8d8h", "QcKd",
	}
	for _, hand := range handsToCheck {
		assert.Contains(t, []string{hand, hand[2:] + hand[:2]}, ParseHand(hand).ToString(),
			"Hands not equal", hand, "!=", ParseHand(hand).ToString())
	}
	assert.Panics(t, func() { ParseHand("AsAs") }, "Passed same cards to hand")
	assert.Panics(t, func() { ParseHand("2h2h") }, "Passed same cards to hand")
	assert.Panics(t, func() { ParseHand("TdTd") }, "Passed same cards to hand")
}

func TestBoardParsing(t *testing.T) {
	boardsToCheck := []string{
		"AsAdTd", "Td4h8h", "8hQcJc", "JsJhJd", "2d7sQcKd", "TsAsAh3c4s", "Ts8h2h9h",
	}
	for _, board := range boardsToCheck {
		assert.Equal(t, ParseCard(board[:2]), ParseBoard(board)[0], "First card no equal")
		lastCard := board[len(board)-2:]
		assert.Equal(t, ParseCard(lastCard), ParseBoard(board)[len(board)/2-1], "First card no equal")
	}
	assert.Panics(t, func() { ParseBoard("AsAdAs") }, "Passed same cards to board")
	assert.Panics(t, func() { ParseBoard("2s3s4s2s") }, "Passed same cards to board")
	assert.Panics(t, func() { ParseBoard("Ts8h2h8h") }, "Passed same cards to board")
}

func TestSuitsValues(t *testing.T) {
	cards := []string{
		"As", "Td", "8h", "Jc", "Js", "2d", "Ts", "Ah", "4s", "8d", "Qc",
	}
	for _, card := range cards {
		suit := card[1]
		value := card[0]
		assert.Equal(t, ParseSuit(suit), ParseCard(card).Suit(), "Suits does not match")
		assert.Equal(t, ParseValue(value), ParseCard(card).Value(), "Suits does not match")
	}
}

func TestParseRange(t *testing.T) {
	ranges := []string{
		"AsAd:0.34,2s2c:0.34",
	}
	for _, rangeStr := range ranges {
		range_ := ParseRange(rangeStr)
		iter := range_.GetIterator()
		for hand, w, end := iter.Next(); !end; hand, w, end = iter.Next() {
			assert.Equal(t, float32(0.34), w, "Weight must me 1")
			assert.True(t, strings.Contains(rangeStr, hand.ToString()), "Unknown hand in range")
		}
	}
}

func TestIsDistinct(t *testing.T) {
	assert.True(t, IsDistinct(parseCards("7hAsQdJs8hThAc")...), "Cards not distinct fail")
	assert.True(t, IsDistinct(parseCards("7hAsQdJs8h4hQc")...), "Cards not distinct fail")
	assert.True(t, IsDistinct(parseCards("7hAsQdJs8hQcKh")...), "Cards not distinct fail")

}
