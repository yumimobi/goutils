package goutils

import (
	"math/rand"
	"time"
)

// can be mocked out for test
var randomIntn = rand.Intn

func init() {
	rand.Seed(time.Now().Unix())
}

// SetRandomIntn sets random function
func SetRandomIntn(f func(int) int) {
	randomIntn = f
}

// Modified from https://github.com/jmcvetta/randutil/blob/master/randutil.go

// A Choice contains a generic item and a weight controlling the frequency with
// which it will be selected.
type Choice struct {
	Weight int
	Item   interface{}
}

// WeightedChoice used weighted random selection to return one of the supplied
// choices.  Weights of 0 are never selected.  All other weight values are
// relative.  E.g. if you have two choices both weighted 3, they will be
// returned equally often; and each will be returned 3 times as often as a
// choice weighted 1.
func WeightedChoice(choices []Choice) (int, Choice) {
	// Based on this algorithm:
	//     http://eli.thegreenplace.net/2010/01/22/weighted-random-generation-in-python/
	sum := 0
	for _, c := range choices {
		sum += c.Weight
	}
	if sum <= 0 {
		l := len(choices)
		if l == 0 {
			return -1, Choice{}
		}

		randomIndex := randomIntn(l)
		return randomIndex, choices[randomIndex]
	}

	n := randomIntn(sum)
	for i, c := range choices {
		n -= c.Weight
		if n < 0 {
			return i, c
		}
	}

	panic("code can never run here")
}
