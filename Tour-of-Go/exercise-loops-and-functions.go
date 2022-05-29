package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for {
		dz := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(dz-z) < 1e-10 {
			break
		}
	}
	return z
}

func main() {
	for i := 1.0; i < 10; i++ {
		fmt.Println(Sqrt(i))
		fmt.Println(math.Sqrt(i))
	}
}
