package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) interface{} {
	if x < 0 {
		return ErrNegativeSqrt(x).Error()
	}
	z := float64(2.)
	s := float64(0)
	for {
		z = z - (z*z-x)/(2*z)
		if math.Abs(s-z) < 1e-15 {
			break
		}
		s = z
	}
	return s
}

func main() {
	fmt.Println(Sqrt(-2))
	fmt.Println(math.Sqrt(2))
}
