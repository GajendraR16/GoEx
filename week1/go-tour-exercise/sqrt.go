package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN()
	}
	z := 1.0
	const epsilon = 1e-9

	for {
		prevZ := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prevZ) < epsilon {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
