// implementing a goroutine
package main

import (
	"fmt"
	"time"
)

func go_routine(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	go go_routine(("Asynchronous call"))
	go_routine("Synchronus call") // cannot call sync func first as it will trigger first over the loop and completes the loop without even calling the async func as the condition got fullfilled with the first sync function
}
