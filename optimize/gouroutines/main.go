package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//begin wait
	const numGoroutines = 5
	wg := sync.WaitGroup{}
	wg.Add(numGoroutines)
	// run go func
	// wg.Done()
	//end wait

	//begin concurrent
	for i := 0; i < numGoroutines; i++ {
		i := i //shadow 
		go func(num int) {
			defer wg.Done()
			sleepDuration := 1 + rand.Intn(3) // Sleep for 5 to 10 seconds
			time.Sleep(time.Duration(sleepDuration) * time.Second)
			fmt.Printf("Goroutine %d: Loop number %d\n", num, i)
		}(i)
	}
	//end concurrent

	//begin wait
	wg.Wait()
	//end wait
	fmt.Println("All goroutines completed")
}
