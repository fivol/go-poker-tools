package combinations

import "go-poker-tools/pkg/types"

type Selector struct {
	hand   cardsInfo
	board  cardsInfo
	total  cardsInfo
	topFDs FDList
}

func newCombinationsSelector(board types.Board, hand types.Hand) (s Selector) {
	c1, c2 := hand.Cards()
	cards := make([]types.Card, len(board)+2)
	copy(cards, board)
	cards[len(board)] = c1
	cards[len(board)+1] = c2
	s.total = newCardsInfo(cards)
	s.hand = newCardsInfo(cards[len(board):])
	s.board = newCardsInfo(cards[:len(board)])
	s.topFDs = s.getFDList()
	return
}
