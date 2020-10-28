package main

import (
	"fmt"
	"math"
)

func exercise() {
	exercise01()
	defer exercise02()
	exercise03()
	exercise04()
	exercise05()
	exercise06()
	exercise07()
	exercise08()
	exercise09()
	exercise010()
	exercise11()
}
func exercise01() {
	fmt.Println("****** running exercise 01 ****")
	i := fooe()
	fmt.Println("fooe ==>", i)
	j, s := baar()
	fmt.Println("baar ==>", j, s)
}

func fooe() int {
	return 42
}
func baar() (int, string) {
	return 42, "Hello from bar"
}

func exercise02() {
	fmt.Println("****** running exercise 02 ****")
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(foeV(x...), barV(x))
}

func foeV(i ...int) int {
	var sum int
	for _, i := range i {
		sum += i
	}
	return sum
}
func barV(x []int) int {
	var sum int
	for i := 0; i < len(x); i++ {
		sum += x[i]
	}
	return sum
}

func exercise03() {
	defer deferexample()
	fmt.Println("****** running exercise 03 ****")
}

func deferexample() {
	defer func() {
		fmt.Println("Executing defer of annynomous function ")
	}()
	fmt.Println("About to exit deferFunction")
}

type person3 struct {
	first string
	last  string
	age   int
}

func (p person3) speak3() {
	fmt.Printf("Hello my name is %v having age %v \n", p.first, p.age)
}
func exercise04() {
	fmt.Println("****** running exercise 04 ****")
	p := person3{
		first: "James",
		last:  "Bond",
		age:   53,
	}
	p.speak3()
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("Area of type %T is %v \n", s, s.area())
}
func exercise05() {
	fmt.Println("****** running exercise 05 ****")
	s := square{
		side: 5,
	}
	c := circle{
		radius: 5,
	}
	info(s)
	info(c)
}

func exercise06() {
	fmt.Println("****** running exercise 06 ****")
	f := func(x int) {
		fmt.Println("Called anynomous func passed value ", x)
	}
	f(42)
	fmt.Printf("Type of f  %T \n", f)
}

func exercise07() {
	fmt.Println("****** running exercise 07 ****")
	f2 = func() {
		fmt.Println("function created inside method")
	}
	fmt.Printf("Type of f2  %T \n", f2)
	f2()
}

var f2 func()

func exercise08() {
	fmt.Println("****** running exercise 08 ****")
	f := funcReturnExample()
	f()
	f()
	f2 := funcReturnExample()
	f2()
}

func funcReturnExample() func() {
	x := 101
	return func() {
		x++
		fmt.Println("This is example of returned function ", x)
	}
}
func exercise09() {
	fmt.Println("****** running exercise 09 ****")
	fmt.Println("causing unnecessary complexity", acceptingFuncArgs(sq, 12))
}

func acceptingFuncArgs(f func(i int) int, x int) int {
	return f(x)
}

var sq = func(x int) int {
	return x * x
}

func exercise010() {
	fmt.Println("****** running exercise 10 ****")
	f := enclosureDemo()
	f()
	f()
	fmt.Println("value of f is ", f())
	f2 := enclosureDemo()
	f2()
	f2()
	fmt.Println("value of f2 is ", f2())
}

func enclosureDemo() func() int {
	x := 1
	return func() int {
		x++
		return x
	}
}

func exercise11() {
	fmt.Println("****** running recursion exercise 11 ****")
	fmt.Println(recusrionDemo(5))
}

func recusrionDemo(x int) int {
	if x == 0 {
		return 1
	}
	return x * recusrionDemo(x-1)
}
