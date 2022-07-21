package main

import (
	"encoding/json"
	"fmt"
	"go-poker-tools/internal/combinations"
	"go-poker-tools/pkg/types"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

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

func main() {
	if len(os.Args) < 2 {
		panic("Must specify board with first argument")
	}
	t0 := time.Now()
	board := types.ParseBoard(os.Args[1])
	var hands types.HandsList
	if len(os.Args) >= 3 {
		hands = readHands(strings.NewReader(os.Args[2]), board)
	} else {
		hands = readHands(os.Stdin, board)
	}

	handsCombos := combinations.GetHandsCombinations(board, hands)
	combosHands := combinations.HandsByCombination(handsCombos)
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
