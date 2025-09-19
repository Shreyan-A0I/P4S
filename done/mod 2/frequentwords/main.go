package main

import (
	"fmt"
	"net/http"
	"io"
	"log"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg" //for vector graphics and image generation
)

func main() {
	fmt.Println("hello")
	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"
	resp, err := http.Get(url)
	//access url
	//close connection at the end of func main
	defer resp.Body.Close()

	if err!=nil {
		panic(err)
	}

	//check status
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK status: %v", resp.Status)
	}

	//io read
	GenomeSymbols, err := io.ReadAll(resp.Body)
	if err!=nil {
		panic(err)
	}

	EcoliGenome := string(GenomeSymbols)
	k:=9
	L:=500
	t:=3
	clumps := FindClumpsFaster(EcoliGenome, k, L, t)
	fmt.Println("found clumps of length", len(clumps))
}


// takes a string text and an integer k

func FindFrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)
	
	freqMap := FrequencyTable(text,k)
	
	max := MaxMapValue(freqMap)

	for pattern, val := range freqMap {
		if val == max {
			freqPatterns = append(freqPatterns, pattern)
		}
	}

	return freqPatterns
}

func FrequencyTable(text string, k int) map[string]int {
	dict := make(map[string]int)
    l := len(text)
	for i := 0; i <= l-k; i++ {
		pattern := text[i:i+k]

	/*	_, exists := freqMap[pattern]
		if !exists {
			freqMap[pattern] = 1
		} else {
			freqMap[pattern]++
		}
	}*/
		dict[pattern]++
	}
    return dict
}

func MaxMapValue(dict map[string]int) int {
	max := 0
	first := true
	for _,v := range dict {
		if first || max<v {
			max = v
			first = false
		}
	}
	return max
}

func FindClumps(text string, k, L, t int) []string {
	patterns := make([]string,0)
	n := len(text)
	foundPatterns := make(map[string]bool)

	for i := 0; i < n-L+1; i++ {
		window := text[i:i+L]
		FreqMap := FrequencyTable(window, k)

		for s, val := range FreqMap{
			if val >= t && !foundPatterns[s] {
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}
	} 
	return patterns
}

func FindClumpsFaster(text string, k, L, t int) []string {
	patterns := make([]string,0)
	n := len(text)
	foundPatterns := make(map[string]bool)

	FirstWindow := text[:L]
	FreqMap := FrequencyTable(FirstWindow, k)

	for s,freq := range FreqMap {
		if freq >= t {
			patterns = append(patterns,s)
			foundPatterns[s] = true
		}
	}

	for i := 1; i < n-L+1; i++ {
		//decrease value by 1 of firstsubstring and add last substring
		oldpattern := text[i-1:i-1+k]
		newpattern := text[i+L-k:i+L]

		FreqMap[oldpattern]--
		if FreqMap[oldpattern] == 0 {
			delete(FreqMap, oldpattern)
		}
		FreqMap[newpattern]++
		for s, val := range FreqMap{
			if val >= t && !foundPatterns[s] {
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}
	}
	return patterns
}

// func contains(patterns []string, s string) bool {
// 	for _, pattern := range patterns {
// 		if pattern == s {
// 			return true
// 		}
// 	}
// 	return false
// }


//skewarray
//input: DNA genome
//output: skew slice of integer showing G-C at each position
func SkewArray(genome string) []int {
	n:=len(genome)

	array := make([]int,n+1)
	array[0] = 0

	//range over remainng values
	for i := 1; i < n+1; i++ {
		array[i] = array[i-1] + Skew(genome[i-1])
	}

	return array
}

func Skew(symbol byte) int {
	if symbol == 'G' {
		return 1
	} else if symbol == 'C' {
		return -1
	}
	return 0
}

func MinSkew(genome string) []int {
	indices := make([]int, 0)

	arr := SkewArray(genome)

	m := Minint(arr)

	for i,val := range arr {
		if m == val {
			indices = append(indices, i)
		}
	}

	return indices
}

func Minint(skewarr []int) int {
	min := skewarr[0]

	for i := range skewarr {
		if skewarr[i] < min {
			min = skewarr[i]
		}
	}
	return min
}