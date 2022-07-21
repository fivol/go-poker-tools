package combinations

import "go-poker-tools/pkg/types"

type Comb string

type CombinationExtractor func(c *Selector) bool

type ExtractorWithName struct {
	extractor CombinationExtractor
	name      Comb
}

var combinations = []ExtractorWithName{
	{findStraightFlush, "straight_flush"},
	{findStraight, "straight"},
	{findFlush, "flush"},
	{findSet, "set"},
	{findTopSet, "top_set"},
	{findMediumSet, "medium_set"},
	{findTopTwoPairs, "top_two_pairs"},
	{findMediumTwoPairs, "medium_two_pairs"},
	{findOverPairOESD, "overpair_oesd"},
	{findOverPairGSH, "overpair_gsh"},
	{findHighOverPair, "high_overpair"},
	{findLowOverPair, "low_overpair"},
	{findHighOverPairFD, "high_overpair_fd"},
	{findLowOverPairFD, "low_overpair_fd"},
	{findTPFDNutsFd, "tp_fd_nuts_fd"},
	{findTPFD, "tp_fd"},
	{findTPOESD, "tp_oesd"},
	{findTPGSH, "tp_gsh"},
	{findTP, "tp"},
	{findPocketTP2FD13Nuts, "pocket_tp_2_fd_1_3_nuts"},
	{findPocketTP2FD4Nuts, "pocket_tp_2_fd_4_nuts"},
	{findPocketTP2OESD, "pocket_tp_2_oesd"},
	{findPocketTP2GSH, "pocket_tp_2_gsh"},
	{findPocketTop2, "pocket_top_2"},
	{findSecondFD13Nuts, "2nd_fd_1_3_nuts"},
	{findSecondFD4Nuts, "2nd_2nd_fd_4_nuts"},
	{findSecondOESD, "2nd_oesd"},
	{findSecondGSH, "2nd_gsh"},
	{findSecond, "2nd"},
	{findPair, "pair"},
	{findPocketBetween23FDNuts, "pocket_between_2_3_fd_nuts"},
	{findPocketBetween23FD23Nuts, "pocket_between_2_3_fd_2_3_nuts"},
	{findPocketBetween23FD4Nuts, "pocket_between_2_3_fd_4_nuts"},
	{findPocketBetween23OESD, "pocket_between_2_3_oesd"},
	{findPocketBetween23GSH, "pocket_between_2_3_gsh"},
	{findPocketBetween23, "pocket_between_2_3"},
	{find3dHands, "3d_hands"},
	{find3dHandsFDNuts, "3d_hands_fd_nuts"},
	{find3dHandsFD23Nuts, "3d_hands_fd_2_3_nuts"},
	{find3dHandsFD4Nuts, "3d_hands_fd_4_nuts"},
	{find3dHandsOESD, "3d_hands_oesd"},
	{find3dHandsGSH, "3d_hands_gsh"},
	{findUnderPocket, "under_pocket"},
	{findUnderPocketFD12Nuts, "under_pocket_fd_1_2_nuts"},
	{findAHigh, "ahigh"},
	{findNomade, "nomade"},
	{findTopCards, "top_cards"},
	{findOverCards, "overcards"},
	{findFdOESD2Cards, "fd_oesd_fd_2_cards"},
	{findGSHFD2Cards, "fd_gsh_fd_2_cards"},
	{findFDNutsFD, "fd_nuts_fd"},
	{findFD2nd3dNutsFD, "fd_2nd_3d_nuts_fd"},
	{findFD4NutsFD, "fd_4_nuts_fd"},
	{findBadOESD, "bad_oesd"},
	{findGoodOESD, "good_oesd"},
	{findGoodGutShot, "good_gutshot"},
	{findBadGutShot, "bad_gutshot"},
}

func GetCombinations(board types.Board, hand types.Hand) []Comb {
	selector := newCombinationsSelector(board, hand)
	var combs []Comb
	for _, combData := range combinations {
		if combData.extractor(&selector) {
			combs = append(combs, combData.name)
		}
	}
	return combs
}
