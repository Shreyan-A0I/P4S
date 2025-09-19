package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := make([]int, 3)
	a = append(a, 5)
	fmt.Println(a)
}

func FindEvenDivisors(n int) []int {
	evenDivisors := make([]int,0)

	for i := 2; i <= n; i++ {
		d := n % i
		if d==0 && i%2==0 {
			evenDivisors = append(evenDivisors, i)
		}
	}

	return evenDivisors
}

func MakeStateFrequencyMap(board [][]int) map[int]int {
	countMap := make(map[int]int) 

	// loop through each row and column and add to countMap for each Key
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			countMap[board[i][j]]++
		}
	}

	return countMap
}


func GenerateRandomDNAString(length int, p1, p2, p3, p4 float64) string {
	s := ""

	//looping through length generating a random float each time and appending character to string based on that
	for i:=0; i<length; i++ {
	
		r := rand.Float64()
		
		if r<p1 {
			s += "A"
		} else if r<p1+p2{
			s += "C"
		} else if r<p1+p2+p3{
			s += "G"
		} else {
			s += "T"
		}
	}

	return s
}


func FindFlankingRegions(text, pattern1, pattern2 string)([]string){
	flankingRegions := make([]string,0)

	pat1, pat2 := startIndices(text, pattern1, pattern2)

	l1 := len(pattern1)
	l2 := len(pattern2)

	for _,i1 := range pat1{
		for _,i2 := range pat2{
			// pattern2 start should not be before pattern1 ends
			if i1+l1<=i2 && i2+l2 <= len(text){
				flankingRegions = append(flankingRegions, text[i1:i2+l2])
			}
		}
	}
	return flankingRegions
}

// return starting indices of two patterns
func startIndices(text, pattern1, pattern2 string) ([]int, []int) {
	pat1 := make([]int,0)
	pat2 := make([]int,0)

	l1 := len(pattern1)
	l2 := len(pattern2)

	// check pattern 1
	for i := 0; i <= len(text) - l1; i++ {
		if compareToPattern(text[i:i+l1], pattern1) {
			pat1 = append(pat1, i)
		}
	}

	for i := 0; i <= len(text) - l2; i++ {
		if compareToPattern(text[i:i+l2], pattern2) {
			pat2 = append(pat2, i)
		}
	}

	return pat1, pat2
}

//function to compare two strings
func compareToPattern(str, pattern string) bool {
	if len(str) != len(pattern) {
		panic("comparing strings of unequal length")
	}

	for i:=0; i<len(str); i++ {
		if str[i] != pattern[i] {
			return false
		}
	}

	return true
}