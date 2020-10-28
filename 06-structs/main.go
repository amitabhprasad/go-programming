package main

import (
	"fmt"
)

type person struct {
	first    string
	last     string
	age      int
	interest []string
}
type secretAgent struct {
	p   person
	ltk bool
}

func main() {
	p := person{
		first:    "James",
		last:     "Bond",
		age:      32,
		interest: []string{"hike", "run", "code"},
	}
	sa := secretAgent{
		p:   p,
		ltk: true,
	}
	fmt.Println("secret agent--> ", sa.p.first, " has ltk--> ", sa.ltk, "interest : ", sa.p.interest)
	ap := []person{
		{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		{
			first: "James2",
			last:  "Bond2",
			age:   32,
		},
	}
	fmt.Println(p.first)
	//fmt.Println(ap)
	for i, v := range ap {
		fmt.Println(i, v.first, v.age)
	}
	anonymousStructExample()
	exercise()
}

func anonymousStructExample() {
	p1 := struct {
		first string
		last  string
	}{
		first: "James",
		last:  "Bond",
	}
	fmt.Printf("%T data %v\n", p1, p1)
}
