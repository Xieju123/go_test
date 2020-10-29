package main

import (
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(x-z*z)>1e-5{
		z+=(x-z*z)/(2*z)
	}
	return z
}

func main() {

}
