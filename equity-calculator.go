package main

import (
	wr "github.com/mroth/weightedrand"
	"go-poker-equity/poker"
	"math/rand"
	"time"
)

type equityCalculator struct {
	choosers  []wr.Chooser
	oppRanges []poker.Range
}

type Equity float32

func newEquityCalculator(oppRanges *[]poker.Range) equityCalculator {
	return equityCalculator{oppRanges: *oppRanges}
}

func selectHand(chooser *wr.Chooser) poker.Hand {
	return chooser.Pick().(poker.Hand)
}

func (c *equityCalculator) createOppRangesChoosers() {
	for _, range_ := range c.oppRanges {
		var choices []wr.Choice
		for _, hand := range range_ {
			choices = append(choices, wr.Choice{Item: hand.Hand, Weight: uint(hand.Weight * 1000)})
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

func (c *equityCalculator) selectOppCards() []poker.Hand {
	var hands []poker.Hand
	for i := 0; i < len(c.choosers); i++ {
		hands = append(hands, selectHand(&c.choosers[i]))
	}
	return hands
}

func (c *equityCalculator) calcHandEquity(hand poker.Hand) Equity {
	oppCards := c.selectOppCards()
}

func CalculateEquity(board poker.Board, myRange poker.Range, oppRanges []poker.Range) EquityRange {
	initRandom()
	calculator := newEquityCalculator(&oppRanges)
	calculator.createOppRangesChoosers()
	var equityRange EquityRange
	for _, hand := range myRange {
		equity := calculator.calcHandEquity(hand.Hand)
		equityRange = append(equityRange, HandEquity{Hand: hand.Hand, Equity: equity})
	}
	return equityRange
}
