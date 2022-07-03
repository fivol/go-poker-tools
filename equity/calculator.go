package equity

import (
	"github.com/jmcvetta/randutil"
	"go-poker-equity/combinations"
	"go-poker-equity/poker"
	"math/rand"
	"time"
)

type equityResult struct {
	hand   poker.Hand
	equity Equity
}

type equityCalculator struct {
	choosers      [][]randutil.Choice
	oppRanges     []poker.Range
	board         poker.Board
	runIterations chan uint32
	done          chan bool
	result        chan equityResult
	runnersCount  int
}

func newEquityCalculator(board poker.Board, oppRanges *[]poker.Range) equityCalculator {
	return equityCalculator{
		oppRanges:     *oppRanges,
		board:         board,
		runIterations: make(chan uint32),
		result:        make(chan equityResult),
		done:          make(chan bool),
	}
}

func (c *equityCalculator) createOppRangesChoosers() {
	for _, range_ := range c.oppRanges {
		choices := make([]randutil.Choice, 0, 2)
		iter := poker.NewRangeIterator(&range_)
		for hand, weight, end := iter.Next(); !end; hand, weight, end = iter.Next() {
			choices = append(choices, randutil.Choice{int(weight * 1000), hand})
		}
		c.choosers = append(c.choosers, choices)
	}
}

func initRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (c *equityCalculator) selectOppHands() []poker.Hand {
	var hands []poker.Hand
	for i := 0; i < len(c.choosers); i++ {
		result, _ := randutil.WeightedChoice(c.choosers[i])
		hands = append(hands, result.Item.(poker.Hand))
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

func (c *equityCalculator) calcHandWinCount(hand poker.Hand, iterations uint32) uint32 {
	var winsCount uint32
	for i := uint32(0); i < iterations; i++ {
		if c.iterHandWinCheck(hand) {
			winsCount++
		}
	}
	return winsCount
}

func runHandEquityCalc(calculator *equityCalculator, hand poker.Hand) {
	var totalIterations uint32
	var totalWinCount uint32
	for {
		iterations := <-calculator.runIterations
		if iterations == 0 {
			break
		}
		totalIterations += iterations
		totalWinCount += calculator.calcHandWinCount(hand, iterations)
		calculator.done <- true
	}
	equity := Equity(float32(totalWinCount) / float32(totalIterations))
	result := equityResult{
		equity: equity,
		hand:   hand,
	}
	calculator.result <- result
}

func runUntilStop(res *ResultData, calculator *equityCalculator, params *RequestParams) {
	t0 := time.Now()
	totalIterations := uint32(0)
	var iterationsBunch uint32 = 2000
	for {
		if params.Iterations > 0 && iterationsBunch+totalIterations > params.Iterations {
			iterationsBunch = params.Iterations - totalIterations
		}
		if params.Timeout > 0 && time.Since(t0).Seconds() >= params.Timeout {
			break
		}
		if params.Iterations > 0 && iterationsBunch == 0 {
			break
		}
		for i := 0; i < calculator.runnersCount; i++ {
			calculator.runIterations <- iterationsBunch
		}
		totalIterations += iterationsBunch
		for i := 0; i < calculator.runnersCount; i++ {
			<-calculator.done
		}
	}
	for i := 0; i < calculator.runnersCount; i++ {
		calculator.runIterations <- 0
	}
	for i := 0; i < calculator.runnersCount; i++ {
		equity := <-calculator.result
		res.Equity[equity.hand] = equity.equity
	}
	res.Iterations = totalIterations
}

func CalculateEquity(params *RequestParams) (res ResultData) {
	t0 := time.Now()
	res.Equity = make(map[poker.Hand]Equity)
	initRandom()
	calculator := newEquityCalculator(params.Board, &params.OppRanges)
	calculator.createOppRangesChoosers()
	iter := poker.NewRangeIterator(&params.MyRange)
	for hand, _, end := iter.Next(); !end; hand, _, end = iter.Next() {
		go runHandEquityCalc(&calculator, hand)
		calculator.runnersCount++
	}
	runUntilStop(&res, &calculator, params)
	res.TimeDelta = float32(time.Since(t0).Seconds())
	return
}
