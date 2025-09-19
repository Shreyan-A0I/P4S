package main

import ( "fmt" )

func main() {
	T = min(2,3)
	if T {
		fmt.Println("b is less than a")
	} else {
		fmt.Println("a is not less than b")
	}
}

func min2(a, b int) bool {
	return a>b
}

func samesign(a, b int) bool {
	return a*b>=0
}

func mod(a, b int) int {
	if a>b{
		return a-b 
	}
	return b-a
}