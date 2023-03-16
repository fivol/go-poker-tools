package combinations

import (
	"go-poker-tools/pkg/generics"
	"go-poker-tools/pkg/types"
)

func ifThenElse(condition bool, then, other uint8) uint8 {
	if condition {
		return then
	}
	return other
}

type Source int

const (
	Total Source = iota
	Board
)

func (s *Selector) getSource(source Source) cardsInfo {
	if source == Total {
		return s.total
	}
	if source == Board {
		return s.board
	}
	return s.hand
}

func (s *Selector) handSameValues() bool {
	return s.hand.maxSameValues == 2
}
func (s *Selector) handSameSuit() bool {
	return s.firstCard().Suit() == s.secondCard().Suit()
}
func (s *Selector) isPokerPair() bool {
	return s.handSameValues() && s.board.values[s.poketPairValue()] == 0
}
func (s *Selector) boardCardsCount() uint8 {
	return uint8(len(s.board.cards))
}
func (s *Selector) hasBoardPairs() bool {
	return s.board.maxSameValues == 1
}
func (s *Selector) otherCard(card types.Card) types.Card {
	if s.hand.cards[0] == card {
		return s.secondCard()
	}
	return s.firstCard()
}
func (ci *cardsInfo) getFullHouse() (bool, uint8, uint8) {
	if ci.maxSameValues != 3 {
		return false, 0, 0
	}
	foundPair := false
	setIdx := uint8(0)
	pairIdx := uint8(0)
	for i := ci.minValue; i <= ci.maxValue; i++ {
		if ci.values[i] >= 2 {
			if ci.values[i] == 3 {
				setIdx = generics.Max(i, setIdx)
			}
			if i != setIdx {
				pairIdx = i
			}
			if foundPair {
				return true, setIdx, pairIdx
			}
			foundPair = true
		}
	}
	return false, 0, 0
}
func (s *Selector) badOESDCard(card types.Card) bool {
	if card.Value() >= s.board.minValue {
		return false
	}
	return s.twoWaySD(card) && !s.twoWaySD(s.otherCard(card))
}
func (s *Selector) badGutShotCard(card types.Card) bool {
	if card.Value() >= s.board.minValue && card.Value() != 12 {
		return false
	}
	up := s.board.upStairLen(card.Value()+1) + 1
	up2 := s.board.upStairLen(card.Value() + up + 1)
	return up == 4 || up+up2 >= 4
}
func (s *Selector) pocketPairLessBoardCount(count uint8) bool {
	return s.isPokerPair() && s.board.graterValuesCount[s.poketPairValue()] == count
}
func (s *Selector) getFDIdx(fd FD) int {
	for i, topFD := range s.topFDs {
		if topFD == fd {
			return i + 1
		}
	}
	return 4
}
func (s *Selector) anySD() bool {
	return s.handOneCardSD() || s.handTwoWaySD()
}
func (s *Selector) noCombos() bool {
	return !s.FD() && !s.anySD() && !s.handSameValues()
}
func (s *Selector) isFDBetween(from, to int) bool {
	found, fd := s.getFD()
	if !found {
		return false
	}
	idx := s.getFDIdx(fd)
	return idx >= from && idx <= to
}
func (s *Selector) hasFD() bool {
	found, _ := s.getFD()
	return found
}
func (s *Selector) isTopFD() bool {
	return s.isFDBetween(1, 1)
}
func (s *Selector) inStraightMiddle(card types.Card) bool {
	if card.Value() == 0 {
		return false
	}
	if s.total.downStairLen(card.Value()-1) == 0 {
		return false
	}
	return s.total.upStairLen(card.Value()+1)+s.total.downStairLen(card.Value()-1)+1 >= 5
}
func (s *Selector) isWeakFD() bool {
	return s.isFDBetween(4, 10)
}
func (s *Selector) firstCard() types.Card {
	return s.hand.cards[0]
}
func (s *Selector) handMinValue() uint8 {
	return s.hand.minValue
}
func (s *Selector) card(idx uint8) types.Card {
	return s.hand.cards[idx]
}
func (s *Selector) secondCard() types.Card {
	return s.hand.cards[1]
}
func (s *Selector) poketPairValue() uint8 {
	return s.hand.cards[0].Value()
}
func (s *Selector) twoWaySD(card types.Card) bool {
	up := s.total.upStairLen(card.Value())
	down := s.total.downStairLen(card.Value())
	size := up + down - 1
	if s.total.isOpenUpStair(card.Value()) && s.total.isOpenDownStair(card.Value()) && size == 4 {
		return true
	}
	if size+s.total.upStairLen(card.Value()+up+1) >= 4 &&
		s.total.upStairLen(card.Value()+up+1) > 0 &&
		s.total.downStairLen(card.Value()-down-1) > 0 &&
		size+s.total.downStairLen(card.Value()-down-1) >= 4 {
		return true
	}
	return false
}
func (s *Selector) handTwoWaySD() bool {
	return s.twoWaySD(s.firstCard()) || s.twoWaySD(s.secondCard())
}
func (s *Selector) gutShot(card types.Card) bool {
	return (s.SD(card) || s.gutShotWhole(card)) && !s.twoWaySD(card)
}
func (s *Selector) handOneCardSD() bool {
	return s.gutShot(s.firstCard()) || s.gutShot(s.secondCard())
}
func (s *Selector) SD(card types.Card) bool {
	up := s.total.upStairLen(card.Value())
	down := s.total.downStairLen(card.Value())
	if card.Value() == 12 {
		if up > down {
			down = 1
		} else {
			up = 1
		}
	}
	return up+down-1 == 4
}
func (s *Selector) FD() bool {
	return s.maxSuitsWithHand() == 4
}
func (s *Selector) getFDSuit(suit uint8) (bool, FD) {
	if s.board.suits[suit] == 4 {
		return false, 0
	}
	if s.total.suits[suit] == 4 {
		return true, FD(s.hand.maxValueSuits[suit])
	}
	return false, 0
}
func (s *Selector) getFD() (bool, FD) {
	found1, fd1 := s.getFDSuit(s.firstCard().Suit())
	found2, fd2 := s.getFDSuit(s.secondCard().Suit())
	if !found1 {
		return found2, fd2
	}
	if !found2 {
		return found1, fd1
	}
	if fd1 < fd2 {
		return true, fd2
	}
	return true, fd1
}
func (s *Selector) upGutshotWhole(card types.Card) bool {
	down := s.total.downStairLen(card.Value())
	up := s.total.upStairLen(card.Value())
	size := up + down - 1
	if size >= 4 {
		return false
	}
	return s.total.upStairLen(card.Value()+up+1)+size == 4
}
func (ci *cardsInfo) getStraight() (bool, uint8) {
	if ci.maxOrderLen < 5 {
		return false, 0
	}
	return true, ci.maxOrderIdx + ci.maxOrderLen - 5
}
func (s *Selector) downGutshotWhole(card types.Card) bool {
	down := s.total.downStairLen(card.Value())
	up := s.total.upStairLen(card.Value())
	size := up + down - 1
	if size >= 4 {
		return false
	}
	return s.total.downStairLen(card.Value()-down-1)+size >= 4
}
func (s *Selector) gutShotWhole(card types.Card) bool {
	return s.upGutshotWhole(card) || s.downGutshotWhole(card)
}
func (s *Selector) HasFlush(card types.Card) bool {
	if s.total.suits[card.Suit()] < 5 {
		return false
	}
	if s.board.suits[card.Suit()] < 5 {
		return true
	}
	if s.board.graterValuesCount[card.Value()] < 5 {
		return true
	}
	return false
}
func (s *Selector) pairHandBoardIdx(card types.Card) uint8 {
	return s.board.valueOrder[card.Value()]
}
func (s *Selector) pairWithBoardIdx() uint8 {
	if s.handSameValues() {
		return 0
	}
	first := s.pairHandBoardIdx(s.firstCard())
	second := s.pairHandBoardIdx(s.secondCard())
	if first > 0 && second > 0 {
		return generics.Min(first, second)
	}
	if first > 0 {
		return first
	}
	if second > 0 {
		return second
	}
	return 0
}
func (s *Selector) handPairTopBoard() bool {
	return s.pairWithBoardIdx() == 1
}
func (s *Selector) hasTurn() bool {
	return len(s.board.cards) >= 4
}
func (s *Selector) turnCard() types.Card {
	return s.board.cards[3]
}
func (s *Selector) isOneCardPairWithBoard(idx uint8) bool {
	return s.pairWithBoardIdx() == idx
}
func (s *Selector) maxSuitsWithHand() uint8 {
	return generics.Max(s.total.suits[s.hand.cards[0].Suit()], s.total.suits[s.hand.cards[1].Suit()])
}
func (s *Selector) isPokerPairGraterBoard() bool {
	return s.pocketPairLessBoardCount(0)
}
func (ci *cardsInfo) upStairLen(value uint8) uint8 {
	if value > 12 {
		return 0
	}
	if value < 0 {
		return 0
	}
	return ci.stairsUpLen[value]
}
func (ci *cardsInfo) isOpenUpStair(value uint8) bool {
	size := ci.upStairLen(value)
	return value+size < 12+1
}
func (ci *cardsInfo) isOpenDownStair(value uint8) bool {
	size := ci.downStairLen(value)
	if size == 0 {
		return false
	}
	return value+1 > size
}
func (ci *cardsInfo) downStairLen(value uint8) uint8 {
	if value > 12 {
		return 0
	}
	if value < 0 {
		return 0
	}
	return ci.stairsDownLen[value]
}

