package combinations

func sameValuesComb(s *Selector, repeatCount uint8, combName CombinationName) (c Combination, found bool) {
	idx := repeatCount
	for invertedVal, valCount := range s.invertedValues {
		val := uint8(12 - invertedVal)
		if valCount >= repeatCount && !found {
			for i := uint8(0); i < repeatCount; i++ {
				c.values[i] = val
				valCount--
			}
			c.name = combName
			found = true
		}
		for valCount > 0 && idx < 5 {
			c.values[idx] = val
			valCount--
			idx++
		}
	}
	return
}

func findHighComb(s *Selector) (Combination, bool) {
	return sameValuesComb(s, 1, High)
}

func findPairComb(s *Selector) (Combination, bool) {
	return sameValuesComb(s, 2, Pair)
}

func findSetComb(s *Selector) (Combination, bool) {
	return sameValuesComb(s, 3, Set)
}

func findQuadsComb(s *Selector) (Combination, bool) {
	return sameValuesComb(s, 4, Quads)
}

func findTwoPairsComb(s *Selector) (c Combination, found bool) {
	fistPairFound := false
	idx := uint8(4)
	for invertedVal, valCount := range s.invertedValues {
		val := uint8(12 - invertedVal)
		if valCount >= 2 && !found {
			valCount -= 2
			if !fistPairFound {
				c.values[0] = val
				c.values[1] = val
				fistPairFound = true
			} else {
				c.values[2] = val
				c.values[3] = val

				c.name = TwoPairs
				found = true
			}
		}
		if valCount > 0 && idx < 5 {
			c.values[idx] = val
			idx++
		}
	}
	return
}
func findFullHouseComb(s *Selector) (c Combination, found bool) {
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

func findStraightComb(s *Selector) (c Combination, round bool) {
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

func findFlushComb(s *Selector) (c Combination, found bool) {
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

func findStraightFlushComb(s *Selector) (c Combination, found bool) {
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
