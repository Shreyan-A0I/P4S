package main

import ("log";"time"; "math")

func main () {
	// x := 219948
	// y := 828418

	//timing
	// start := time.Now()
	// fmt.Println(trivialGCD(x,y))
	// elapsed := time.Since(start)
	// log.Printf("elapsed time: %s", elapsed)

	// start2 := time.Now()
	// fmt.Println(euclidGCD(x,y))
	// elapsed2 := time.Since(start2)
	// log.Printf("elapsed time: %s", elapsed2)

	start := time.Now()
	trivialprimefinder(1000000)
	elapsed := time.Since(start)
	log.Printf("elapsed time: %s", elapsed)

	start2 := time.Now()
	sieveoferatos(1000000)
	elapsed2 := time.Since(start2)
	log.Printf("elapsed time: %s", elapsed2)
}

func min(a, b int) int {
	if a<b {return a}
	return a
}


// trivial GCD
func trivialGCD(a,b int) int {
	d := 1

	m := min(a,b)
	for i:=1; i<=m; i++ {
		if a%i == 0 && b%i == 0 {
			d = i
		}
	}
	return d
}

func euclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a-b
		} else {
			b = b-a
		}
	}
	return a
}

/*
func factorialarray (n int) [n+1]int {
	if n<0 {
		panic("error")
	}
	var fact [n+1]int
	fact[0]=1
	for i:=1; i<=n; i++ {
		fact[i] = fact[i-1]*i
	}

	return fact
}
*/

func factorialarray (n int) []int {
	if n<0 {
		panic("error")
	}
	fact := make([]int, n+1)
	fact[0]=1
	for i:=1; i<=n; i++ {
		fact[i] = fact[i-1]*i
	}

	return fact
}

//slice input output min
func minintarray (list []int) int {
	if len(list) == 0 {
		panic("empty slice given as input")
	}
	m := list[0]

	// for i,val := range list
	for i:=1; i<len(list); i++ {
		if list[i] < m {
			m = list[i]
		}
	}
	return m
}

/*variadic function
func minint(n ...int) int {
	//n becomes a slice of variable length
}*/

func trivialprimefinder (n int) []bool {
	PrimeBool := make([]bool, n+1)
	for p:=2; p<=n; p++ {
		PrimeBool[p] = isPrime(p)
	}
	return PrimeBool
}

func isPrime (p int) bool {
	for k:=2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k ==0 {
			return false
		}
	}
	return true
}

func sieveoferatos (n int) []bool {
	PrimeBool := make([]bool, n+1)

	for i:=2; i<=n; i++ {
		PrimeBool[i] = true
	}

	for p:=2; float64(p)<=math.Sqrt(float64(n)); p++ {
		if PrimeBool[p] {
			crossoffmultiples(PrimeBool, p)
		}
	}

	return PrimeBool
}

func crossoffmultiples (pb []bool, p int) []bool {
	n:=len(pb)-1
	for k:=2*p; k<=n; k+=p {
		pb[k] = false
	}
	return pb
}

func ListPrimes(n int) []int {
    pb := sieveoferatos(n)
    plist := make([]int, 0)
    for i,v:= range pb {
        if v {
            plist = append(plist,i)
        }
    }
    return plist
}