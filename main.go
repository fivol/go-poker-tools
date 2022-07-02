package main

import (
	"flag"
	"fmt"
	"go-poker-equity/equity"
	"go-poker-equity/poker"
)

func printEquity(equityRange []equity.HandEquity) {
	for _, hand := range equityRange {
		fmt.Printf("%s\n", hand.ToString())
	}
}

func main() {
	flag.Parse()
	if flag.NArg() < 3 {
		panic("must be specified at least board and two ranges")
	}
	var iterations = flag.Int("n", 1234, "iterations count")
	if *iterations <= 0 {
		panic("iterations must me grater then zero")
	}
	board := poker.ParseBoard(flag.Args()[0])
	var ranges []poker.Range
	for _, rangeStr := range flag.Args()[1:] {
		ranges = append(ranges, poker.ParseRange(rangeStr))
	}
	equityRange := equity.CalculateEquity(board, ranges[0], ranges[1:], uint32(*iterations))
	printEquity(equityRange)
}