func findFlush(s *Selector) bool {
	/*
		flush
		5 карт одной масти
	*/
	return s.HasFlush(s.firstCard()) || s.HasFlush(s.secondCard())
}

func findHighFlushJ(s *Selector) bool {
	/*
		high_flush_j
		Флеш, когда на доске лежит 4 карты одной масти, и старшая карта у нас из руки, которая используется, выше или равна J
	*/
	return s.firstCard().Value() >= 9 && s.HasFlush(s.firstCard()) ||
		s.secondCard().Value() >= 9 && s.HasFlush(s.secondCard())
}

func findLowFlush(s *Selector) bool {
	/*
		low_flush
		Флеш, когда на доске лежит 4 карты одной масти, и старшая карта у нас из руки, которая используется,ниже J
	*/
	return findFlush(s) && !findHighFlushJ(s)
}

func findStraightFlush(s *Selector) bool {
	/*
		straight_flush
		5 карт одной масти подряд
	*/
	if !findStraight(s) {
		return false
	}
	if !findFlush(s) {
		return false
	}
	var chart [4][13]bool
	for _, card := range s.total.cards {
		chart[card.Suit()][card.Value()] = true
	}
	for suit, count := range s.total.suits {
		if count >= 5 {
			for value, up := range s.total.stairsUpLen {
				if up >= 5 && s.board.stairsUpLen[value] < 5 {
					match := true
					for i := value; i < value+5; i++ {
						val := i % 13
						if !chart[suit][val] {
							match = false
							break
						}
					}
					if match {
						return true
					}
				}
			}
		}
	}
	return false
}

