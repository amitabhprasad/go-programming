package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) Error() string {
	var s string
	if p.age < 0 {
		s = "Age can not be less then 0 "
	}
	return s
}

func main() {
	examples()
	//exampleUsingWaitGroup()
}

func customError() {
	p := person{
		first: "John",
		age:   1,
	}
	fmt.Println("Person ", p.Error())
}
