package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//go run -race *.go
func raceConditionDemoUsingMutex() {
	t1 := time.Now().UnixNano()
	fmt.Println("************ Running race Condition Mutex Demo **********")
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines initial", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			//fmt.Println("Read value ", counter)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			//fmt.Println("Updated counter value ", counter)
			wg.Done()
		}()
		fmt.Println("Goroutines count during execution", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines count after execution", runtime.NumGoroutine())
	fmt.Println("Total execution time ", (time.Now().UnixNano() - t1))
	fmt.Println("Count ", counter)
}

func raceConditionDemoUsingAtomic() {
	t1 := time.Now().UnixNano()
	fmt.Println("************ Running race Condition Atomic Demo **********")
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines initial", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			//fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
		//fmt.Println("Goroutines count during execution", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines count after execution", runtime.NumGoroutine())
	fmt.Println("Total execution time ", (time.Now().UnixNano() - t1))
	fmt.Println("Count ", counter)
}
