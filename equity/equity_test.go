package equity

import (
	"github.com/stretchr/testify/assert"
	"go-poker-equity/poker"
	"testing"
)

func TestCalculateEquity(t *testing.T) {
	table := []struct {
		board      string
		myRange    string
		ranges     []string
		iterations int
		equity     map[string]float32
	}{{
		"6s9c4hQcKd",
		"KsKc",
		[]string{"2s3h"},
		1,
		map[string]float32{"KsKc": 1},
	},
		{
			"6s9c4hQcKd",
			"2s3h",
			[]string{"6h3s"},
			1,
			map[string]float32{"2s3h": 0},
		},
	}

	for _, testCase := range table {
		board := poker.ParseBoard(testCase.board)
		var ranges []poker.Range
		for _, rangeStr := range testCase.ranges {
			ranges = append(ranges, poker.ParseRange(rangeStr))
		}
		params := RequestParams{
			Board:      board,
			MyRange:    poker.ParseRange(testCase.myRange),
			OppRanges:  ranges,
			Iterations: uint32(testCase.iterations),
			Timeout:    1,
		}
		result := CalculateEquity(&params)
		for hand, equity := range testCase.equity {
			assert.Equal(t, equity, float32(result.Equity[poker.ParseHand(hand)]), "Equity not match")
		}
	}
}
