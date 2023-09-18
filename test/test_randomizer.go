package test

import "math/rand"

const SEED = 1040
const RANDOM_ITERATIONS = 1000
const DELTA = 1e-7

func NewRandom() *rand.Rand {
	return rand.New(rand.NewSource(SEED))
}
