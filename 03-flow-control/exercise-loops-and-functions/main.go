package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z := float64(1)

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Step %d: z = %v\n", i, z)
	}
	return z
}

func main() {
	fmt.Println("My sqrt(2) =", sqrt(2))
	fmt.Println("Math sqrt(2) =", math.Sqrt(2))
}
