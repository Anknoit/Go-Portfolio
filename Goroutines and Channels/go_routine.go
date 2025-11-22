// implementing a goroutine
package main

import (
	"fmt"
	"sync"
	"time"
)

func go_routine(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	go_routine("some random", &wg) // A

	wg.Add(1)
	go go_routine(("Asynchronous call"), &wg) // B - will not wait for completion to move to next line as it starts running in background
	go_routine("Synchronus call", &wg)        //C // cannot call sync func first as it will trigger first over the loop and completes the loop without even calling the async func as the condition got fullfilled with the first sync function
	// time.Sleep(5000 * time.Millisecond)       // 5 second hold To ensure there isnt any go_routine task left before main closes
	wg.Wait()
}

//time.Sleep(500 * time.Millisecond) gives the scheduler a natural point to switch between goroutines. Because both goroutines spend most of their time sleeping, the scheduler alternates them, producing the alternating lines you observed.
