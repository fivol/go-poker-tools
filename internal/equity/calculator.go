package equity

import (
	"go-poker-tools/pkg/types"
	"go-poker-tools/pkg/winner"
	//wr "go-types-go-types-tools/random"
	wr "github.com/mroth/weightedrand"
	"math/rand"
	"time"
)

type equityResult struct {
	hand   types.Hand
	equity Equity
}

type equityCalculator struct {
	choosers      [types.RangeLen][]wr.Chooser
	oppRanges     []types.Range
	board         types.Board
	runIterations chan uint32
	done          chan bool
	result        chan equityResult
	runnersCount  int
}

func newEquityCalculator(board types.Board, oppRanges *[]types.Range) equityCalculator {
	return equityCalculator{
		oppRanges:     *oppRanges,
		board:         board,
		runIterations: make(chan uint32),
		result:        make(chan equityResult),
		done:          make(chan bool),
	}
}

func selectHand(chooser *wr.Chooser) types.Hand {
	return chooser.Pick().(types.Hand)
}

func (c *equityCalculator) createOppRangesChoosers(myRange *types.Range) {
	myRangeIter := types.NewRangeIterator(myRange)
	for myHand, _, ended := myRangeIter.Next(); !ended; myHand, _, ended = myRangeIter.Next() {
		for _, range_ := range c.oppRanges {
			var choices []wr.Choice
			oppRangeCopy := range_
			oppRangeCopy.RemoveCards(myHand.Cards())
			iter := types.NewRangeIterator(&oppRangeCopy)
			weightsSum := uint(0)
			for hand, weight, end := iter.Next(); !end; hand, weight, end = iter.Next() {
				weightInt := uint(weight * 1000)
				choices = append(choices, wr.Choice{Item: hand, Weight: weightInt})
				weightsSum += weightInt
			}
			if weightsSum == 0 {
				myRange[myHand] = 0
				continue
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

func (c *equityCalculator) sampleOppHands(myHand types.Hand, oppOrder []int) []types.Hand {
	hands := make([]types.Hand, len(c.choosers[myHand]))
	for i, oppIdx := range oppOrder {
		hands[i] = selectHand(&c.choosers[myHand][oppIdx])
	}
	return hands
}

func (c *equityCalculator) selectOppHands(myHand types.Hand) []types.Hand {
	oppsIdx := make([]int, len(c.choosers[myHand]))
	for i := range c.choosers[myHand] {
		oppsIdx[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(oppsIdx), func(i, j int) { oppsIdx[i], oppsIdx[j] = oppsIdx[j], oppsIdx[i] })
	hands := c.sampleOppHands(myHand, oppsIdx)
	if !types.IsDistinct(types.ToCards(hands...)...) {
		return c.selectOppHands(myHand)
	}
	return hands
}

func (c *equityCalculator) handsWinCheck(myHand types.Hand, oppHands []types.Hand) (bool, int) {
	hands := make([]types.Hand, len(oppHands)+1)
	copy(hands, oppHands)
	hands[len(oppHands)] = myHand
	winners := winner.DetermineWinners(c.board, hands)
	isWin := false
	for _, winner := range winners {
		if winner == len(hands)-1 {
			isWin = true
		}
	}
	return isWin, len(winners)
}

func (c *equityCalculator) iterHandWinCheck(hand types.Hand) (bool, int) {
	hands := c.selectOppHands(hand)
	hands = append(hands, hand)
	winners := winner.DetermineWinners(c.board, hands)
	isWin := false
	for _, winner := range winners {
		if winner == len(hands)-1 {
			isWin = true
		}
	}
	return isWin, len(winners)
}

func (c *equityCalculator) calcHandWinCount(hand types.Hand, iterations uint32) float32 {
	var winsCount float32
	for i := uint32(0); i < iterations; i++ {
		won, winnersCount := c.iterHandWinCheck(hand)
		if won {
			winsCount += 1.0 / float32(winnersCount)
		}
	}
	return winsCount
}

func runHandEquityCalc(calculator *equityCalculator, hand types.Hand) {
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

func calcHandEquity3way(params *RequestParams, calculator *equityCalculator, hand types.Hand) {
	opp1 := params.OppRanges[0]
	opp2 := params.OppRanges[1]
	opp1.RemoveCards(hand.Cards())
	opp2.RemoveCards(hand.Cards())
	hands := make([]types.Hand, 2)
	iter1 := types.NewRangeIterator(&opp1)

	var weightsAll float32
	var weightsWin float32

	for hand1, w1, end1 := iter1.Next(); !end1; hand1, w1, end1 = iter1.Next() {
		c1, c2 := hand1.Cards()
		iter2 := types.NewRangeIterator(&opp2)
		hands[0] = hand1
		for hand2, w2, end2 := iter2.Next(); !end2; hand2, w2, end2 = iter2.Next() {
			c3, c4 := hand2.Cards()
			if c1 == c3 || c1 == c4 || c2 == c3 || c2 == c4 {
				continue
			}
			hands[1] = hand2
			pairWeight := w1 * w2
			weightsAll += pairWeight
			win, winnersCount := calculator.handsWinCheck(hand, hands)
			if win {
				weightsWin += pairWeight / float32(winnersCount)
			}
			winner.DetermineWinners(params.Board, hands)
		}
	}
	res := equityResult{
		hand:   hand,
		equity: Equity(weightsWin / weightsAll),
	}
	calculator.result <- res
}

func calculateEquity3way(params *RequestParams) (res ResultData) {
	t0 := time.Now()
	res.Equity = make(map[types.Hand]Equity)

	calculator := newEquityCalculator(params.Board, &params.OppRanges)

	iter := types.NewRangeIterator(&params.MyRange)
	for hand, _, end := iter.Next(); !end; hand, _, end = iter.Next() {
		go calcHandEquity3way(params, &calculator, hand)
		calculator.runnersCount++
	}

	for i := 0; i < calculator.runnersCount; i++ {
		equity := <-calculator.result
		res.Equity[equity.hand] = equity.equity
	}

	res.TimeDelta = float32(time.Since(t0).Seconds())
	return
}

func CalculateEquity(params *RequestParams) (res ResultData) {
	if len(params.OppRanges) == 2 {
		return calculateEquity3way(params)
	}
	t0 := time.Now()
	res.Equity = make(map[types.Hand]Equity)
	initRandom()
	calculator := newEquityCalculator(params.Board, &params.OppRanges)
	calculator.createOppRangesChoosers(&params.MyRange)
	iter := types.NewRangeIterator(&params.MyRange)
	for hand, _, end := iter.Next(); !end; hand, _, end = iter.Next() {
		go runHandEquityCalc(&calculator, hand)
		calculator.runnersCount++
	}
	runUntilStop(&res, &calculator, params)
	res.TimeDelta = float32(time.Since(t0).Seconds())
	return
}
