package dice

import (
	"math/rand"
)

func DieRoll() int {
	d1 := rand.Intn(9)
	d2 := rand.Intn(9)
	return 2 + d1 + d2
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
