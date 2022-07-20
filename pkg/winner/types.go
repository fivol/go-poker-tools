package winner

type CombinationName uint8

const (
	High CombinationName = iota
	Pair
	TwoPairs
	Set
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
)

type Combination struct {
	name   CombinationName
	values [5]uint8
}

func (c Combination) GraterThen(other Combination) bool {
	if c.name > other.name {
		return true
	}
	if c.name < other.name {
		return false
	}
	for i := 0; i < 5; i++ {
		if c.values[i] > other.values[i] {
			return true
		}
		if c.values[i] < other.values[i] {
			return false
		}
	}
	return false
}
