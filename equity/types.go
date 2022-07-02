package equity

import (
	"fmt"
	"go-poker-equity/poker"
)

type HandEquity struct {
	Hand   poker.Hand
	Equity Equity
}
type Equity float32

type RequestParams struct {
	Board      poker.Board
	MyRange    poker.Range
	OppRanges  []poker.Range
	Iterations uint32
	Timeout    float32
}

type ResultData struct {
	Equity     map[poker.Hand]Equity
	Iterations uint32
	TimeDelta  float32
}

func (h HandEquity) ToString() string {
	return fmt.Sprintf("%s %f\n", h.Hand.ToString(), h.Equity)
}
