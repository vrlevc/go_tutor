package test

import (
	"fmt"
	"math"
	"testing"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func TestPow(t *testing.T) {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
