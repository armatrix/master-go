// Timeout using select 单个channel的超时我们直接返回
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// The time.After function returns a channel that blocks for the specified duration.
// After the interval, the channel delivers the current time, once.
func main() {
	c := boring("Aaron")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(900 * time.Millisecond): // P = 1/10
			fmt.Println("You are too slow.")
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
