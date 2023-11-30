package combinations

import (
	"fmt"
	"go-poker-tools/pkg/types"
	"testing"
)

func TestCombinations(t *testing.T) {
	table := []struct {
		board string
		hand  string
		comb  string
	}{
		{
			"As2sKs",
			"5s4s",
			"flush",
		},
		{
			"3s4c5d",
			"7d6s",
			"straight",
		},
		{
			"Kc7d2h",
			"KsKd",
			"top_set",
		},
		{
			"Kc7c2c",
			"7s7d",
			"medium_set",
		},
		{
			"Kc7c2c",
			"2s2d",
			"medium_set",
		},
		{
			"AcKs6d",
			"AsKd",
			"top_two_pairs",
		},
		{
			"AcKc6h",
			"As6c",
			"medium_two_pairs",
		},
		{
			"AcKc6h",
			"Ks6d",
			"medium_two_pairs",
		},
		{
			"QdJcTh",
			"KsKd",
			"overpair_oesd",
		},
		{
			"JdTc9h",
			"KsKd",
			"overpair_gsh",
		},
		{
			"2s3c8h",
			"QsQd",
			"high_overpair",
		},
		{
			"2s3c8h",
			"JsJc",
			"high_overpair",
		},
		{
			"6h2d4c",
			"9s9h",
			"low_overpair",
		},
		{
			"2s3c4d",
			"7c7h",
			"low_overpair",
		},
		{
			"9s6s2s",
			"QdQs",
			"high_overpair_fd_3nuts",
		},
		{
			"7s6s2s",
			"8d8s",
			"low_overpair_fd",
		},
		{
			"Qd5s3s",
			"AsQs",
			"tp_fd_nuts",
		},
		{
			"Qs5s3s",
			"AsQd",
			"tp_fd_nuts",
		},
		{
			"Qd5s3s",
			"KsQs",
			"tp_fd_2nuts",
		},
		{
			"Qd5d3d",
			"Qs7d",
			"tp_fd_4nuts",
		},
		{
			"KcJdTs",
			"KsQd",
			"tp_oesd",
		},
		{
			"KsQcJd",
			"AsKd",
			"tp_gsh",
		},
		{
			"Jd3c2h",
			"AsJh",
			"tp",
		},
		{
			"As9s3s",
			"JsJd",
			"pocket_tp_2_fd_3_nuts",
		},
		{
			"Ad8d2d",
			"QdQc",
			"pocket_tp_2_fd_2_nuts",
		},
		{
			"As2s3s",
			"6s6d",
			"pocket_tp_2_fd_4_nuts",
		},
		{
			"Js9d8c",
			"TsTd",
			"pocket_tp_2_oesd",
		},
		{
			"Js9d7c",
			"TsTd",
			"pocket_tp_2_gsh",
		},
		{
			"Js9d6c",
			"TsTd",
			"pocket_tp_2",
		},
		{
			"9s8s2s",
			"As8d",
			"2nd_fd_nuts",
		},
		{
			"Kd8dAd",
			"KsJd",
			"2nd_fd_2_nuts",
		},
		{
			"Ts7s2s",
			"8s7d",
			"2nd_fd_4_nuts",
		},
		{
			"9s7s6d",
			"8s7d",
			"2nd_oesd",
		},
		{
			"9s7s5d",
			"8s7d",
			"2nd_gsh",
		},
		{
			"9s7s2d",
			"8s7d",
			"2nd",
		},
		{
			"AsKs2s",
			"QsQd",
			"pocket_between_2_3_fd_nuts",
		},
		{
			"AsKs2s",
			"JsJd",
			"pocket_between_2_3_fd_2_nuts",
		},
		{
			"AsKs2s",
			"7s7d",
			"pocket_between_2_3_fd_4_nuts",
		},
		{
			"KsQcTd",
			"JsJd",
			"pocket_between_2_3_oesd",
		},
		{
			"KsQc9d",
			"JsJd",
			"pocket_between_2_3_gsh",
		},
		{
			"KsQc8d",
			"JsJd",
			"pocket_between_2_3",
		},
		{
			"KsQs3s",
			"As3d",
			"3d_hands_fd_nuts",
		},
		{
			"KsQs3s",
			"Js3d",
			"3d_hands_fd_2_nuts",
		},
		{
			"KsQs3s",
			"5s3d",
			"3d_hands_fd_4_nuts",
		},
		{
			"6s4c3h",
			"5s3d",
			"3d_hands_oesd",
		},
		{
			"3s4c6h",
			"7s3d",
			"3d_hands_gsh",
		},
		{
			"3s4d8h",
			"7s3d",
			"3d_hands",
		},
		{
			"AsKsQs",
			"JsJd",
			"fd_gsh_fd_1_card",
		},
		{
			"AsKdQc",
			"2s2d",
			"under_pocket",
		},
		{
			"KcTs2d",
			"Ac6d",
			"ahigh",
		},
		{
			"AsQdKc",
			"5s6d",
			"nomade",
		},
		{
			"3c2d2h",
			"JsTc",
			"top_cards",
		},
		{
			"7s3d2c",
			"Td9s",
			"overcards",
		},
		{
			"6s7sAd",
			"4s5s",
			"fd_oesd_fd_2_cards",
		},
		{
			"6s8sAd",
			"4s5s",
			"fd_gsh_fd_2_cards",
		},
		{
			"As7s2s",
			"KsQd",
			"fd_nuts_fd",
		},
		{
			"AsQs3d",
			"Js2s",
			"fd_2nd_nuts_fd",
		},
		{
			"As9s3d",
			"QsTs",
			"fd_2nd_nuts_fd",
		},
		{
			"6s8sTs",
			"5s4d",
			"fd_gsh_fd_1_card",
		},
		{
			"JsTc3h",
			"KsQd",
			"oesd_nuts",
		},
		{
			"4s5d6h",
			"As3d",
			"bad_oesd",
		},
		{
			"Tc9dQs",
			"8c7c",
			"good_oesd",
		},
		{
			"Js9d6c",
			"KsQd",
			"gutshot_nuts",
		},
		{
			"3d4c6s",
			"As2s",
			"gutshot_bdfd_nuts",
		},
		{
			"8c9hJd",
			"As7d",
			"bad_gutshot",
		},
		{
			"As2s3s",
			"5s4s",
			"straight_flush",
		},
		{
			"Ac2hKc",
			"2s2d",
			"set",
		},
		{
			"3c4d7h",
			"JsJd",
			"overpair",
		},
		{
			"QsTd2c",
			"KcJd",
			"oesd",
		},
		{
			"AcJc2d",
			"KsQd",
			"gutshot",
		},
		{
			"4c8c9cTd",
			"Ac3h",
			"fd",
		},
		{
			"5h5c5d",
			"5s3d",
			"quads",
		},
		{
			"3c2h3d",
			"2c2d",
			"full_house",
		},
		{
			"QhQc2d",
			"KcQd",
			"trips",
		},
		{
			"AcKc6h",
			"As6c",
			"two_pairs",
		},
		{
			"AcKc6h",
			"Ks6d",
			"two_pairs",
		},
		{
			"Ks3d8s7c",
			"AsKd",
			"tp_bdfd_nuts", // old_tp
		},
		{
			"8s3c2s6d",
			"8d7c",
			"old_tp",
		},
		{
			"3c8d9cKs",
			"AsKd",
			"new_tp",
		},
		{
			"6c9d8cQs",
			"KsQc",
			"new_tp",
		},
		{
			"As2s3s",
			"6s6d",
			"pocket_tp_2_fd",
		},
		{
			"Ts7s2s",
			"8s7d",
			"2nd_fd",
		},
		{
			"AsKs2s",
			"7s7d",
			"pocket_between_2_3_fd",
		},
		{
			"KsQs3s",
			"5s3d",
			"3d_hands_fd",
		},
		{
			"AsQs3d",
			"Ks2s",
			"fd_1st_3d_nuts_fd",
		},
		{
			"As9s3d",
			"KsTs",
			"fd_1st_3d_nuts_fd",
		},
		{
			"3s4s8s9s",
			"AsKd",
			"high_flush_j",
		},
		{
			"3s9sTs2s",
			"Ks8s",
			"high_flush_j",
		},
		{
			"3s4s8s9s",
			"TsKd",
			"low_flush",
		},
		{
			"3s9sTs2s",
			"4s8s",
			"low_flush",
		},
		{
			"7c9dThJc",
			"7s8d",
			"high_straight",
		},
		{
			"9cTdJsQd",
			"Ks2c",
			"high_straight",
		},
		{
			"6s7d8cJs",
			"4s5d",
			"high_straight",
		},
		{
			"2s3d4h5d",
			"AcAd",
			"low_straight",
		},
		{
			"4s5d6c7s",
			"3hTd",
			"low_straight",
		},
		{
			"9s6s2s",
			"QdQs",
			"high_overpair_fd_3nuts",
		},
		{
			"Ks2c2d",
			"As3s",
			"ahigh",
		},
		{
			"Ks2c2d",
			"AsQd",
			"ahigh",
		},
		{
			"Ks2c2d",
			"AcJd",
			"ahigh",
		},
		{
			"6c6d9c",
			"5s8d",
			"good_gutshot",
		},
		// TODO turn think
		{
			"4s5dTcJs",
			"8s7d",
			"oesd_nuts",
		},
		// TODO turn think
		{
			"8c9dThKh",
			"AsQh",
			"gutshot_overcard",
		},
		{
			"7c7d9c",
			"4s4d",
			"under_pocket",
		},
		{
			"9c3d2h",
			"KcAc",
			"overcards_bdfd_nuts",
		},
		{
			"As4d3h",
			"5d2c",
			"high_straight",
		},
		{
			"7c8cJh9dTd",
			"KhQh",
			"high_straight",
		},
		{
			"Ac2h3d4s5d",
			"6d7d",
			"high_straight",
		},
		{
			"Ac2h3d4s5d",
			"5c6d",
			"high_straight",
		},
		{
			"9cTdJsQhKd",
			"AcKh",
			"high_straight",
		},
		{
			"2h3h4h5h6h",
			"6d7h",
			"straight_flush",
		},
		// {
		// 	"2h3h4h6h7h",
		// 	"Ah3d",
		// 	"flush",
		// },
		{
			"2c2d2h2s3c",
			"AhAc",
			"high_overpair",
		},
		{
			"2cJs6s",
			"AhQh",
			"ahigh",
		},
		{
			"2dQsQd",
			"5h5s",
			"pocket_tp_2",
		},
		{
			"AsAc2d",
			"8s8d",
			"pocket_tp_2",
		},
		{
			"AsAc9d",
			"8s8d",
			"under_pocket",
		},
		{
			"AsKs3d",
			"QsQd",
			"pocket_between_2_3",
		},
		{
			"AsAc3d",
			"QsQd",
			"pocket_tp_2",
		},
		{
			"AhTcAc",
			"QdQh",
			"pocket_tp_2",
		},
		{
			"Ts3cTd",
			"6h6c",
			"pocket_tp_2",
		},
		// {
		// 	"3dAdKh",
		// 	"Qd5d",
		// 	"fd_2nd_3d_nuts_fd",
		// },
		{
			"8s9dJsQc",
			"Qh4h",
			"new_tp",
		},
		{
			"8s8d2s",
			"7s7d",
			"pocket_tp_2",
		},
		{
			"8s2s2d",
			"7s7d",
			"pocket_between_2_3",
		},
		{
			"6s6d8s",
			"7s7d",
			"pocket_between_2_3",
		},
		// {
		// 	"2d7hQs",
		// 	"AdQd",
		// 	"tp_bdfd_nuts",
		// },
		{
			"8h2h5s",
			"As8s",
			"tp_bdfd_nuts",
		},
		{
			"Th6s5s",
			"Ah6h",
			"2nd_bdfd_nuts",
		},
		{
			"6d3sQc",
			"Ac6c",
			"2nd_bdfd_nuts",
		},
		{
			"KdTh3s",
			"Ah3h",
			"3nd_bdfd_nuts",
		},
		{
			"TdJc5c",
			"Ad5d",
			"3nd_bdfd_nuts",
		},
		{
			"Td8h7s",
			"Ad9d",
			"oesd_bdfd_nuts",
		},
		{
			"8s9c7c",
			"AsTs",
			"oesd_bdfd_nuts",
		},
		{
			"Jh7h8s",
			"AsTs",
			"gutshot_bdfd_nuts",
		},
		{
			"Jd3s2s",
			"Ad5d",
			"gutshot_bdfd_nuts",
		},
		{
			"5d5h4c",
			"AcKc",
			"overcards_bdfd_nuts",
		},
		{
			"2d5h8s",
			"AdTd",
			"overcards_bdfd_nuts",
		},
		{"9s6s2s", "AdAs", "high_overpair_fd_nuts"},
		{"Ts5s2s", "KcKs", "high_overpair_fd_2nuts"},
		{"Js6s5s", "QdQs", "high_overpair_fd_3nuts"},
		{"Ts5s2s", "JdJs", "high_overpair_fd_4nuts"},
		{"Qd5s3s", "AsQs", "tp_fd_nuts"},
		{"Qh7d6d", "KdQd", "tp_fd_2nuts"},
		{"As8s4s", "AcJs", "tp_fd_3nuts"},
		{"AsKs2s", "Ac9s", "tp_fd_4nuts"},
		{"As9s3s", "KsKd", "pocket_tp_2_fd_nuts"},
		{"Ks7s2s", "QsQd", "pocket_tp_2_fd_2_nuts"},
		{"Qs7s2s", "JsJd", "pocket_tp_2_fd_3_nuts"},
		{"Qs6s2s", "As6c", "2nd_fd_nuts"},
		{"Js5s2s", "Ks5c", "2nd_fd_2_nuts"},
		{"Js5s2s", "Qs5d", "2nd_fd_3_nuts"},
		{"Ks8s7s", "Ts9d", "fd_oesd_fd_1_card"},
		{"Ks7s6s", "Ts9d", "fd_gsh_fd_1_card"},
		{"KsQs3s", "JsJc", "pocket_between_2_3_fd_2_nuts"},
		{"KsJs3s", "TsTc", "pocket_between_2_3_fd_3_nuts"},
		{"KsQs3s", "Js3d", "3d_hands_fd_2_nuts"},
		{"QcJc6c", "Tc6d", "3d_hands_fd_3_nuts"},
		{"AsQs3d", "Ks2s", "fd_nuts_fd"},
		{"AcQc5d", "Jc7c", "fd_2nd_nuts_fd"},
		{"2c9dTh", "JcQd", "oesd_nuts"},
		{"9s8c2c", "QdJh", "gutshot_nuts"},
		{"3c6s8h", "Ts7d", "gutshot_overcard"},
		{"4dTc7s", "Jh9h", "gutshot_overcard"},
		{"AcJs6c", "AsTs", "tp_bdfd_nuts"},
	}
	skipCombos := map[int][]Comb{
		66:  {"top_set", "medium_set"},
		67:  {"high_overpair"},
		68:  {"good_oesd"},
		69:  {"good_gutshot"},
		70:  {"fd_nuts_fd", "fd_1st_3d_nuts_fd"},
		74:  {"medium_two_pairs"},
		75:  {"medium_two_pairs"},
		80:  {"pocket_tp_2_fd_4_nuts"},
		81:  {"2nd_fd_4_nuts"},
		82:  {"pocket_between_2_3_fd_4_nuts"},
		83:  {"3d_hands_fd_4_nuts"},
		84:  {"fd_nuts_fd"},
		85:  {"fd_nuts_fd"},
		0:   {"low_flush"},
		1:   {"high_straight", "low_straight"},
		95:  {"high_overpair_fd"},
		100: {"good_oesd"},
		110: {"high_flush_j"},
	}
	skipCombosAllTests := []Comb{
		"high_overpair_fd",
		"tp_fd_nuts_fd",
		"pocket_tp_2_fd_1_3_nuts",
		"2nd_fd_1_3_nuts",
		"pocket_between_2_3_fd_2_3_nuts",
		"3d_hands_fd_2_3_nuts",
		"fd_2nd_3d_nuts_fd",
	}
	allCombos := GetAllCombos()
	for i, testCase := range table {
		hand := types.ParseHand(testCase.hand)
		board := types.ParseBoard(testCase.board)
		trueComb := testCase.comb
		combos := subCombos(allCombos, skipCombos[i])
		combos = subCombos(combos, skipCombosAllTests)
		extractors := GetExtractors(combos)
		comb := GetCombinations(board, hand, extractors)
		if comb != Comb(trueComb) {
			t.Error(fmt.Sprintf("Test %d, board: %s, hand: %s, %s != %s", i, testCase.board, testCase.hand, comb, trueComb))
			return
		}
	}

}

func subCombos(source []Comb, skip []Comb) []Comb {
	var result []Comb
	for _, comb := range source {
		found := false
		for _, c := range skip {
			if c == comb {
				found = true
				break
			}
		}
		if !found {
			result = append(result, comb)
		}
	}
	return result
}
