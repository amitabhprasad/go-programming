package main

import "fmt"

func exercise() {
	exercise01()
	exercise03()
	exercise04()
}

type person1 struct {
	first   string
	last    string
	flavors []string
}

func exercise01() {
	fmt.Println("****** running exercise 01 & 02 ****")
	p1 := person1{
		first: "James",
		last:  "Bond",
		flavors: []string{
			"chocolate",
			"vanilla",
		},
	}

	p2 := person1{
		first: "Sandra",
		last:  "Bullock",
		flavors: []string{
			"dark chocolate",
			"sitafal",
		},
	}

	fmt.Println("Reading details of p1 ", p1.first)
	for i, v := range p1.flavors {
		fmt.Println(i, v)
	}
	fmt.Println("Reading details of p2 ", p2.first)
	for i, v := range p2.flavors {
		fmt.Println(i, v)
	}

	m := map[string]person1{
		p1.last: p1,
		p2.last: p2,
	}
	for k, v := range m {
		fmt.Println("Reading details of ", v.first, "with key ", k)
		for i := 0; i < len(v.flavors); i++ {
			fmt.Println(v.flavors[i])
		}
	}
}

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func exercise03() {
	fmt.Println("****** running exercise 03 ****")
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: false,
	}

	fmt.Printf("%T, %v\n", t, t)
	fmt.Printf("%T, %v\n", s, s)

	fmt.Printf("%v, %v\n", t.vehicle, t.fourWheel)
	fmt.Printf("%v, %v\n", s.vehicle, s.luxury)
}

func exercise04() {
	fmt.Println("****** running exercise 04 ****")
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"money":  1234,
			"penney": 007,
		},
		favDrinks: []string{"martini", "scotch"},
	}
	fmt.Println(s)
	for i, v := range s.favDrinks {
		fmt.Println(i, v)
	}
}
