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
	startMs   int64
}

func newEquityCalculator(board poker.Board, oppRanges *[]poker.Range) equityCalculator {
	return equityCalculator{oppRanges: *oppRanges, board: board, startMs: time.Now().UnixMilli()}
}

func selectHand(chooser *wr.Chooser) poker.Hand {
	return chooser.Pick().(poker.Hand)
}

func (c *equityCalculator) createOppRangesChoosers() {
	for _, range_ := range c.oppRanges {
		var choices []wr.Choice
		iter := poker.NewRangeIterator(&range_)
		for hand, weight, end := iter.Next(); !end; hand, weight, end = iter.Next() {
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
		if winner == len(hands)-1 {
			return true
		}
	}
	return false
}

func (c *equityCalculator) calcHandEquity(hand poker.Hand, params *RequestParams) (Equity, uint32) {
	var winsCount uint32
	var iterationsDone uint32
	for i := uint32(0); params.Iterations == 0 || i < params.Iterations; i++ {
		if c.iterHandWinCheck(hand) {
			winsCount++
		}
		iterationsDone++
		if i%100 == 0 && params.Timeout > 0 {
			if time.Now().UnixMilli()-c.startMs >= int64(params.Timeout*1000) {
				break
			}
		}
	}
	return Equity(float32(winsCount) / float32(iterationsDone)), iterationsDone
}

func addHandEquity(res *ResultData, hand poker.Hand, eq Equity, iterations uint32) {
	res.Equity[hand] = eq
	if res.Iterations == 0 {
		res.Iterations = iterations
	} else {
		res.Iterations = (res.Iterations + iterations) / 2
	}
}

func runHandEquityCalc(params *RequestParams, res *ResultData, calculator *equityCalculator, hand poker.Hand) {
	equity, iterations := calculator.calcHandEquity(hand, params)
	addHandEquity(res, hand, equity, iterations)
}

func CalculateEquity(params *RequestParams) (res ResultData) {
	res.Equity = make(map[poker.Hand]Equity)
	initRandom()
	calculator := newEquityCalculator(params.Board, &params.OppRanges)
	calculator.createOppRangesChoosers()
	iter := poker.NewRangeIterator(&params.MyRange)
	for hand, _, end := iter.Next(); !end; hand, _, end = iter.Next() {
		if params.Board.Intersects(hand) {
			continue
		}
		go runHandEquityCalc(params, &res, &calculator, hand)
	}
	res.TimeDelta = float32(time.Now().UnixMilli()-calculator.startMs) / 1000
	return
}
