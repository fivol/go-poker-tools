package combinations

type FD uint8
type FDList []FD

func (s *Selector) getFDList() FDList {
	var fds [13]bool

	for suit := uint8(0); suit < 4; suit++ {
		suitsCount := s.board.suits[suit]
		if suitsCount == 3 || suitsCount == 2 {
			fdCount := 0
			for i := 12; i >= 0 && fdCount < 4; i-- {
				if s.board.values[i] == 0 {
					fdCount++
					fds[i] = true
				}
			}
		}
	}
	var fdList FDList
	for i := 12; i >= 0; i-- {
		if fds[i] {
			fdList = append(fdList, FD(i))
		}
	}
	return fdList
}
