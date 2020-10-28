package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func exampleUsingWaitGroup() {
	wg.Add(2)
	fName := "names_wg.txt"
	go exampleWG(fName)
	go exampleReaderWG(fName)
	wg.Wait()
}

func exampleWG(fName string) {
	fmt.Println("Creating file exampleWG()")
	f, err := os.Create(fName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("string(rand.Intn(WG: taste of india...))")
	io.Copy(f, r)
	fmt.Println("Added data to file")
	wg.Done()
}

func exampleReaderWG(fName string) {
	fmt.Println("Called exampleReaderWG(), waiting to read")

	fmt.Println("Reading data now")
	f, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	wg.Done()
}
