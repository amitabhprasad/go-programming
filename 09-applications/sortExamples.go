package main

import (
	"fmt"
	"sort"
)

type person2 struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

type sortPerson2 []person2

func (p sortPerson2) Len() int {
	return len(p)
}
func (p sortPerson2) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}
func (p sortPerson2) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p person2) string() string {
	return fmt.Sprintf("%s:%d", p.First, p.Age)
}

func sortExample() {
	x := []int{2, 23, 45, 1, 0, 12, 5, 7, 10}
	sort.Ints(x)
	fmt.Println(x)
}

func sortCustomExample() {
	p1 := person2{First: "Brad", Age: 32}
	p2 := person2{First: "Brad Pitt", Age: 52}
	p3 := person2{First: "Michael", Age: 66}
	p4 := person2{First: "Angelina", Age: 43}
	p5 := person2{First: "Deepika", Age: 35}

	p := []person2{p1, p2, p3, p4, p5}
	//var s sortPerson2 = p
	var pa2 []person2 = []person2{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}
	pa := []person2{
		{"Brad", 32},
		{"John", 42},
	}
	fmt.Printf("%T %T\n", pa, pa2)
	fmt.Println(p)
	sort.Sort(sortPerson2(p))
	fmt.Println(p)

	// reverse sorting on the fly
	sort.Slice(pa2, func(i, j int) bool {
		return pa2[i].Age > pa2[j].Age
	})
	//
	fmt.Println("Reverse sorted by age", pa2)
	//sort by age
	sort.Slice(pa2, func(i, j int) bool {
		return pa2[i].First < pa2[j].First
	})
	fmt.Println("Sorted by Name", pa2)
}
