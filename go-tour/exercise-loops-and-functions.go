package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	prev := 0.0
	z := 1.0
	for math.Abs(z-prev) > .005 {
		fmt.Println(z)
		prev = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
