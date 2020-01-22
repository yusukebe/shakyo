package main

import(
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 0.1
	z2 := 0.0
	for {
		z2 -= (z*z - x) / (2*z)
		if math.Abs(z2 - z) < 0.0001 {
			break
		}
		z = z2
		//fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
