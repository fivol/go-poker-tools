package equity

import (
	wr "github.com/mroth/weightedrand"
	"go-poker-equity/combinations"
	"go-poker-equity/poker"
	"math/rand"
	"time"
)

type equityCalculator struct {
	choosers  []wr.Chooser
	oppRanges []poker.Range
	board     poker.Board
}

func newEquityCalculator(board poker.Board, oppRanges *[]poker.Range) equityCalculator {
	return equityCalculator{oppRanges: *oppRanges, board: board}
}

func selectHand(chooser *wr.Chooser) poker.Hand {
	return chooser.Pick().(poker.Hand)
}

func (c *equityCalculator) createOppRangesChoosers() {
	for _, range_ := range c.oppRanges {
		var choices []wr.Choice
		iter := poker.NewRangeIterator(&range_)
		for hand, weight, end := iter.Next(); !end; {
			choices = append(choices, wr.Choice{Item: hand, Weight: uint(weight * 1000)})
		}
		chooser, err := wr.NewChooser(choices...)
		if err != nil {
			panic(err)
		}
		c.choosers = append(c.choosers, *chooser)
	}
}

func initRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (c *equityCalculator) selectOppHands() []poker.Hand {
	var hands []poker.Hand
	for i := 0; i < len(c.choosers); i++ {
		hands = append(hands, selectHand(&c.choosers[i]))
	}
	return hands
}

func (c *equityCalculator) iterHandWinCheck(hand poker.Hand) bool {
	hands := c.selectOppHands()
	hands = append(hands, hand)
	winners := combinations.DetermineWinners(c.board, hands)
	for _, winner := range winners {
		if winner == hand {
			return true
		}
	}
	return false
}

func (c *equityCalculator) calcHandEquity(hand poker.Hand, iterations uint32) Equity {
	var winsCount uint32
	for i := uint32(0); i < iterations; i++ {
		if c.iterHandWinCheck(hand) {
			winsCount++
		}
	}
	return Equity(float32(winsCount) / float32(iterations))
}

func CalculateEquity(board poker.Board, myRange poker.Range, oppRanges []poker.Range, iterations uint32) []HandEquity {
	initRandom()
	calculator := newEquityCalculator(board, &oppRanges)
	calculator.createOppRangesChoosers()
	var equityRange []HandEquity
	for hand, _ := range myRange {
		equity := calculator.calcHandEquity(poker.Hand(hand), iterations)
		equityRange = append(equityRange, HandEquity{Hand: poker.Hand(hand), Equity: equity})
	}
	return equityRange
}
