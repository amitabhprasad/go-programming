package main

import (
	"fmt"
	"time"
)

func example() {
	//exampleBufferedChannle()
	//channelType()
	//usingChannels()
	//channelWithRange()
	channelWithSelect()
}

func channelWithSelect() {
	fmt.Println("Executing channel with select")
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)
	receive(eve, odd, quit)
	fmt.Println("Exiting done")
}
func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Reading data from even channel: ", v)
		case v := <-o:
			fmt.Println("Reading data from odd channel: ", v)
		case i, ok := <-q:
			// when channel is closed `ok` will be set to false
			if !ok {
				fmt.Println("From , ok  as false ", i, ok)
				return
			}
			fmt.Println("From , ok as true ", i, ok)
		}
	}
}
func send(eve, odd, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
	close(quit)
}
func channelWithRange() {
	fmt.Println("Executing channel with range")
	c := make(chan int)
	go addToChannel(&c)
	// read data from channle until it's closed
	for v := range c {
		fmt.Println("Reading ", v)
	}
}
func addToChannel(c *chan int) {
	for i := 0; i < 10; i++ {
		*c <- i
	}
	close(*c)
}
func usingChannels() {
	c := make(chan int)
	//send
	go foo(c)
	//receive
	bar(c)
	fmt.Println("About to exit")
}

func foo(c chan<- int) {
	fmt.Println("Adding data to channel")
	time.Sleep(time.Second * 5)
	c <- 42
}
func bar(c <-chan int) {
	fmt.Println("waiting.... to add data in the channel")
	fmt.Println("Reading data from channel ", <-c)
}
func exampleBufferedChannle() {
	fmt.Println("Executing buffered channel example")
	c := make(chan int, 2)
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("------")
	fmt.Printf("%T\n", c)
}
func channelType() {
	fmt.Println("Exploring channel types")
	c := make(chan int)
	//<-chan<-
	cr := make(<-chan int)
	cw := make(chan<- int)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cw)
	go func() {
		cw <- 101
	}()
	//fmt.Println("Reding in receive only ", <-cw)
}
