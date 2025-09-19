package main

import ( 
    "fmt"
    // "math"
    // "strconv"
    "math/rand" // this should be helpful!
)

func main() {
	// numTrials := 10000
	// fmt.Println("The value of pi is : ", EstimatePi(numTrials))
	// fmt.Println("relative prob : ", RelativelyPrimeProbability(1,1000000000,1000000))
	// fmt.Println("relative prob birthday pair for 40 people : ", SharedBirthdayProbability(23,10000))
	fmt.Println(CountShortPeriodSeeds())
}

func WeightedDie() int {
    F := rand.Float64()
    if F <0.1 {
        return 1
    } else if F<0.2 {
        return 2
    } else if F<0.7 {
        return 3
    } else if F<0.8 {
        return 4
    } else if F<0.9 {
        return 5
    }
    return 6
}

func EstimatePi(numPoints int) float64 {	
	hit := 0

	for i:=0; i<numPoints; i++ {
		a := rand.Float64()*2 - 1
		b := rand.Float64()*2 - 1

		if a*a+b*b <= 1 {
			hit++
		}
	}

	return 4.0*float64(hit)/float64(numPoints)
}

func EuclidGCD(a, b int) int {
    for a != b {
        if a > b {
            a = a - b
        } else {
            b = b-a
        }
    }
    return a
}

// write your improved RelativelyPrime() function here. Use EuclidGCD!
func RelativelyPrime(a, b int) bool {
    if a == 1 || b == 1 {
        return true
    }
    if EuclidGCD(a,b) == 1 {
        return true
    }
    return false
}

func RelativelyPrimeProbability(lowerBound, upperBound, numPairs int) float64 {
	l := upperBound-lowerBound+1
	count := 0

	for i:=0; i<numPairs; i++ {
		a := rand.Intn(l)+lowerBound
		b := rand.Intn(l)+lowerBound

		if RelativelyPrime(a,b)==true {
			count++
		}
	}
	return float64(count) / float64(numPairs)
}

func HasRepeat(a []int) bool {
	//checks if an integer is encountered
	checklist := make(map[int]bool)

	for _,i := range a {
		if checklist[i]{
			return true
		}
		checklist[i] = true
	}
	return false
}

func SimulateOneBirthdayTrial(num_people int) bool {
	bdays := make([]int,0)

	for i:=0; i<num_people; i++ {
		bdays = append(bdays,rand.Intn(365))
		if HasRepeat(bdays) {
			return true
		}
	}
	return false
}

func SharedBirthdayProbability(numPeople, numTrials int) float64 {
	count := 0
	for i:=0; i<numTrials; i++ {
		if SimulateOneBirthdayTrial(numPeople) {
			count++
		}
	}
	return float64(count)/float64(numTrials)
}

func CountNumDigits(x int) int {
    // Handle zero
    if x == 0 {
        return 1
    }

    // Make x positive for counting
    if x < 0 {
        x = -x
    }

    count := 0
    for x > 0 {
        x /= 10
        count++
    }
    return count
}

func SquareMiddle(x, numDigits int) int {
    // Fail conditions
    if numDigits <= 0 || numDigits%2 != 0 || x < 0 || CountNumDigits(x) > numDigits {
        return -1
    }

    squared := x * x
    leftTrim := numDigits / 2
    middleSection := (squared / Pow10(leftTrim)) % Pow10(numDigits)

    return middleSection
}

func Pow10(n int) int {
    result := 1
    for i := 0; i < n; i++ {
        result *= 10
    }
    return result
}

func GenerateMiddleSquareSequence(seed, numDigits int) []int {
	seq := make([]int, 0)
	seq = append(seq, seed)

	for {
        next := SquareMiddle(seed, numDigits)
		seq = append(seq, next)
		if HasRepeat(seq){
			return seq
		}
        seed = next
	}
}

func ComputePeriodLength(a []int) int {
	checklist := make(map[int]int)

	for dex,i := range a {
		_,exists := checklist[i]
		if exists {
			return dex - checklist[i] + 1
		}
		checklist[i] = dex
	}
	return 0
}


func CountShortPeriodSeeds() int {
    count := 0
    for seed := 1; seed <= 9999; seed++ {
        seq := GenerateMiddleSquareSequence(seed, 4)
        period := ComputePeriodLength(seq)
        if period <= 10 {
            count++
        }
    }
    return count
}

func GenerateLinearCongruenceSequence(seed, a, c, m int) []int {
    seq := make([]int,0)
	seq = append(seq,seed)

    for {
        next := (a*seq[len(seq)-1] + c) % m
        seq = append(seq, next)

        if HasRepeat(seq) {
            return seq // stop as soon as a repeat occurs, including the repeated element
        }
    }
}
