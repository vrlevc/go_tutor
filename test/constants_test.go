package test

import (
	"fmt"
	"testing"
)

func TestConstants(t *testing.T) {
	const Pi = 3.14
	const World = "Червоні маки"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
