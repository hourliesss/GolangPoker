package main

import (
	"math"
	"math/rand"
	"time"
)

func expoFunction(odds float64) float64 {
	return math.Exp(3 * (odds - 1))
}

func normaleLawRandomization(risk int) float64 {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	randNumber := r.Intn(risk*100*2) - 100*risk
	x := float64(randNumber) / 100
	if x < 0 {
		return math.Exp(-0.5 * x * x)
	}

	return 2.0 - math.Exp(-0.5*x*x)

}
