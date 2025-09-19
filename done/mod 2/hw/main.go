package main

import (
	"fmt"
)

func main()  {

}

func SimpsonsIndex(sample map[string]int) float64 {
	// n := Richness(sample)
	t := SumOfValues(sample)

	var SI float64

	for _,val := range sample {
		SI += float64(val)/float64(t)*float64(val)/float64(t)
	}

	return SI
}

func Richness(sample map[string]int) int {
	count := 0
	for _,val := range sample {
		if val > 0 {
			count++
		}
	}
	return count
}



func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	s1 := SumOfValues(sample1)
	s2 := SumOfValues(sample2)
	avg := float64(s1+s2)/2.0
	sum_min := SumOfMinima(sample1,sample2)

    return 1-float64(sum_min)/float64(avg)
}

func SumOfValues(sample map[string]int) int {
	sum := 0
	for _,v := range sample {
		sum += v
	}
	return sum
}

func JaccardDistance(sample1, sample2 map[string]int) float64 {
	sum_min := SumOfMinima(sample1,sample2)
	sum_max := SumOfMaxima(sample1,sample2)

	return 1.0 - float64(sum_min)/float64(sum_max)
}

func SumOfMinima(sample1, sample2 map[string]int) int {
	//to keep track of species that we have seen already
	crossoff := make(map[string]bool, 0)
	sum := 0

	//ranging over first sample
	for i,_ := range sample1 {
		sum += Min2(sample1[i],sample2[i])
		crossoff[i] = true
	}

	for i,val := range sample2 {
		if !crossoff[i] {
            sum += Min2(0,val)
		}
	}

	return sum
}

func SumOfMaxima(sample1, sample2 map[string]int) int {
	//to keep track of species that we have seen already
	crossoff := make(map[string]bool, 0)
	sum := 0

	//ranging over first sample
	for i,_ := range sample1 {
		sum += Max2(sample1[i],sample2[i])
		crossoff[i] = true
	}

	for i,val := range sample2 {
		if !crossoff[i] {
            sum += val
		}
	}

	return sum
}

// Note: for the sake of convenience, we are providing a Min2() function below.
func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}