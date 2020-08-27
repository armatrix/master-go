// Quit channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// We can turn this around and tell Aaron to stop when we're tired of listening to him.
func main() {
	quit := make(chan bool)

	c := boring("Aaron", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
}

func boring(msg string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s, I am tired of listening to you for %d times.", msg, i):
			case <-quit:
				//	Think about we need to do sth. before return,
				//	like remove temp files, some cleaning stuff,
				//	we have to make sure all those things been done before return.
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
