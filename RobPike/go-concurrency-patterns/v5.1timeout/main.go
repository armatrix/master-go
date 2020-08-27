// Timeout for whole conversation using select
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create the timer once, outside the loop, to time out the entire conversation.
// In the previous program, we had a timeout for each message.
func main() {
	c := boring("Gina")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s,%d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
