package test

import (
	"fmt"
	"testing"
)

func TestWhile(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
