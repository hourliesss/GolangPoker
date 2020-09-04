package main

import (
	"math"
	"math/rand"
	"time"
)

func expoFunction(odds float64, index float64) float64 {
	return math.Exp(index * (odds - 1.0))
}

func raiseExpo(x float64) float64 {
	return (1.5 - 0.5*x) * math.Exp(x-1.0)
}

func normaleLawRandomization(risk float64, offset float64) float64 {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	riskInt := int(math.Ceil(risk * 100))
	randNumber := r.Intn(riskInt*2) - riskInt
	x := float64(randNumber)/100 + offset
	if x < 0 {
		return math.Exp(-0.5 * x * x)
	}

	return 2.0 - math.Exp(-0.5*x*x)

}

func ratioStacks(stack float64, opponentStack float64) float64 {
	return (math.Log10(stack/opponentStack) + 2.0) / 2.0
}

func ratioToWintoBet(amountToBet float64, pot float64) float64{
	return pot/amountToBet
}


