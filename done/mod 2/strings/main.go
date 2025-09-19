package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Strings.")
	fmt.Println(strconv.Itoa(45))
	fmt.Println(ReverseComplement("ATTA"))
}

func ReverseComplement(pattern string) string {
	return Reverse(Complement(pattern))
}

func Complement (dna string) string {
	dna2 := make([]byte, len(dna))
	for i, symbol := range dna {
		switch symbol {
			case 'A':
				dna2[i] = 'T'
			case 'T':
				dna2[i] = 'A'
			case 'C':
				dna2[i] = 'G'
			case 'G':
				dna2[i] = 'C'
			default:
				panic("error not a neucleotide")
		}
	}
	return string(dna2)
}

func Reverse(pattern string) string {
	pat := make([]byte, len(pattern))
	n := len(pattern)

	for i := range pattern {
		pat[i] = pattern[n-1-i]
	}
	return string(pat)
}

func patterncount(pattern,text string) int {
	len(StartingIndices(pattern,text))
	// count := 0
	// k := len(pattern)

	// for i := 0; i<=len(text)-k; i++ {
	// 	if pattern == text[i:i+k]{
	// 		count++
	// 	}
	// }
	// return count
}

func StartingIndices(pattern, text string) []int {
	k := len(pattern)
	position := make([]int, 0)

	for i := 0; i<=len(text)-k; i++ {
		if pattern == text[i:i+k]{
			position = append(position, i)
		}
	}
	return position
}