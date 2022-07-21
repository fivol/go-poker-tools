package generics

func Max(x, y uint8) uint8 {
	if x > y {
		return x
	}
	return y
}

func Min(x, y uint8) uint8 {
	if x < y {
		return x
	}
	return y
}
