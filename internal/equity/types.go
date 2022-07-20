package equity

import (
	"fmt"
	"go-poker-tools/pkg/types"
)

type HandEquity struct {
	Hand   types.Hand
	Equity Equity
}
type Equity float32

type RequestParams struct {
	Board      types.Board
	MyRange    types.Range
	OppRanges  []types.Range
	Iterations uint32
	Timeout    float64
}

type ResultData struct {
	Equity     map[types.Hand]Equity
	Iterations uint32
	TimeDelta  float32
}

func (h HandEquity) ToString() string {
	return fmt.Sprintf("%s %f\n", h.Hand.ToString(), h.Equity)
}
