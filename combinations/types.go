package combinations

type CombinationName uint8

const (
	HIGH CombinationName = iota
	PAIR
	SET
	STRAIGHT
	FLUSH
	FULL_HOUSE
	QUADS
	STRAIGHT_FLUSH
	ROYAL_FLUSH
)

type Combination struct {
	name  CombinationName
	value uint8
}

func (c Combination) GraterThen(other Combination) bool {

}