func findStraight(s *Selector) bool {
	/*
		straight
		5 карт подряд
	*/
	found, totalStraight := s.total.getStraight()
	if !found {
		return false
	}
	boardFound, boardStraight := s.board.getStraight()
	if !boardFound {
		return true
	}
	return totalStraight > boardStraight
}

func findHighStraight(s *Selector) bool {
	/*
		high_straight
		Стрит с использованием 4 карт борда и 1 карты с руки и эта карта не
		является первой (младшей, за исключением случая, когда стрит A2345) картой в стрите
	*/
	if !findStraight(s) {
		return false
	}
	return s.inStraightMiddle(s.firstCard()) || s.inStraightMiddle(s.secondCard())
}

func findLowStraight(s *Selector) bool {
	/*
		low_straight
		Стрит с использованием 4 карт борда и 1 карты с руки и эта карта является первой
		(младшей, за исключением случая, когда стрит A2345) картой в стрите
	*/
	return findStraight(s) && !findHighStraight(s)
}

func findSet(s *Selector) bool {
	/*
		set
		У нас карманная пара и есть совпадение с одной из карт борда
	*/
	return s.handSameValues() && s.total.values[s.firstCard().Value()] == 3
}

func findTrips(s *Selector) bool {
	/*
		trips
		Одна наша карта совпадает сразу с двумя на борде по номиналу
	*/
	return (s.total.values[s.firstCard().Value()] == 3 || s.total.values[s.secondCard().Value()] == 3) &&
		!findSet(s)
}

