package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func exercise() {
	//exercise01()
	//exercise02()
	//exercise03()
	//exercise04()
	//exercise05()
	//exercise06()
	exercise07()
}

func exercise01() {
	fmt.Println("******* Running exercise 01 channel basics")
	c := make(chan int)
	go func() { c <- 42 }()
	fmt.Println(<-c)
}

func exercise02() {
	fmt.Println("******* Running exercise 02 channel basics")
	cs := make(chan int)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("-----------\n")
	fmt.Printf("cs \t %T\n", cs)
}

func exercise03() {
	fmt.Println("******* Running exercise 03")
	c := gen()
	receiveChan(c)
	fmt.Println("About to Exit")
}
func receiveChan(c <-chan int) {
	for v := range c {
		fmt.Println("Reading data ", v)
	}
}
func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("About to sleep Zzzz")
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		//close(c)
	}()
	return c
}
func exercise04() {
	fmt.Println("******* Running exercise 04 using select statment with channels")
	q := make(chan int)
	c := generate(q)
	receive4(c, q)
	fmt.Println("About to exit")
}
func generate(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 1; i < 11; i++ {
			fmt.Println("Generated data taking nap Zzzz")
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			c <- i * (rand.Intn(100))
		}
		close(q)
		close(c)
		//q <- 1
	}()
	return c
}
func receive4(c <-chan int, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Reading data ", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("Done reding data all channel closed returning", v, ok)
			}
			return
		}
	}
}

func exercise05() {
	fmt.Println("******* Running exercise 05 comma ok idiom *****")
	c := make(chan int)
	go func() { c <- 42 }()
	v, ok := <-c
	fmt.Println(v, ok)
	close(c)
	v, ok = <-c
	fmt.Println(v, ok)
}

func exercise06() {
	fmt.Println("******* Running exercise 06 ")
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
}

func exercise07() {
	fmt.Println("******* Running exercise 07 channel throttle")
	c := make(chan string)
	go throttleChannels(c)
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("Done !!!")
}

func throttleChannels(c chan<- string) {
	fmt.Println("--->>>> Entered function throttleChannels")
	var wg sync.WaitGroup
	const MaxGoRoutines = 10
	wg.Add(MaxGoRoutines)

	for i := 0; i < MaxGoRoutines; i++ {
		go func(i int) {
			for j := 0; j < 10; j++ {
				c <- fmt.Sprintf("Go-Routine: %v %v", i, j)
			}
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			wg.Done()
		}(i)
		fmt.Println(" Number of subroutines ", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("***** Done processing data closing channel now ****")
	close(c)
	fmt.Println("Exiting function throttleChannels <<<<---")
}
