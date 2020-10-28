package main

/**
* Go natively take advantage of multiple cores
	2006 - multi core processor was first released
	2007 - work on Go dev started

- Parallelism
	- Code can run in parallel, happens only when computer has more then one core
- Concurrency
	- Design pattern, that allows your code to run in parallel, Concurrent code is required
	to run in multi core machine to run in parallel
**/

import (
	"fmt"
)

type myString string

type formatting interface {
	toString()
}

type hello interface {
	sayHello()
}

func (s myString) sayHello() {
	fmt.Printf("Hello \t %v \n", s)
}

func (s *myString) toString() {
	fmt.Printf("Custom formatting \n type \t\t value \n %T\t %v \n", s, s)
}

func sayHelloImplByType(h hello) {
	h.sayHello()
}

func formatterImplByPointer(f formatting) {
	f.toString()
}

// methodset of a type determines the interfaces that the types implements
// if it's the pointer type that, implements interface then method through interface
// can be called only by using pointers

// if interface is implemented by type then method through interface can be called by
//pointer type or by type

// and the methods that can be called using a receiver of that type

// method set of myString are sayHello() & toString()
// method set of pointer type i.e. *myString are sayHello(), toString(),sayHelloPointer()
func methodsetExample() {
	s := myString("Represents type")
	sayHelloImplByType(s)
	//formatterImpl(s)
	fmt.Println("Running pointer example")
	s2 := myString("I am pointer")
	var sp *myString = &s2
	sayHelloImplByType(sp)
	formatterImplByPointer(sp)
}

func main() {
	//concurrencyExample()
	//methodsetExample()
	//raceConditionDemoUsingMutex()
	//raceConditionDemoUsingAtomic()
	exercise()
}