func findQuads(s *Selector) bool {
	/*
		quads
		4 одинаковых карты, для составления
		которых используется хотя бы одна карта из нашей руки
	*/
	return s.total.values[s.firstCard().Value()] == 4 ||
		s.total.values[s.secondCard().Value()] == 4
}

func findFullHouse(s *Selector) bool {
	/*
		full_house
		Три одинаковых карты плюс пара
	*/
	found, setIdxT, pairIdxT := s.total.getFullHouse()
	if !found {
		return false
	}
	foundBoard, setIdxB, pairIdxB := s.board.getFullHouse()
	if !foundBoard {
		return true
	}
	return (setIdxT > setIdxB) || (setIdxT == setIdxB && pairIdxT > pairIdxB)
}

func findTopSet(s *Selector) bool {
	// top_set
	return s.handSameValues() && s.board.valueOrder[s.hand.cards[0].Value()] == 1
}

func findMediumSet(s *Selector) bool {
	// medium_set
	return s.handSameValues() && s.board.valueOrder[s.hand.cards[0].Value()] > 1
}

func findTopTwoPairs(s *Selector) bool {
	// top_two_pairs
	firstCardBoardOrder := s.board.valueOrder[s.hand.cards[0].Value()]
	secondCardBoardOrder := s.board.valueOrder[s.hand.cards[1].Value()]
	return !s.handSameValues() &&
		firstCardBoardOrder != 0 &&
		secondCardBoardOrder != 0 &&
		firstCardBoardOrder <= 2 &&
		secondCardBoardOrder <= 2
}

func findMediumTwoPairs(s *Selector) bool {
	/*
		medium_two_pairs
		Две пары, которые образуются из совпадения обеих наших карт с двумя картами борда (c высшей и низшей или со средней и низшей)
	*/
	firstCardBoardOrder := s.board.valueOrder[s.hand.cards[0].Value()]
	secondCardBoardOrder := s.board.valueOrder[s.hand.cards[1].Value()]
	return !s.handSameValues() &&
		firstCardBoardOrder != 0 &&
		secondCardBoardOrder != 0 &&
		(firstCardBoardOrder == uint8(len(s.board.cards)) || secondCardBoardOrder == uint8(len(s.board.cards)))
}

func findTwoPairs(s *Selector) bool {
	/*
		two_pairs
		Две пары, которые образуются из совпадения обеих наших карт с двумя картами борда
	*/
	return !s.handSameValues() &&
		s.board.values[s.firstCard().Value()] == 1 &&
		s.board.values[s.secondCard().Value()] == 1
}

func findOverPairOESD(s *Selector) bool {
	/*
		overpair_oesd
		Карманная пара выше всех карт флопа и образующая двухстороннее стритдро
	*/
	return s.isPokerPairGraterBoard() && s.twoWaySD(s.firstCard())
}

func findOverPair(s *Selector) bool {
	/*
		overpair
		Карманная пара, выше всех карт борда
	*/
	return s.isPokerPairGraterBoard()
}

