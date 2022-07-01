package main

import "go-poker-equity/poker"

type HandEquity struct {
	Hand   poker.Hand
	Equity Equity
}

type EquityRange []HandEquity

func (h HandEquity) ToString() string {
	return h.ToString()
}
