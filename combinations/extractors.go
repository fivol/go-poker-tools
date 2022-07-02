package combinations

func sameValuesComb(s *Selector, repeatCount uint8, combName CombinationName) (c Combination, found bool) {
	for invertedVal, valCount := range s.invertedValues {
		if valCount >= repeatCount {
			c.values[0] = uint8(12 - invertedVal)
			c.name = combName
			found = true
			return
		}
	}
	return
}

func FindHighComb(s *Selector) (Combination, bool) {
	return sameValuesComb(s, 1, High)
}

func FindPairComb(s *Selector) (Combination, bool) {
	return sameValuesComb(s, 2, Pair)
}

func FindSetComb(s *Selector) (Combination, bool) {
	return sameValuesComb(s, 3, Set)
}

func FindQuadsComb(s *Selector) (Combination, bool) {
	return sameValuesComb(s, 4, Quads)
}

func FindTwoPairsComb(s *Selector) (c Combination, found bool) {
	fistPairFound := false
	for invertedVal, valCount := range s.invertedValues {
		if valCount >= 2 {
			if !fistPairFound {
				c.values[0] = uint8(12 - invertedVal)
				c.name = TwoPairs
				fistPairFound = true
			} else {
				c.values[1] = uint8(12 - invertedVal)
				found = true
				return
			}
		}
	}
	return
}
func FindFullHouseComb(s *Selector) (c Combination, found bool) {
	sumEqualValuesCount := 0
	for invertedVal, valCount := range s.invertedValues {
		if valCount >= 2 {
			if sumEqualValuesCount != 0 && int(valCount)+sumEqualValuesCount < 5 {
				continue
			}
			if valCount == 3 {
				c.values[0] = uint8(12 - invertedVal)
			} else {
				c.values[1] = uint8(12 - invertedVal)
			}
			if sumEqualValuesCount > 0 {
				found = true
				c.name = FullHouse
				return
			}
			sumEqualValuesCount += int(valCount)
		}
	}
	return
}

func FindStraightComb(s *Selector) (c Combination, round bool) {
	var chainLen uint8
	for invertedVal, valCount := range s.invertedValues {
		if valCount == 0 {
			chainLen = 0
			continue
		}
		chainLen++
		if chainLen < 5 {
			continue
		}
		val := uint8(12 - invertedVal)
		round = true
		c.name = Straight
		c.values[0] = val + 4
		return
	}
	return
}

func FindFlushComb(s *Selector) (c Combination, found bool) {
	for suit, suitsCount := range s.suits {
		if suitsCount < 5 {
			continue
		}
		found = true
		c.name = Flush
		var sameSuitValues [13]bool
		for _, card := range s.cards {
			if card.Suit() == uint8(suit) {
				sameSuitValues[12-card.Value()] = true
			}
		}
		idx := 0
		for invertedValue, exist := range sameSuitValues {
			if !exist {
				continue
			}
			c.values[idx] = uint8(12 - invertedValue)
			idx++
			if idx == 5 {
				break
			}
		}
	}
	return
}

func FindStraightFlushComb(s *Selector) (c Combination, found bool) {
	for suit, suitsCount := range s.suits {
		if suitsCount < 5 {
			continue
		}
		var sameSuitValues [13]bool
		for _, card := range s.cards {
			if card.Suit() == uint8(suit) {
				sameSuitValues[12-card.Value()] = true
			}
		}
		var chainLen uint8
		for invertedValue, exist := range sameSuitValues {
			if !exist {
				chainLen = 0
				continue
			}
			chainLen++
			if chainLen < 5 {
				continue
			}
			c.values[0] = uint8(12 - invertedValue + 4)
			found = true
			c.name = StraightFlush
			return
		}
	}
	return
}
