package main

import (
	"flag"
	"fmt"
	"go-poker-equity/poker"
)

func printEquity(equityRange EquityRange) {
	for _, hand := range equityRange {
		fmt.Printf("%s\n", hand.ToString())
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 3 {
		panic("must be specified at least board and two ranges")
	}
	board := poker.ParseBoard(flag.Args()[0])
	var ranges []poker.Range
	for _, rangeStr := range flag.Args()[1:] {
		ranges = append(ranges, poker.ParseRange(rangeStr))
	}
	equityRange := CalculateEquity(board, ranges[0], ranges[1:])
	printEquity(equityRange)
}
