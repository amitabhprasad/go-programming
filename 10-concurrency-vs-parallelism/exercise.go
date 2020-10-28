package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func exercise() {
	//exercise01()
	exercise01_1()
	exercise02()
	exercise03()
	exercise05()
	exercise06()
	fmt.Println("*********** exiting ******")
}

var wg1 sync.WaitGroup

func exercise01() {
	fmt.Println("******* Running exercise 01 creates 3 go routines with wt groups")
	fmt.Println("Go Routine Count (Start):\t", runtime.NumGoroutine())

	wg1.Add(2)
	go foo1()
	go bar1()
	fmt.Println("Go Routine Count (during execution):\t", runtime.NumGoroutine())
	wg1.Wait()
	fmt.Println("Go Routine Count (End):\t", runtime.NumGoroutine())
}
func foo1() {
	fmt.Println("Executing foo() ")
	wg1.Done()
}

func bar1() {
	fmt.Println("Executing bar() ")
	wg1.Done()
}

func exercise01_1() {
	fmt.Println("******* Running exercise 01_1 creates 3 go routines with wt groups using anynomous func")
	fmt.Println("Go Routine Count (Start):\t", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Executing 1st function....Go routine count ", runtime.NumGoroutine())
		wg.Done()
	}()

	go func() {
		fmt.Println("Executing 2nd function.... Go routine count ", runtime.NumGoroutine())
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Go Routine Count (At Exit):\t", runtime.NumGoroutine())
}

//******************************* 2 *************************************
type person struct {
	First    string   `json:"first"`
	Age      int64    `json:"age"`
	Language []string `json:"language"`
}

func (p *person) speak() {
	fmt.Printf("I, %v can speak following language : \n", p.First)
	for i, j := range p.Language {
		fmt.Println(i, j)
	}
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func exercise02() {
	fmt.Println("******* Running exercise 2 exploring method set using interface *******")
	p := person{
		First:    "James",
		Age:      53,
		Language: []string{"English", "Hindi"},
	}
	saySomething(&p)
}

//******************************* 3 *************************************
func exercise03() {
	fmt.Println("******* Running exercise 03 Example of race condition and Mutex")
	var counter int64
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			runtime.Gosched()
			counter = v
			fmt.Println("Number of goroutines currently executing ", runtime.NumGoroutine(), counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Number of go routines at the end %v \t counter val %v \n", runtime.NumGoroutine(), counter)
}

//******************************* 5 *************************************

func exercise05() {
	fmt.Println("******* Running exercise 05 - Fix race caondition by using Atomic")
	var counter int64
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			//time.Sleep(time.Second)
			runtime.Gosched()
			fmt.Println("Number of goroutines currently executing ", runtime.NumGoroutine(), atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("Number of go routines at the end %v \t counter val %v \n", runtime.NumGoroutine(), counter)
}

func exercise06() {
	fmt.Println("******* Running exercise 06 *****")
	fmt.Println("CPUs:\t", runtime.NumCPU())
	fmt.Println("ARCH:\t", runtime.GOARCH)
}
