package main

import "fmt"

func examples() {
	//receiverExample()
	anonymousExample()
	f("Hello expresion")
	rf := returnFunctionxample()
	fmt.Printf("Returing function type is %T and retrun values %v \n", rf, rf())
	//callbackExample()
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	si := sum(s...)
	fmt.Println(si)
	s2 := sumEven(sum, s...)
	fmt.Println("sum of all even numbers are ", s2)
	fmt.Println("sum of all odd numbers are ", sumOdd(sum, s...))
}

var sum = func(x ...int) int {
	s := 0
	for _, i := range x {
		s += i
	}
	return s
}

func sumEven(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func sumOdd(f func(x ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func anonymousExample() {
	fmt.Println("running  anonymousExample")

	func(s ...string) {
		fmt.Println("anonymous function called ", s)
	}("Hello", "Greet ", " meaning of life")

	//example of func expression
	g := func() {
		fmt.Println("combination of anonymous and expression")
	}
	g()
}

func returnFunctionxample() func() int {
	f := func() int {
		return 42
	}
	return f
}

var f = func(s ...string) {
	fmt.Println("funcExpression example", s)
}

type person struct {
	first string
	last  string
}

type person2 struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak(lang ...string)
}

func receiverExample() {
	fmt.Println("Running receiver example")
	sa := secretAgent{
		person: person{
			first: "James ",
			last:  "Bond",
		},
		ltk: true,
	}
	lang := []string{"English", "Hindi"}
	sa.speak(lang...)
	p := person{
		first: "Dr.",
		last:  "Yes",
	}
	p2 := person2{
		first: "Dr.2",
		last:  "Yes2",
	}
	bar(sa)
	bar(p)
	bar(p2)
}

func (p person) greet() string {
	return fmt.Sprintf("Hello from %v", p.first)
}

func (sa secretAgent) speak(lang ...string) {
	fmt.Printf("This secret agent is of the type %T\n", sa)
	fmt.Println(sa.greet(), "I can speak following languages ...")
	for i, v := range lang {
		fmt.Printf("%v\t%v\n", i, v)
	}
	sa.person.speak("Marathi")
}

func (p person) speak(lang ...string) {
	fmt.Printf("%v is of the type %T\n", p.first, p)
	fmt.Println("I can speak following languages ...")
	for i, v := range lang {
		fmt.Printf("%v\t%v\n", i, v)
	}
}

func (p person2) speak(lang ...string) {
	fmt.Printf("%v is of the type %T\n", p.first, p)
	fmt.Println("I can speak following languages ...", p)
}

type myInt int

func bar(h human) {
	fmt.Printf("Called from bar I am of type %T\n", h)
	switch h.(type) {
	case person:
		fmt.Println("person Coming from bar... ", h.(person).first) //assertion
	case person2:
		fmt.Println("person2 Coming from bar... ", h.(person2).first)
	case secretAgent:
		fmt.Println("secretAgent Coming from bar... ", h.(secretAgent).ltk)
	}
	if sa, ok := h.(secretAgent); ok {
		fmt.Println("Hello I am secretAgent, I have liscence to kill ", sa.ltk)
	} else {
		fmt.Println("Sorry  I am not secretAgent ")
	}
	//conversion example
	var m myInt = 42
	fmt.Printf("%T %v \n", m, m)

	y := int(m)
	fmt.Printf("%T %v \n", y, y)
}
