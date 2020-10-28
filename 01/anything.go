package main

import (
	"fmt"
)

func main() {
	fmt.Println("File name can be anything but needs ", "package main", "and func main")
	foo()
	exercise()
	fmt.Println("Program exits when main() exits")
}

var x = 12

func foo() {
	fmt.Println(" I am in foo()", x)
	fmt.Printf("Type of x is %T binary is %b\n", x, x)
	fmt.Printf("Type of x in hexdecimal is %x\n", x)

	a := `It's said 
	"Growth and comfort do not co-exists ! "`
	fmt.Println(a)
	showType()
}

type hotdog int

var b hotdog

func showType() {
	x = 42
	b = 43
	fmt.Println(" I am in showType()", x, "balue of b is ", b)
	fmt.Printf("Type of x is %T\n", x)
	fmt.Printf("Type of b is %T\n ", b)
	x = int(b)
	fmt.Println("After conversion ", x, b)
}