func findOverPairGSH(s *Selector) bool {
	// overpair_gsh
	return s.isPokerPairGraterBoard() && s.gutShot(s.firstCard())
}

func findHighOverPair(s *Selector) bool {
	// high_overpair
	return s.isPokerPairGraterBoard() && s.poketPairValue() >= 8
}

func findLowOverPair(s *Selector) bool {
	// low_overpair
	return s.isPokerPairGraterBoard() && s.poketPairValue() <= 7
}

func findHighOverPairFD(s *Selector) bool {
	// high_overpair_fd
	return findHighOverPair(s) && s.maxSuitsWithHand() == 4
}

func findLowOverPairFD(s *Selector) bool {
	/*
		low_overpair_fd
		Оверпара от 99 включительно и ниже,имеющая флешдро
	*/
	return findLowOverPair(s) && s.maxSuitsWithHand() == 4
}

func findOverPairFD(s *Selector) bool {
	/*
		overpair_fd
		Оверпара, имеющая флешдро
	*/
	return findOverPair(s) && s.hasFD()
}

func findTPFDNutsFd(s *Selector) bool {
	// tp_fd_nuts_fd
	return s.handPairTopBoard() && s.isTopFD()
}

func findTPFD(s *Selector) bool {
	// tp_fd
	return s.handPairTopBoard() && s.FD()
}

func findTPOESD(s *Selector) bool {
	// tp_oesd
	var otherIdx = 0
	if s.board.valueOrder[s.hand.cards[0].Value()] != 1 {
		otherIdx = 1
	}
	return findTP(s) && s.twoWaySD(s.hand.cards[otherIdx])
}

func findOESD(s *Selector) bool {
	/*
		oesd
		Двухстороннее Стритдро, для составления стрита которому необходима одна карта
		и используется хотя бы одна карта, которая у нас на руках
	*/
	return s.twoWaySD(s.firstCard()) || s.twoWaySD(s.secondCard())
}

func findTPGSH(s *Selector) bool {
	// tp_gsh
	var otherIdx = 0
	if s.board.valueOrder[s.hand.cards[0].Value()] != 1 {
		otherIdx = 1
	}
	return findTP(s) && s.gutShot(s.hand.cards[otherIdx])
}

func findTP(s *Selector) bool {
	/*
		tp
		Совпадение одной из наших карманных карт с высшей картой стола
	*/
	return s.handPairTopBoard()
}

func findOldTP(s *Selector) bool {
	/*
		old_tp
		Совпадение одной из наших карманных карт с высшей картой стола,
		образовавшаяся на флопе и не изменившаяся к терну
		(на терне вышла карта ниже самой высокой карты флопа и
		у нас осталось совпадение с самой высокой картой доски)
	*/
	return s.hasTurn() &&
		s.handPairTopBoard() &&
		s.hand.values[s.turnCard().Value()] == 0 &&
		s.board.valueOrder[s.turnCard().Value()] > 1
}

func findNewTP(s *Selector) bool {
	/*
			new_tp
			Совпадение одной из наших карманных карт с высшей картой стола,
		 	образовавшаяся на терне
			(на терне вышла карта выше всех карт флопа и у нас совпадение с этой картой)
	*/
	return s.hasTurn() &&
		s.handPairTopBoard() &&
		!findOldTP(s)
}

func findPocketTP2FD13Nuts(s *Selector) bool {
	/*
		pocket_tp_2_fd_1_3_nuts
		Карманная пара ниже одной карты борда и  имеющая от первого до третьего по силе флешдро
	*/
	return findPocketTop2(s) && s.isFDBetween(1, 3)
}

func findPocketTP2FD4Nuts(s *Selector) bool {
	/*
		pocket_tp_2_fd_4_nuts
		Карманная пара ниже одной карты борда и  имеющая 4е или ниже по силе флешдро
	*/
	return findPocketTop2(s) && s.isWeakFD()
}

