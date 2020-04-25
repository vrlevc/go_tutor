package test

import (
	"fmt"
	"testing"
)

func SqrtEx(x float64) float64 {
	if z := x; x > 0.0 {
		for i := 0; i < 100; i++ {
			if zn := z - (z*z-x)/(2.0*z); z-zn < 0.000000050001 {
				fmt.Printf("%v iterations done\n", i)
				break
			} else {
				z = zn
			}
		}
		return z
	}
	return -1
}

func TestSqrtEx(t *testing.T) {
	fmt.Println(SqrtEx(2))
	fmt.Println(sqrt(2))
}
