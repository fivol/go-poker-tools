package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"go-poker-tools/pkg/combinations"
	"go-poker-tools/pkg/types"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func GetHandsCombinations(board types.Board, hands []types.Hand, extractors []combinations.ExtractorWithName) map[types.Hand]combinations.Comb {
	result := make(map[types.Hand]combinations.Comb)
	for _, hand := range hands {
		result[hand] = combinations.GetCombinations(board, hand, extractors)
	}
	return result
}

func HandsByCombination(handsCombos map[types.Hand]combinations.Comb) map[combinations.Comb]types.HandsList {
	result := make(map[combinations.Comb]types.HandsList)
	for hand, comb := range handsCombos {
		result[comb] = append(result[comb], hand)
	}
	return result
}

type ResultModel struct {
	HandsByCombination map[string][]string `json:"combinations"`
	TimeDelta          float64             `json:"time_delta"`
}

func printResults(result ResultModel) {
	resultJson, err := json.Marshal(&result)
	if err != nil {
		panic("fail to dump json")
	}
	fmt.Println(string(resultJson))
}

func readHands(input io.Reader, board types.Board) types.HandsList {
	file, err := ioutil.ReadAll(input)
	if err != nil {
		panic("reading ranges error: " + err.Error())
	}
	handsStr := strings.Trim(string(file), " \n\r")
	if handsStr == "" {
		panic("have no hands")
	}
	var handsList types.HandsList
	hands := types.ParseRange(handsStr)
	hands.RemoveCards(board...)
	iter := types.NewRangeIterator(&hands)
	for hand, _, end := iter.Next(); !end; hand, _, end = iter.Next() {
		handsList = append(handsList, hand)
	}
	return handsList
}

func printCombos() {
	combos := combinations.GetAllCombos()
	for _, comb := range combos {
		fmt.Println(comb)
	}
}

func main() {
	var combosArg = flag.String("combos", "", "list of combos to use: trips,top_set,medium_set (use all by default)")
	var handsArg = flag.String("hands", "", "list of hands: 2s2d,KcQd")
	flag.Parse()
	if flag.NArg() < 1 {
		panic("Must specify board with first argument")
	}
	t0 := time.Now()
	if flag.Args()[0] == "combinations" {
		printCombos()
		return
	}
	board := types.ParseBoard(flag.Args()[0])
	var hands types.HandsList
	if *handsArg != "" {
		hands = readHands(strings.NewReader(*handsArg), board)
	} else {
		hands = readHands(os.Stdin, board)
	}
	combos := combinations.GetAllCombos()
	if *combosArg != "" {
		combos = []combinations.Comb{}
		for _, comb := range strings.Split(*combosArg, ",") {
			combos = append(combos, combinations.Comb(comb))
		}
	}
	extractors := combinations.GetExtractors(combos)
	handsCombos := GetHandsCombinations(board, hands, extractors)
	combosHands := HandsByCombination(handsCombos)
	handsByCombination := make(map[string][]string)
	for comb, handsList := range combosHands {
		for _, hand := range handsList {
			handsByCombination[string(comb)] = append(handsByCombination[string(comb)], hand.ToString())
		}
	}
	printResults(ResultModel{
		handsByCombination,
		time.Since(t0).Seconds(),
	})
}
