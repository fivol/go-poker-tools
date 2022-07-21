package combinations

import (
	"go-poker-tools/pkg/combinations"
	"go-poker-tools/pkg/types"
)

func GetHandsCombinations(board types.Board, hands []types.Hand) map[types.Hand][]combinations.Comb {
	result := make(map[types.Hand][]combinations.Comb)
	for _, hand := range hands {
		result[hand] = combinations.GetCombinations(board, hand)
	}
	return result
}

func HandsByCombination(handsCombos map[types.Hand][]combinations.Comb) map[combinations.Comb]types.HandsList {
	result := make(map[combinations.Comb]types.HandsList)
	for hand, combos := range handsCombos {
		for _, comb := range combos {
			result[comb] = append(result[comb], hand)
		}
	}
	return result
}
