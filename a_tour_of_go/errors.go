package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	
	z0 := 0.0
	z1 := 1.0
	for math.Abs(z0 - z1) > 1e-10 {
		z0 = z1
		z1 = z1 - (z1*z1-x)/(2*z1)
	}
	
	return z1, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
