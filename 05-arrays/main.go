package main

import "fmt"

func main() {
	exercise01()
	exercise02()
	exercise03()
	exercise04()
	exercise05()
	exercise06()
	exercise07()
	exercise08()
	exercise09()
}

func exercise01() {
	fmt.Println("****** running exercise 01 ****")
	x := [5]int{2, 4, 5, 7, 9}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func exercise02() {
	fmt.Println("****** running exercise 02 ****")
	x := []int{2, 4, 5, 7, 9}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func exercise03() {
	fmt.Println("****** running exercise 03 ****")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}

func exercise04() {
	fmt.Println("****** running exercise 04 ****")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
}

func exercise05() {
	fmt.Println("****** running exercise 05 ****")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}

func exercise06() {
	fmt.Println("****** running exercise 06 ****")
	x := make([]string, 50, 50)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	s := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	x = append(x, s...)
	fmt.Println("new length ", len(x))
	fmt.Println("new capacity ", cap(x))
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
func exercise07() {
	fmt.Println("****** running exercise 07 ****")
	x1 := []string{"James", "Bond", "Shaken"}
	x2 := []string{"Miss", "Money", "Penney"}
	x := [][]string{{"x1"}, {"x2"}}
	xs := [][]string{x1, x2}
	fmt.Println(x)
	fmt.Println(xs)
	for i, v := range xs {
		fmt.Println("Iterating over ", i)
		for j, v2 := range v {
			fmt.Printf("\t index position: %v\t value \t%v\n", j, v2)
		}

	}
}

func exercise08() {
	fmt.Println("****** running exercise 08 ****")
	m := map[string][]string{
		"amitabh": {"running", "coding", "politics"},
		"james":   {"movies", "booze", "drugs"},
	}
	m["bond"] = []string{"movies", "tv", "serials"}

	for k, v := range m {
		fmt.Println("Reading data for ", k)
		for i, j := range v {
			fmt.Printf("\t\t %v \t %v\n", i, j)
		}
	}
}

func exercise09() {
	fmt.Println("****** running exercise 09 for delete ****")
	m := map[string][]string{
		"amitabh": {"running", "coding", "politics"},
		"james":   {"movies", "booze", "drugs"},
	}
	m["bond"] = []string{"movies", "tv", "serials"}
	if _, ok := m["james"]; ok {
		delete(m, "james")
	}
	for k, v := range m {
		fmt.Println("Reading data for ", k)
		for i, j := range v {
			fmt.Printf("\t\t %v \t %v\n", i, j)
		}
	}
}
