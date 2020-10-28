package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/**
* Show the usage of json binding and unbinding
**/

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"Age"`
}

func jsonExample() {
	fmt.Println("********** Running json examples ********")
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   52,
	}
	p2 := person{
		First: "Money",
		Last:  "Penney",
		Age:   42,
	}

	people := []person{p1, p2}
	md := marshalData(people)
	fmt.Printf("Marshaled data type %T and data %v \n", md, md)
	var ud []person
	unmarshalData(md, &ud)
	fmt.Printf("UnMarshaled data type %T and data %v \n", ud, ud)
}

func marshalData(d interface{}) string {
	bs, err := json.Marshal(d)
	if err != nil {
		fmt.Println("Error during marshaling")
	}
	return string(bs)
}

func unmarshalData(md string, v *[]person) {
	bs := []byte(md)
	err := json.Unmarshal(bs, v)
	if err != nil {
		fmt.Println("Error during un-mashal")
	}
}

func dataEncoder(d interface{}) {
	stdEncoder := json.NewEncoder(os.Stdout)
	stdEncoder.Encode(d)
}
