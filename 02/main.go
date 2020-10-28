package main

import (
	"fmt"
	"runtime"
)

const (
	a = 12
	b = 12.12
	c = "James Bond"
)
const (
	a1 = iota
	b1
	c1
)

func main() {
	// ascii
	fmt.Println("********* Understanding ascii **********")
	x := 1
	y := 12.34
	fmt.Printf("%T, %T \n", x, y)
	fmt.Println(runtime.GOARCH, runtime.GOOS)
	s := "Hello, Playground"
	a := "A"
	bs := []byte(s)
	fmt.Println(bs, []byte(a))
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println()
	for i, v := range s {
		fmt.Printf("Formatted data at index position %d decimal value is %d, UTF-8 val is %#U, hex value is %#x, ASCII is %b\n", i, v, v, v, v)
	}
	i := 911
	fmt.Printf("Value of i in hexadecimal is %#X\n", i)
	fmt.Printf("Value of i in UTF-8 is %#U\n", i)
	fmt.Printf("Value of i in decimal is %d\n", i)
	fmt.Printf("Value of i in binary is %b\n", i)

	s1 := "H"
	fmt.Printf("User input data %v\n ", s1)
	bs1 := []byte(s1)
	fmt.Printf("byte slice %v\n ", bs1)
	n := bs1[0]

	fmt.Printf("Value of s1 in hexadecimal is %#X\n", n)
	fmt.Printf("Value of s1 in UTF-8 is %#U\n", n)
	fmt.Printf("Value of s1 in decimal is %d\n", n)
	fmt.Printf("Value of s1 in binary is %b\n", n)
	iotaDemo()
	exercise()
}

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb
	gb
)

func iotaDemo() {
	fmt.Println("IOTA value", b1)
	x := 4
	fmt.Printf("%d \t\t %b \n", x, x)

	y := x << 1
	fmt.Printf("%d \t\t %b \n", y, y)

	fmt.Printf("%d \t\t %b \n", kb, kb)
	fmt.Printf("%d \t %b \n", mb, mb)
	fmt.Printf("%d \t %b \n", gb, gb)
}
