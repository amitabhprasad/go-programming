package main

import "fmt"

func exercise() {
	exercise01()
	exercise02()
	exercise03()
	exercise04()
	exercise05()
	exercise06()
	exercise07()
	exercise08()
}

func exercise01() {
	fmt.Println("****** running exercise 01 **** printing UPPER CASE 3 times")
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t%#U \n", i)
		}
	}

}

func exercise02() {
	fmt.Println("****** running exercise 02 for cond {} ****")
	bd := 1985
	for bd <= 2020 {
		fmt.Println(bd)
		bd++
	}
}

func exercise03() {
	fmt.Println("****** running exercise 03 for {} ****")
	bd := 1985
	for {
		fmt.Println(bd)
		if bd == 2020 {
			break
		}
		bd++
	}
}

func exercise04() {
	fmt.Println("****** running exercise 04 ****")
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4 remainder is %v\n", i, i%4)
	}
}

func exercise05() {
	fmt.Println("****** running exercise 05 ****")
	x := 42
	if x < 42 {
		fmt.Println("Values is less")
	} else if x > 42 {
		fmt.Println("Values is more")
	} else {
		fmt.Println("Values is equal")
	}
}
func exercise06() {
	fmt.Println("****** running exercise 06 ****")
	switch {
	case true:
		fmt.Println("Should print")
	case false:
		fmt.Println("Shouldn't print")
	default:
		fmt.Println("Inside switch with default")
	}
}
func exercise07() {
	fmt.Println("****** running exercise 07 ****")
	favSport := "Ice Hockey"
	switch favSport {
	case "Hockey":
		fmt.Println("Indias national sports")
	case "Cricket":
		fmt.Println("Sports watched by billions of indian")
	default:
		fmt.Printf("Data for this sports %v not yet known\n", favSport)
	}
}

func exercise08() {
	fmt.Println("****** running exercise 08 ****")
}
