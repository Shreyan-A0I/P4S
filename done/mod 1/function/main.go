package main

import (
	"fmt"
)

func main() {
	x := 3
	fmt.Println("The value of x is:", add(x, 5))
}

// adds two integers and returns the result
func add(a int, b int) int {
	return a + b
}

// subtracts two integers and returns the result
func doubleduplicate(a float64) (float64, float64) {
	return a * 2.0, a * 2.0
}

// returns the value of pi
func pi() (float64) {
	return 3.14
}

// prints a greeting message
func hi() {
	fmt.Println("Hi there!")
}