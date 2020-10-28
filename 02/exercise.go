package main

import "fmt"

func exercise() {
	exercise01()
	exercise03()
	exercise04()
	exercise05()
}

func exercise01() {
	fmt.Println("****** running exercise 01 ****")
	x := 42
	fmt.Printf("value of x in decimal %v\n", x)
	fmt.Printf("value of x in binary %b\n", x)
	fmt.Printf("value of x in Hexadecimal %#X \n", x)
	fmt.Printf("value of x in UTF %#U \n", x)
}

const (
	a3     = 42
	b3 int = 42
)

func exercise03() {
	fmt.Println("****** running exercise 03 ****")
	fmt.Println(a3, b3)
}

func exercise04() {
	fmt.Println("****** running exercise 04 ****")
	x := 42
	fmt.Printf("%d \t %b \t %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d \t %b \t %#x\n", y, y, y)

	z := x >> 1
	fmt.Printf("%d \t %b \t %#x\n", z, z, z)
}

const (
	_ = iota
	y1
	y2
	y3
	y4
	current = 2020
)

func exercise05() {
	fmt.Println("****** running exercise 05 ****")
	fmt.Println(current + y1)
	fmt.Println(current + y2)
	fmt.Println(current + y3)
	fmt.Println(current + y4)

}
