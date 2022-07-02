package equity

import (
	"go-poker-equity/poker"
)

type HandEquity struct {
	Hand   poker.Hand
	Equity Equity
}

type Equity float32

func (h HandEquity) ToString() string {
	return h.ToString()
}
