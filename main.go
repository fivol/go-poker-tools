package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-poker-equity/equity"
	"go-poker-equity/poker"
	"io"
	"io/ioutil"
	"os"
	"strings"
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
		panic("fail to dump json")
	}
	fmt.Println(string(resultJson))
}

func readRanges(input io.Reader, rangeLines *[]string) {
	file, err := ioutil.ReadAll(input)
	if err != nil {
		panic("reading ranges error: " + err.Error())
	}
	for _, line := range strings.Split(string(file), "\n") {
		rangeStr := strings.Trim(line, " \n\r")
		if rangeStr != "" {
			*rangeLines = append(*rangeLines, rangeStr)
		}
	}
}

func main() {
	var iterations = flag.Int("iter", 0, "iterations count (0 for unlimited)")
	var timeout = flag.Float64("timeout", 0, "timeout in seconds (fractional)")
	var rangesFile = flag.String("ranges", "", "path to file with ranges lines")
	flag.Parse()
	var rangeLines []string
	if flag.NArg() < 1 {
		panic("must specify board")
	}
	if flag.NArg() >= 3 {
		for _, rangeStr := range flag.Args()[1:] {
			rangeLines = append(rangeLines, rangeStr)
		}
	}
	if *rangesFile != "" {
		file, err := os.Open(*rangesFile)
		if err != nil {
			panic("ranges file open error: " + err.Error())
		}
		readRanges(file, &rangeLines)
	}
	if len(rangeLines) == 0 && *rangesFile == "" {
		readRanges(os.Stdin, &rangeLines)
	}
	if len(rangeLines) < 2 {
		panic("must specify at least 2 ranges")
	}

	if *iterations < 0 {
		panic("iterations must me grater or equal zero")
	}
	board := poker.ParseBoard(flag.Args()[0])

	var ranges []poker.Range
	for _, rangeStr := range rangeLines {
		range_ := poker.ParseRange(rangeStr)
		range_.RemoveCards(board...)
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
