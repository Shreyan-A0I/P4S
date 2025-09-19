package main

import ("fmt";"math")

func main () {
	// fmt.Println((Permutation(5,1)))
	fmt.Println(ListMersennePrimes(61))
}

func Permutation(n, k int) int {
    return factorial(n)/factorial(n-k)
}

func Combination(n, k int) int {
	return factorial(n)/factorial(n-k)/factorial(k)
}

//subroutine for factorial
func factorial(n int) int {
	if n<0 {
		panic("error: int less than 0")
	}
	f := 1
	for i:=2; i<=n; i++ {
		f *= i
	}
	return f
}

func SumProperDivisors(n int) int {
    sd:=0
    for i:=1; i<n; i++ {
        if n%i==0 {
            sd+=i
        }
    }
    return sd
}

func IsPerfect(n int) bool {
    if SumProperDivisors(n) == n {
        return true}
    return false
}

func NextPerfectNumber(n int) int {
	n++
	for !IsPerfect(n) {
		n++
	}
	return n
}

func FibonacciArray(n int) []int {
	fib_array := make([]int, n+1)
	fib_array[0] = 1
    if n+1>1{
        fib_array[1] = 1}
	for i := 2; i<=n; i++ {
		fib_array[i] = fib_array[i-1] + fib_array[i-2]
	}
	return fib_array
}

func DividesAll(a []int, d int) bool {
	if d == 0 {
		return false
	}
	for i := range a {
		if a[i]%d!=0 {
			return false
		}
	}
	return true
}

func MaxIntegerArray(list []int) int {
	var max int
	if len(list) > 0 {
		max = list[0]
	}

	for i := 1; i<len(list); i++ {
		if max<list[i] {
			max = list[i]
		}
	}
	return max
}

func GCDArray(a []int) int {
	n := MaxIntegerArray(a)
	gcd := 1
	for i := 2; i<n; i++ {
		if DividesAll(a,i){
			gcd = i
		}
	}
	return gcd
}

func ListMersennePrimes(n int) []int {
	list := make([]int, 0)
	t := 0
	for i := 2; i<=n; i++ {
		t = Power(2, i) - 1
		if isPrime(t) {
			list = append(list, t)
		}
	}
	return list
}

func Power(a, b int) int {
    if b == 0 {
        return 1
    }
    pow := a
    for i:=2; i<=b; i++ {
        pow *= a
    }
    return pow
}

func isPrime (p int) bool {
	for k:=2; float64(k) <= math.Sqrt(float64(p)); k++ {
		if p%k ==0 {
			return false
		}
	}
	return true
}

func NextPrime(n int) int {
	for n=n+1; !isPrime(n); n++ {}
	return n
}

func NextTwinPrimes(n int) (int, int) {
	new := n+1
	var old int
	for new-old != 2 {
		old = new
		new = NextPrime(new)
	}
	return old, new
}

func Permutation(n, k int) int {
	if n<0 || k<0 || k>n {
		panic("error: int less than 0")
	}
    return factorial_opt(n,n-k)
}

func Combination(n, k int) int {
	if n<0 || k<0 || k>n {
		panic("error: int less than 0")
	}
	return factorial_opt(n,min(k,n-k))/factorial(max(k,n-k))
}

//subroutine for factorial
func factorial(n int) int {
	if n<0 {
		panic("error: int less than 0")
	}
	if n == 0 {
		return 1
	}
	f := 1
	for i:=2; i<=n; i++ {
		f *= i
	}
	return f
}

func factorial_opt(a,b int) int {
	f := 1
	for i:=a; i>b; i-- {
		f *= i
	}
	return f
}

func min (a, b int) int {
	if a>b {
		return b
	}
	return a
}

func max (a, b int) int {
	if a>b {
		return a
	}
	return b
}