func findPocketTP2FD(s *Selector) bool {
	/*
		pocket_tp_2_fd
		Карманная пара ниже одной карты борда и имеющая флешдро
	*/
	return findPocketTop2(s) && s.hasFD()
}

func findPocketTP2OESD(s *Selector) bool {
	/*
		pocket_tp_2_oesd
		Карманная пара ниже одной карты борда и имеющая двустороннее стритдро
	*/
	return findPocketTop2(s) && s.twoWaySD(s.firstCard())
}

func findPocketTP2GSH(s *Selector) bool {
	/*
		pocket_tp_2_gsh
		Карманная пара ниже одной карты борда и имеющая  стритдро на одну карту
	*/
	return findPocketTop2(s) && s.gutShot(s.firstCard())
}

func findPocketTop2(s *Selector) bool {
	/*
		pocket_tp_2
		Карманная пара ниже одной карты борда
	*/
	return s.pocketPairLessBoardCount(1)
}

func findSecondFD13Nuts(s *Selector) bool {
	/*
		2nd_fd_1_3_nuts
		Совпадение одной из наших карманных карт со второй по номиналу картой борда плюс 1-3 флешдро по силе
		(на основе номинала карты, образующей флешдро)
	*/
	return findSecond(s) && s.isFDBetween(1, 3)
}

func findSecondFD4Nuts(s *Selector) bool {
	/*
		2nd_fd_4_nuts
		Совпадение одной из наших карманных карт со второй по номиналу картой борда плюс флешдро, которое не входит в топ-3 по номиналу карты, образующее это флешдро
		(на основе номинала карты, образующей флешдро)
	*/
	return findSecond(s) && s.isWeakFD()
}

func findSecondFD(s *Selector) bool {
	/*
			2nd_fd
			Совпадение одной из наших карманных карт со второй по номиналу картой борда плюс флешдро
		 	(на основе номинала карты, образующей флешдро)
	*/
	return findSecond(s) && s.hasFD()
}

func findSecondOESD(s *Selector) bool {
	/*
		2nd_oesd
		Совпадение одной из наших карманных карт со второй по номиналу картой борда
		плюс двустороннее стритдро, образованное нашей второй картой
	*/
	first := s.pairHandBoardIdx(s.firstCard())
	return findSecond(s) && s.twoWaySD(s.card(ifThenElse(first == 2, 1, 0)))
}

func findSecondGSH(s *Selector) bool {
	/*
		2nd_gsh
		Совпадение одной из наших карманных карт со второй по
		номиналу картой борда плюс  стритдро на одну карту, образованное нашей второй картой
	*/
	first := s.pairHandBoardIdx(s.firstCard())
	return findSecond(s) && s.gutShot(s.card(ifThenElse(first == 2, 1, 0)))
}

func findSecond(s *Selector) bool {
	/*
		2nd
		Совпадение одной из наших карманных карт со второй по номиналу картой борда
	*/
	return s.isOneCardPairWithBoard(2)
}

func findPocketBetween23FDNuts(s *Selector) bool {
	/*
		pocket_between_2_3_fd_nuts
		Карманная пара ниже двух карт борда и  имеющая 1е по силе флешдро
	*/
	return findPocketBetween23(s) && s.isTopFD()
}

func findPocketBetween23FD23Nuts(s *Selector) bool {
	/*
		pocket_between_2_3_fd_2_3_nuts
		Карманная пара ниже двух карт борда и  имеющая 2-3е  по силе флешдро
	*/
	return findPocketBetween23(s) && s.isFDBetween(2, 3)
}

func findPocketBetween23FD4Nuts(s *Selector) bool {
	/*
		pocket_between_2_3_fd_4_nuts
		Карманная пара ниже двух карт борда и  имеющая 4е или ниже по силе флешдро
	*/
	return findPocketBetween23(s) && s.isWeakFD()
}

