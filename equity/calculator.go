package equity

import (
	"go-poker-equity/combinations"
	"go-poker-equity/poker"
	//wr "go-poker-equity/random"
	wr "github.com/mroth/weightedrand"
	"math/rand"
	"time"
)

type equityResult struct {
	hand   poker.Hand
	equity Equity
}

type equityCalculator struct {
	choosers      [poker.RangeLen][]wr.Chooser
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

func selectHand(chooser *wr.Chooser) poker.Hand {
	return chooser.Pick().(poker.Hand)
}

func (c *equityCalculator) createOppRangesChoosers(myRange *poker.Range) {
	myRangeIter := poker.NewRangeIterator(myRange)
	for myHand, _, ended := myRangeIter.Next(); !ended; myHand, _, ended = myRangeIter.Next() {
		for _, range_ := range c.oppRanges {
			var choices []wr.Choice
			oppRangeCopy := range_
			oppRangeCopy.RemoveCards(myHand.Cards())
			iter := poker.NewRangeIterator(&oppRangeCopy)
			for hand, weight, end := iter.Next(); !end; hand, weight, end = iter.Next() {
				choices = append(choices, wr.Choice{Item: hand, Weight: uint(weight * 1000)})
			}
			chooser, err := wr.NewChooser(choices...)
			if err != nil {
				panic(err)
			}
			c.choosers[myHand] = append(c.choosers[myHand], *chooser)
		}
	}
}

func initRandom() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (c *equityCalculator) sampleOppHands(myHand poker.Hand) []poker.Hand {
	hands := make([]poker.Hand, len(c.choosers[myHand]))
	for i := 0; i < len(c.choosers[myHand]); i++ {
		hands[i] = selectHand(&c.choosers[myHand][i])
	}
	return hands
}

func (c *equityCalculator) selectOppHands(myHand poker.Hand) []poker.Hand {
	hands := c.sampleOppHands(myHand)
	cards := make([]poker.Card, len(hands)*2)
	for i, hand := range hands {
		c1, c2 := hand.Cards()
		cards[i*2] = c1
		cards[i*2+1] = c2
	}
	if !poker.IsDistinct(cards...) {
		return c.selectOppHands(myHand)
	}
	return hands
}

func (c *equityCalculator) iterHandWinCheck(hand poker.Hand) (bool, int) {
	hands := c.selectOppHands(hand)
	hands = append(hands, hand)
	winners := combinations.DetermineWinners(c.board, hands)
	isWin := false
	for _, winner := range winners {
		if winner == len(hands)-1 {
			isWin = true
		}
	}
	return isWin, len(winners)
}

func (c *equityCalculator) calcHandWinCount(hand poker.Hand, iterations uint32) float32 {
	var winsCount float32
	for i := uint32(0); i < iterations; i++ {
		won, winnersCount := c.iterHandWinCheck(hand)
		if won {
			winsCount += 1.0 / float32(winnersCount)
		}
	}
	return winsCount
}

func runHandEquityCalc(calculator *equityCalculator, hand poker.Hand) {
	var totalIterations uint32
	var totalWinCount float32
	for {
		iterations := <-calculator.runIterations
		if iterations == 0 {
			break
		}
		totalIterations += iterations
		totalWinCount += calculator.calcHandWinCount(hand, iterations)
		calculator.done <- true
	}
	equity := Equity(totalWinCount / float32(totalIterations))
	result := equityResult{
		equity: equity,
		hand:   hand,
	}
	calculator.result <- result
}

func runUntilStop(res *ResultData, calculator *equityCalculator, params *RequestParams) {
	t0 := time.Now()
	if params.Timeout == 0 && params.Iterations == 0 {
		panic("Limits not selected")
	}
	totalIterations := uint32(0)
	var iterationsBunch uint32 = 2000
	if params.Timeout == 0 {
		iterationsBunch = params.Iterations
	}
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
	calculator.createOppRangesChoosers(&params.MyRange)
	iter := poker.NewRangeIterator(&params.MyRange)
	for hand, _, end := iter.Next(); !end; hand, _, end = iter.Next() {
		go runHandEquityCalc(&calculator, hand)
		calculator.runnersCount++
	}
	runUntilStop(&res, &calculator, params)
	res.TimeDelta = float32(time.Since(t0).Seconds())
	return
}
