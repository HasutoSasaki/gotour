package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) float64 {
	z := float64(1)

	if x < 0 {
		err := ErrNegativeSqrt(x)
		fmt.Println(err)
		return 0 // or handle the error as needed
	}
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Step %d: z = %v\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
