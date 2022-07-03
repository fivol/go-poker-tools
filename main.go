package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-poker-equity/equity"
	"go-poker-equity/poker"
)

type EquityResultModel struct {
	Equity     map[string]equity.Equity `json:"equity"`
	TimeDelta  float32                  `json:"time_delta"`
	Iterations uint32                   `json:"iterations"`
}

func printResults(result equity.ResultData) {
	outputModel := EquityResultModel{}
	outputModel.Equity = make(map[string]equity.Equity)
	outputModel.Iterations = result.Iterations
	outputModel.TimeDelta = result.TimeDelta
	for key := range result.Equity {
		outputModel.Equity[key.ToString()] = result.Equity[key]
	}
	resultJson, err := json.Marshal(&outputModel)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resultJson))
}

func main() {
	var iterations = flag.Int("iter", 0, "iterations count (0 for unlimited)")
	var timeout = flag.Float64("timeout", 0, "timeout in seconds (fractional)")
	flag.Parse()
	if flag.NArg() < 3 {
		panic("must be specified at least board and two ranges")
	}

	if *iterations < 0 {
		panic("iterations must me grater or equal zero")
	}
	board := poker.ParseBoard(flag.Args()[0])
	var ranges []poker.Range
	for _, rangeStr := range flag.Args()[1:] {
		range_ := poker.ParseRange(rangeStr)
		range_.RemoveCards(board)
		ranges = append(ranges, range_)
	}
	if *iterations == 0 && *timeout == 0 {
		panic("infinite run, set at least timeout or iterations")
	}
	params := equity.RequestParams{
		Board:      board,
		MyRange:    ranges[0],
		OppRanges:  ranges[1:],
		Iterations: uint32(*iterations),
		Timeout:    *timeout,
	}
	result := equity.CalculateEquity(&params)
	printResults(result)
}