func findPocketBetween23FD(s *Selector) bool {
	/*
		pocket_between_2_3_fd
		Карманная пара ниже двух карт борда и имеющая флешдро
	*/
	return findPocketBetween23(s) && s.hasFD()
}

func findPocketBetween23OESD(s *Selector) bool {
	/*
		pocket_between_2_3_oesd
		Карманная пара ниже двух карт борда и имеющая двустороннее стритдро
	*/
	return findPocketBetween23(s) && s.twoWaySD(s.firstCard())
}

func findPocketBetween23GSH(s *Selector) bool {
	/*
		pocket_between_2_3_gsh
		Карманная пара ниже двух карт борда и имеющая  стритдро на одну карту
	*/
	return findPocketBetween23(s) && s.gutShot(s.firstCard())
}

func findPocketBetween23(s *Selector) bool {
	/*
		pocket_between_2_3
		Карманная пара ниже двух карт борда
	*/
	return s.pocketPairLessBoardCount(2)
}

func find3dHands(s *Selector) bool {
	/*
		3d_hands
		Пара с третьей картой борда по номиналу
	*/
	return s.pairWithBoardIdx() == 3
}

func find3dHandsFDNuts(s *Selector) bool {
	/*
		3d_hands_fd_nuts
		Совпадение одной из наших карманных карт с третьей по номиналу
		картой борда плюс флешдро, которое входит
		в топ-1 по номиналу карты, образующее это флешдро
		(на основе номинала карты, образующей флешдро)
	*/
	return find3dHands(s) && s.isTopFD()
}

func find3dHandsFD23Nuts(s *Selector) bool {
	/*
		3d_hands_fd_2_3_nuts
		Совпадение одной из наших карманных карт с
		третьей по номиналу картой борда плюс флешдро,
		которое входит в топ-2-3 по номиналу карты, образующее это флешдро
		(на основе номинала карты, образующей флешдро)
	*/
	return find3dHands(s) && s.isFDBetween(2, 3)
}

func find3dHandsFD4Nuts(s *Selector) bool {
	/*
		3d_hands_fd_4_nuts
		Совпадение одной из наших карманных карт с третьей по
		номиналу картой борда плюс флешдро, которое не входит в топ-3 по номиналу карты,
		образующее это флешдро
		(на основе номинала карты, образующей флешдро)
	*/
	return find3dHands(s) && s.isWeakFD()
}

func find3dHandsFD(s *Selector) bool {
	/*
			3d_hands_fd
			Совпадение одной из наших карманных карт с третьей по номиналу картой борда плюс флешдро
		 	(на основе номинала карты, образующей флешдро)
	*/
	return find3dHands(s) && s.hasFD()
}

func find3dHandsOESD(s *Selector) bool {
	/*
		3d_hands_oesd
		Совпадение одной из наших карманных карт с третьей
		по номиналу картой борда плюс двустороннее стритдро,
		образованное нашей второй картой
	*/
	return find3dHands(s) && s.handTwoWaySD()
}

func find3dHandsGSH(s *Selector) bool {
	/*
		3d_hands_gsh
		Пара с третьей картой борда по номиналу плюс стритдро на одну карту
	*/
	return find3dHands(s) && s.handOneCardSD()
}

func findUnderPocket(s *Selector) bool {
	/*
		under_pocket
		Карманная пара ниже всех карт борда
	*/
	return s.pocketPairLessBoardCount(s.boardCardsCount())
}

func findUnderPocketFD12Nuts(s *Selector) bool {
	/*
		under_pocket_fd_1_2_nuts
		Карманная пара ниже всех карт борда, имеющая первое или второе по силе флешдро
	*/
	return findUnderPocket(s) && s.isFDBetween(1, 2)
}

func findAHigh(s *Selector) bool {
	/*
		ahigh
		Рука, которая никак не зацепилась за доску, и на руках старшая карта туз
	*/
	return s.hand.values[12] == 1 && s.noCombos()
}

