package main

import "fmt"

func main() {
	fmt.Println("Control flow in GO")
	for i := 0; i < 5; i++ {
		fmt.Printf("The outer loop :%d \n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t inner loop is : %d \n", j)
		}
	}

	x := 101
	for x < 205 {
		fmt.Printf("Value of variable is %d, in binary %b\n", x, x)
		x = x << 1
	}
	//never ending loop
	x = 7
	for {
		if x > 9 {
			break
		}
		fmt.Println("value of x is ", x)
		x++
	}
	printEvenOnly()
	printascii()
	conditionalExample()
	bringItAll()
	switchStatement()
	switchOnString()
	exercise()
}

func printEvenOnly() {
	x := 0
	for {
		x++
		if x >= 4 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Printf("Decimal value %d binary value %b \n", x, x)
	}
}

func printascii() {
	for i := 33; i < 122; i++ {
		fmt.Printf("%v \t %#X \t %#U\n", i, i, i)
	}
}

func conditionalExample() {
	x := 42
	if x == 40 {
		fmt.Println("Value of x is 40")
	} else if x == 41 {
		fmt.Println("Value of x is 41")
	} else {
		fmt.Println("Value of x is not 40")
	}
}

func bringItAll() {
	// loop madulous condition
	for i := 0; i%10 == 0; i++ {
		fmt.Println("Found MOD of 10")
		if i > 10 {
			fmt.Println("Looped more then 10 time breaking")
			break
		}
	}
}

//Fallthrough will happen even when the case evaluates to false.
func switchStatement() {
	switch {
	case false:
		fmt.Println("This will not print")
	case 4%2 == 0:
		fmt.Println("This will print 4%2==0")
		fallthrough
	case 4 == 5:
		fmt.Println("This will print 4==5 ")
		fallthrough
	case 2 == 2:
		fmt.Println("This will print 2==2 ")
		fallthrough
	default:
		fmt.Println("Printing default")
	}
}

func switchOnString() {
	n := "Bond"
	switch n {
	case "Moneypenney", "Bond":
		fmt.Println("Miss money for Bond")
	case "M":
		fmt.Println("Mr M")
		fallthrough
	default:
		fmt.Println("Miss default")
	}
}
