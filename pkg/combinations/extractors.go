package combinations

func findFlush(s *Selector) bool {
	return s.total.maxSameSuitsCount >= 5
}

func findStraight(s *Selector) bool {
	return s.total.maxOrderLen >= 5
}

func findSet(s *Selector) bool {
	return s.total.maxSameValues == 3
}

func findTopSet(s *Selector) bool {
	return s.hand.maxSameValues == 2 && s.board.valueOrder[s.hand.cards[0].Value()] == 1
}

func findMinimumSet(s *Selector) bool {
	return s.hand.maxSameValues == 2 && s.board.valueOrder[s.hand.cards[0].Value()] > 1
}

func findTwoPairs(s *Selector) bool {
	firstCardBoardOrder := s.board.valueOrder[s.hand.cards[0].Value()]
	secondCardBoardOrder := s.board.valueOrder[s.hand.cards[1].Value()]
	return s.hand.maxSameValues == 1 &&
		firstCardBoardOrder != 0 &&
		secondCardBoardOrder != 0 &&
		firstCardBoardOrder <= 2 &&
		secondCardBoardOrder <= 2
}

func findMediumTwoPairs(s *Selector) bool {
	firstCardBoardOrder := s.board.valueOrder[s.hand.cards[0].Value()]
	secondCardBoardOrder := s.board.valueOrder[s.hand.cards[1].Value()]
	return s.hand.maxSameValues == 1 &&
		firstCardBoardOrder != 0 &&
		secondCardBoardOrder != 0 &&
		(firstCardBoardOrder == len(s.board.cards) || secondCardBoardOrder == len(s.board.cards))
}

func findPair(s *Selector) bool {
	return s.total.maxSameValues == 2
}