func findNomade(s *Selector) bool {
	/*
		nomade
		Совсем пустые руки: без комбинаций,
		которые еще и не подходят в категории overcards и topcards
	*/
	return s.handMinValue() < 8 && s.board.maxValue > s.handMinValue() && s.noCombos()
}

func findTopCards(s *Selector) bool {
	/*
		top_cards
		Карманные карты, состоящие из двух картинок и не собравшие никакую из комбинаций сильнее
	*/
	return s.handMinValue() >= 8 && s.noCombos()
}

func findOverCards(s *Selector) bool {
	/*
		overcards
		Карманные карты, обе из которых выше карты борда
	*/
	return s.board.maxValue < s.handMinValue() && s.noCombos()
}

func findFdOESD2Cards(s *Selector) bool {
	/*
		fd_oesd_fd_2_cards
		Флешдро, для которого используются обе наших карты на руках+двустороннее стритдро
	*/
	return s.handSameSuit() && s.FD() && s.handTwoWaySD()
}

func findGSHFD2Cards(s *Selector) bool {
	/*
		fd_gsh_fd_2_cards
		Флешдро, для которого используются обе наших карты на руках+стритдро на 1 карту
	*/
	return s.handSameSuit() && s.FD() && s.handOneCardSD()
}

func findFDNutsFD(s *Selector) bool {
	/*
		fd_nuts_fd
		Сильнейшее флешдро, которое может быть на доске
		(сила определяется по старшей карте на руке, т.е.
		As4s и As2s на доске ssd будут считаться одинаково сильнейшими FD
	*/
	return s.isTopFD()
}

func findFD(s *Selector) bool {
	/*
		fd
		Для составления комбинации флеш не хватает одной
		карты и используется хотя бы одна карта, которая у нас на руках
	*/
	return s.FD()
}

func findFD2nd3dNutsFD(s *Selector) bool {
	/*
		fd_2nd_3d_nuts_fd
		Второе и третье флешдро по силе флешдро,
		которое может быть на доске (сила определяется по старшей карте на руке
	*/
	return s.isFDBetween(2, 3)
}

func findFD1nd3dNutsFD(s *Selector) bool {
	/*
		fd_1st_3d_nuts_fd
		От первого до третьего флешдро по силе флешдро, которое может быть на доске (сила определяется по старшей карте на руке
	*/
	return s.isFDBetween(1, 3)
}

func findFD4NutsFD(s *Selector) bool {
	/*
		fd_4_nuts_fd
		Четвертое и ниже по силе первой карты флешдро,
		которое может быть на доске
		(все остальные FD, не попавшие в nuts FD и 2,3 FD)
	*/
	return s.isWeakFD()
}

func findBadOESD(s *Selector) bool {
	/*
		bad_oesd
		Двухстороннее стритдро, которое для собирания хотя бы
		одного стрита будет использовать только одну карту,
		которая ниже всех карт борда
	*/
	return !s.handSameValues() && (s.badOESDCard(s.firstCard()) || s.badOESDCard(s.secondCard()))
}

func findGoodOESD(s *Selector) bool {
	/*
		good_oesd
		Все остальные OESD, которые не вошли в категорию bad OESD
	*/
	return findOESD(s) && !findBadOESD(s)
}

func findGoodGutShot(s *Selector) bool {
	/*
		good_gutshot
		Все остальные гатшоты, которые не вошли в категорию bad gutShotWhole
	*/
	return s.handOneCardSD() && !findBadGutShot(s)
}

func findGutShot(s *Selector) bool {
	/*
		gutshot
		Стритдро, для составления стрита которому необходима одна
		карта и используется хотя бы одна карта, которая у нас на руках
	*/
	return s.handOneCardSD()
}

func findBadGutShot(s *Selector) bool {
	/*
		bad_gutshot
		Гатшот на одну карту, которая ниже всех карт борда
	*/
	return s.badGutShotCard(s.firstCard()) || s.badGutShotCard(s.secondCard())
}
