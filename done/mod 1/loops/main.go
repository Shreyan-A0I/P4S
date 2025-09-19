package main

import ("fmt")

func main () {
	fmt.Println(factorial(66))
	fmt.Println(EuclidGCD(10,5))
}

func factorial(n int) int {
	if n<0 {
		panic("Error: negative input given to factorial")
	}

	p := 1
	q := 1
	for q<=n {
		p = p*q
		q = q+1
	}
	return p
}

func sum (n int) int{
	if n<0 {
		panic("Error: negative input given to factorial")
	}
	s := 0
	for n>0 {
		s = s + n
		n = n-1
	}
	return s
}

func EuclidGCD (a,b int) int {
	for a!=b{
	if a>b {
		a = a-b
	} else {
		b = b-a
	}}
	return a
}

// another factorial, takes input n and gives factorial n! as output
func anotherfact (n int) int {
	if n<0 {
		panic("error: integer less than 0")
	}
	p := 1 //initialzation
	for i := 1; i<=n; i++{
		p *= i
	}
	return p
}

func anothersum (n int) int {
	if n<0 {
		panic("error")
	}
	p := 0 
	for i := 1; i <=n; i++{
		p+=i
	}
	return p
}

func gausssum (n int) int {
	if n<0 {
		panic("error")
	}
	return (n*(n+1))/2
}

func sumeven (n int) int {
	if n<0 {
		panic("error")
	}
	p := 0 
	for i := 1; i <=n; i++{
		if i%2==0
			p+=i
	}
	return p
}