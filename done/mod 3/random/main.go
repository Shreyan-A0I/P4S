package main

import(
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int()) //print integer
	fmt.Println(rand.Intn(10)) //print integer between 0 to 9
	fmt.Println(rand.Float64()) //print decimal [0,1)

}

func RollDie() int {
	return rand.Intn(6)+1
}

func SumTwoDice() int {
	return RollDie()+RollDie()
}

func SumDice(n int) int {
	sum := 0
	for i := 0; i<n; i++ {
		sum += RollDie()
	}
	return sum
}

func PlayCrapsOnce() bool {
	firstroll := SumDice(2)
	if firstroll == 7 || firstroll == 11 {
		return true
	} else if firstroll == 2 || firstroll == 3 || firstroll == 12 {
		return false
	} else {
		for {
			nr := SumDice(2)
			if nr == firstroll {
				return true
			} else if nr == 7 {
				return false
			}
		}
	}
}

func ComputeCrapsHouseEdge(numTrials int) float64 {
	count := 0
	for i:=0; i<numTrials; i++ {
		if PlayCrapsOnce() {
			count++
		} else {
			count--
		}
	}
	return float64(count)/float64(numTrials)
}