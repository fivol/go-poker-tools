package random

import (
	"github.com/valyala/fastrand"
	"math/rand"
	"testing"
)

const max = 1000000

func BenchmarkRandomMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Intn(max)
	}
}

func BenchmarkFastRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fastrand.Uint32()
	}
}

func BenchmarkPseudoRandom(b *testing.B) {
	r := fastrand.RNG{}
	for i := 0; i < b.N; i++ {
		r.Uint32()
	}
}
