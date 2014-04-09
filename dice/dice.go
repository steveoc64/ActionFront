package dice

import (
	"math/rand"
)

func DieRoll() int {
	d1 := rand.Intn(9)
	d2 := rand.Intn(9)
	return 2 + d1 + d2
}

func D6() int {
	return rand.Intn(6) + 1
}

func D12() int {
	return rand.Intn(12) + 1
}

// Throw a bucket of D6, and return the number of HITs, where dice >= the score needed
func BucketD6(numDice int, scoreToHit int) int {

	numHits := 0
	for i := 0; i < numDice; i++ {
		if D6() >= scoreToHit {
			numHits++
		}
	}
	return numHits
}

// Throw a bucket of D12, and return the number of HITs, where dice <= the score needed
// NOTE - this is the reverse of a D6 roll above !!!!
func BucketD12(numDice int, scoreToHit int) int {

	numHits := 0
	for i := 0; i < numDice; i++ {
		if D12() <= scoreToHit {
			numHits++
		}
	}
	return numHits
}

func ShootDice(ammoOut int) (int, bool) {
	d1 := rand.Intn(9) + 1
	d2 := rand.Intn(9) + 1

	if d1 <= ammoOut {
		return d1 + d2, true
	} else {
		return d1 + d2, false
	}
}

func Percent(p int) bool {

	d := rand.Intn(99)
	if d < p {
		return true
	}
	return false
}
