package test

import (
	"fmt"
	"testing"
)

func TestInference(t *testing.T) {
	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("%v is type of: %T\n", i, i)
	fmt.Printf("%v is type of: %T\n", f, f)
	fmt.Printf("%v is type of: %T\n", g, g)
}
