package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func examples() {
	//example1()
	c := make(chan int)
	go example2(c)
	exampleReader(c)
	i, ok := <-c
	fmt.Println("status of channle before exit ", i, ok)
}

func example1() {
	var answer1, answer2, answer3 string
	fmt.Println("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sports: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)
}

func example2(c chan<- int) {
	fmt.Println("Creating file")
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("string(rand.Intn(taste of india...))")
	io.Copy(f, r)
	fmt.Println("Added data to file")
	c <- 101
	close(c)
}

func exampleReader(c <-chan int) {
	fmt.Println("Called exampleReader(), waiting for chanel to be ready")
	i, ok := <-c
	if ok {
		fmt.Println("Reading data now")
		f, err := os.Open("names.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bs))
	}
	fmt.Println(i, ok)
}
