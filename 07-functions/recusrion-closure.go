package main

import "fmt"

func closureRecusrionExample() {
	f := incrementor()
	fmt.Println("1st ", f())
	fmt.Println("2nd ", f())
	fmt.Println("3rd ", f())
	f2 := factorial(4)
	fmt.Printf("%T \t %v\n", f2, f2)
	fmt.Println("loopFactorial ", loopFactorial(4))
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

/**
	- recursion function calling itself
	- should have someway to stop that function
	- example factorial function
**/
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func foo(f func(n int) int, c int) int {
	v := f(c)
	return v
}
func loopFactorial(n int) int {
	var x = 1
	for ; n > 0; n-- {
		x *= n
	}
	return x
}
