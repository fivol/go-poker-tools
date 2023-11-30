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
	{findQuads, "quads"},
	{findFullHouse, "full_house"},
	{findHighFlushJ, "high_flush_j"},
	{findLowFlush, "low_flush"},
	{findFlush, "flush"},
	{findHighStraight, "high_straight"},
	{findLowStraight, "low_straight"},
	{findStraight, "straight"},
	{findTopSet, "top_set"},
	{findMediumSet, "medium_set"},
	{findSet, "set"},
	{findTrips, "trips"},
	{findTopTwoPairs, "top_two_pairs"},
	{findMediumTwoPairs, "medium_two_pairs"},
	{findTwoPairs, "two_pairs"},
	{findOverpairFDNuts, "high_overpair_fd_nuts"},
	{findOverpairFD2Nuts, "high_overpair_fd_2nuts"},
	{findOverpairFD3Nuts, "high_overpair_fd_3nuts"},
	{findOverpairFD4Nuts, "high_overpair_fd_4nuts"},
	// {findHighOverPairFD, "high_overpair_fd"},
	{findLowOverPairFD, "low_overpair_fd"},
	{findOverPairFD, "overpair_fd"},
	{findOverPairOESD, "overpair_oesd"},
	{findOverPairGSH, "overpair_gsh"},
	{findHighOverPair, "high_overpair"},
	{findLowOverPair, "low_overpair"},
	{findOverPair, "overpair"},
	// {findTPFDNutsFd, "tp_fd_nuts_fd"},
	{findTPFDNuts, "tp_fd_nuts"},
	{findTPFD2Nuts, "tp_fd_2nuts"},
	{findTPFD3Nuts, "tp_fd_3nuts"},
	{findTPFD4Nuts, "tp_fd_4nuts"},
	{findTPFD, "tp_fd"},
	{findTPOESD, "tp_oesd"},
	{findTPGSH, "tp_gsh"},
	{findTPBDFDNuts, "tp_bdfd_nuts"},
	{findOldTP, "old_tp"},
	{findNewTP, "new_tp"},
	{findTP, "tp"},
	{findPocketTP2FDNuts, "pocket_tp_2_fd_nuts"},
	{findPocketTP2FD2Nuts, "pocket_tp_2_fd_2_nuts"},
	{findPocketTP2FD3Nuts, "pocket_tp_2_fd_3_nuts"},
	{findPocketTP2FD4Nuts, "pocket_tp_2_fd_4_nuts"},
	// {findPocketTP2FD13Nuts, "pocket_tp_2_fd_1_3_nuts"},
	{findPocketTP2FD, "pocket_tp_2_fd"},
	{findPocketTP2OESD, "pocket_tp_2_oesd"},
	{findPocketTP2GSH, "pocket_tp_2_gsh"},
	{findPocketTop2, "pocket_tp_2"},
	{find2FDNuts, "2nd_fd_nuts"},
	{find2FD2Nuts, "2nd_fd_2_nuts"},
	{find2FD3Nuts, "2nd_fd_3_nuts"},
	// {findSecondFD13Nuts, "2nd_fd_1_3_nuts"},
	{findFdOESD2Cards, "fd_oesd_fd_2_cards"},
	{findGutShotFD2Cards, "fd_gsh_fd_2_cards"},
	{findFDOESDFD1Card, "fd_oesd_fd_1_card"},
	{findGutShotFD1Cards, "fd_gsh_fd_1_card"},
	{findSecondFD4Nuts, "2nd_fd_4_nuts"},
	{findSecondFD, "2nd_fd"},
	{findSecondOESD, "2nd_oesd"},
	{findSecondGSH, "2nd_gsh"},
	{find2ndBDFDNuts, "2nd_bdfd_nuts"},
	{findSecond, "2nd"},
	{findPocketBetween23FDNuts, "pocket_between_2_3_fd_nuts"},
	// {findPocketBetween23FD23Nuts, "pocket_between_2_3_fd_2_3_nuts"},
	{findPocketBetween23FD2Nuts, "pocket_between_2_3_fd_2_nuts"},
	{findPocketBetween23FD3Nuts, "pocket_between_2_3_fd_3_nuts"},
	{findPocketBetween23FD4Nuts, "pocket_between_2_3_fd_4_nuts"},
	{findPocketBetween23FD, "pocket_between_2_3_fd"},
	{findPocketBetween23OESD, "pocket_between_2_3_oesd"},
	{findPocketBetween23GSH, "pocket_between_2_3_gsh"},
	{findPocketBetween23, "pocket_between_2_3"},
	{find3dHandsFDNuts, "3d_hands_fd_nuts"},
	{find3dHandsFD2Nuts, "3d_hands_fd_2_nuts"},
	{find3dHandsFD3Nuts, "3d_hands_fd_3_nuts"},
	// {find3dHandsFD23Nuts, "3d_hands_fd_2_3_nuts"},
	{find3dHandsFD4Nuts, "3d_hands_fd_4_nuts"},
	{find3dHandsFD, "3d_hands_fd"},
	{find3dHandsOESD, "3d_hands_oesd"},
	{find3dHandsGSH, "3d_hands_gsh"},
	{find3ndBDFDNuts, "3nd_bdfd_nuts"},
	{find3dHands, "3d_hands"},
	{findUnderPocketFD12Nuts, "under_pocket_fd_1_2_nuts"},
	{findFDNutsFD, "fd_nuts_fd"},
	{findFD2NutsFD, "fd_2nd_nuts_fd"},
	{findFD3NutsFD, "fd_3d_nuts_fd"},
	// {findFD2nd3dNutsFD, "fd_2nd_3d_nuts_fd"},
	{findFD1nd3dNutsFD, "fd_1st_3d_nuts_fd"},
	{findFD4NutsFD, "fd_4_nuts_fd"},
	{findFD, "fd"},
	{findOESDBDFDNuts, "oesd_bdfd_nuts"},
	{findOESDNuts, "oesd_nuts"},
	{findGoodOESD, "good_oesd"},
	{findBadOESD, "bad_oesd"},
	{findOESD, "oesd"},
	{findGutShotBDFDNuts, "gutshot_bdfd_nuts"},
	{findGutShotNuts, "gutshot_nuts"},
	{findGutShotOvercard, "gutshot_overcard"},
	{findGoodGutShot, "good_gutshot"},
	{findBadGutShot, "bad_gutshot"},
	{findGutShot, "gutshot"},
	{findUnderPocket, "under_pocket"},
	{findOverCardsBDFDNuts, "overcards_bdfd_nuts"},
	{findAHigh, "ahigh"},
	{findTopCards, "top_cards"},
	{findOverCards, "overcards"},
	{findNomade, "nomade"},
}

func GetAllCombos() []Comb {
	combs := make([]Comb, len(combinations))
	for i, comb := range combinations {
		combs[i] = comb.name
	}
	return combs
}

func GetExtractors(combos []Comb) []ExtractorWithName {
	set := make(map[Comb]bool)
	for _, comb := range combos {
		set[comb] = true
	}
	var extractors []ExtractorWithName
	for _, extractor := range combinations {
		if set[extractor.name] {
			extractors = append(extractors, extractor)
		}
	}
	return extractors
}

func GetCombinations(board types.Board, hand types.Hand, useExtractors []ExtractorWithName) Comb {
	selector := newCombinationsSelector(board, hand)
	for _, combData := range useExtractors {
		if combData.extractor(&selector) {
			return combData.name
		}
	}
	return "nomade"
}
