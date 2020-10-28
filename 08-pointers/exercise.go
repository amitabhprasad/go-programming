package main

import "fmt"

func exercise() {
	exercise01()
	exercise02()
	fmt.Println("*********** exiting ******")
}

func exercise01() {
	fmt.Println("******* Running exercise 01")
	x := 42
	fmt.Printf("Details of variable x addressType: %T \t address:%v\t type: %T \n", &x, &x, x)
}

type person2 struct {
	first   string
	address string
}

func changeMe(p *person2, s string) {
	fmt.Println("Existing details ", *p)
	p.address = s
}
func exercise02() {
	fmt.Println("******* Running exercise 02")
	p := person2{
		first:   "James ",
		address: "Washington DC",
	}
	fmt.Println("Old address ", p.address)
	changeMe(&p, "San Jose")
	fmt.Println("New address ", p.address)
}
