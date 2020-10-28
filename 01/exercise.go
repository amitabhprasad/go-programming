package main

import "fmt"

func exercise() {
	exercise01()
	exercise02()
	exercise03()
	exercise04()
	exercise05()
}
func exercise01() {
	fmt.Println("****** running exercise 01 ****")
	x := 42
	y := "James Bond"
	z := true
	fmt.Printf("Values assigned are %v,%v,%v \n", x, y, z)
}

var x2 int
var y2 string
var z2 bool

func exercise02() {
	fmt.Println("****** running exercise 02 ****")
	fmt.Printf("Compiler assigned type for x2, y2, z2 are %T, %T, %T \n", x2, y2, z2)
	fmt.Printf("Compiler assigned values for x2, y2, z2 are %v, %v, %v \n", x2, y2, z2)
	x2 = 42
	y2 = "James Bond"
	z2 = true
	fmt.Printf("Values assigned using globals are  %v, %v, %v \n", x2, y2, z2)
}

var x3 int = 42
var y3 string = "James Bond"
var z3 bool = false

func exercise03() {
	fmt.Println("****** running exercise 03 ****")
	s := fmt.Sprintf("Values assigned using globals are  %v\t %v\t %v \n", x3, y3, z3)
	fmt.Println(s)
}

type myType int

var m myType

func exercise04() {
	fmt.Println("****** running exercise 04 ****")
	fmt.Println("Values of myType is ", m)
	fmt.Printf("Type of mType is %T\n", m)
	m = 42
	fmt.Println("Values of myType is ", m)
}

var i int

func exercise05() {
	fmt.Println("****** running exercise 05 ****")
	i := int(m)
	fmt.Println("Values of i from custom type is ", i)
	fmt.Printf("Type is %T\n", i)
}
