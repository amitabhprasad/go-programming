package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func exercise() {
	fmt.Println("*********** Running exercises ***********")
	exercise01()
	exercise02()
	exercise03()
	exercise04()
	exercise05()
}

//https://mholt.github.io/json-to-go/
type user struct {
	First  string   `json:"first"`
	Last   string   `json:"last"`
	Age    int      `json:"age"`
	Saying []string `json:"saying"`
}

func exercise01() {
	fmt.Println("########## Executing Exercise 01 ###############")
	u1 := user{
		First: "James",
		Age:   52,
	}
	u2 := user{
		First: "Anderson",
		Age:   53,
	}
	u3 := user{
		First: "Angela",
		Age:   48,
	}

	users := []user{u1, u2, u3}
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error during marshal ", err)
	}
	fmt.Println("Marshaled data ", string(bs))
}

func exercise02() {
	fmt.Println("########## Executing Exercise 02 Unmarshal ###############")
	jsonString := `[{"first":"James","age":52},{"first":"Anderson","age":53},{"first":"Angela","age":48}]`
	var people []user
	err := json.Unmarshal([]byte(jsonString), &people)
	if err != nil {
		fmt.Println("Error during Unmarshal ", err)
	}
	fmt.Println("Unmarshaled data ", people)
	for i, j := range people {
		fmt.Println("Person # ", i)
		fmt.Printf("\t Details of User %v with age %v\n", j.First, j.Age)
	}
}

func exercise03() {
	fmt.Println("########## Executing Exercise 03 ###############")
	jsonEncoder := json.NewEncoder(os.Stdout)
	u1 := user{
		First: "James",
		Age:   52,
	}
	u2 := user{
		First: "Anderson",
		Age:   53,
	}
	u3 := user{
		First: "Angela",
		Age:   48,
	}
	people := []user{u1, u2, u3}
	fmt.Println("Data to encode ", people)
	fmt.Println("Starting encoding straight to stdout")
	err := jsonEncoder.Encode(people)
	if err != nil {
		fmt.Println("Error duing Encode flow ")
	}
}

func exercise04() {
	fmt.Println("########## Executing Exercise 04 ###############")
	xi := []int{101, 123, 99, 78, 565, 43, 2, 45, 102, 2, 1, 10, 20}
	var xs []string
	xs = []string{"random", "rainbow", "snow", "wind", "cloud", "rain", "cold"}
	fmt.Println("Unsorted xi ", xi)
	fmt.Println("Unsorted xs ", xs)
	sort.Ints(xi)
	fmt.Println("Sorted data xi ", xi)
	sort.Strings(xs)
	fmt.Println("Sorted data xs ", xs)
}

type peopleByAge []user

func (p peopleByAge) Len() int           { return len(p) }
func (p peopleByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p peopleByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type byLast []user

func (p byLast) Len() int           { return len(p) }
func (p byLast) Less(i, j int) bool { return p[i].Last < p[j].Last }
func (p byLast) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func exercise05() {
	fmt.Println("########## Executing Exercise 05 ###############")
	u1 := user{
		First:  "James",
		Last:   "Bond",
		Age:    52,
		Saying: []string{"secret agent", "spy", "mission", "USSR"},
	}
	u2 := user{
		First:  "Anderson",
		Last:   "Andrey",
		Age:    53,
		Saying: []string{"Project", "Manager", "Canada", "Zebra"},
	}
	u3 := user{
		First:  "Angela",
		Last:   "Markel",
		Age:    48,
		Saying: []string{"Prime minister", "Eagle", "Amazing"},
	}
	fmt.Println("Sorting people by age")
	var p peopleByAge = []user{u1, u2, u3}
	sort.Sort(p)
	printUserDetails(p)
	fmt.Println("Sorting people by Last Name")
	var p2 byLast = []user{u1, u2, u3}
	sort.Sort(p2)
	printUserDetails(p2)
}

func printUserDetails(p []user) {
	for i, j := range p {
		fmt.Printf("Getting details  #: %v\n", i)
		fmt.Printf("%v %v : %v \n", j.First, j.Last, j.Age)
		sort.Strings(j.Saying)
		for m, n := range j.Saying {
			fmt.Printf("Fact  %v \t %v \n", m, n)
		}
	}
}
