package main

import (
	"fmt"
	"math"
)

func main() {
	exercise()
	//example()
	methodSetsExampleNonPointerRecivers()
	methodSetsExamplePointerRecivers()
}

/**
	- All the methods attached to the type are known as methodsets
	- NON-Pointer receiver works with values that are pointers or non pointers
	- Pointer receiver only works with values that are pointers
**/
type shape interface {
	area() float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	fmt.Println("Calculating area of circle with radius ", c.radius)
	return c.radius * c.radius * math.Pi
}

func info(s shape) {
	fmt.Println("Area ", s.area())
}

func methodSetsExampleNonPointerRecivers() {
	c := circle{5}
	info(c)
	var x float64 = c.radius
	var y *float64 = &c.radius
	*y = 6
	fmt.Println(x, y, *y)

	fmt.Printf("Type address of circle & radius %v\t %T \t %T \t %v \n", &c, &c, &c.radius, &c.radius)
	info(&c)
}

func infoPointer(s shapePointer) {
	fmt.Println("Area ", s.areaPointer())
}

type shapePointer interface {
	areaPointer() float64
}

func (c *circle) areaPointer() float64 {
	fmt.Println("Calculating area of circle with Pointer ", c.radius)
	return c.radius * c.radius * math.Pi
}
func methodSetsExamplePointerRecivers() {
	c := circle{5}
	fmt.Printf("Type of circle %T \n", &c)
	infoPointer(&c)
}

type person struct {
	first  string
	second string
	age    int
}

func example() {

	a := 42
	fmt.Println(a, &a, *&a)
	fmt.Printf("Details of variable a %T , %T \n", &a, a)
	var b *int = &a
	fmt.Println(b, *b)
	s := "word"
	fmt.Println(&s, s)
	fmt.Printf("Type s %T \n", &s)
	var s2 *string = &s
	*s2 = "world"
	fmt.Println(*s2, s)

	p := person{
		first:  "James",
		second: "Bond",
		age:    42,
	}
	fmt.Println(&p)
	var p2 *person = &p
	fmt.Println(p2)
	p2.age = 52
	fmt.Println(p)
}
