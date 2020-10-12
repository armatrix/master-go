// Receive on quite channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// How do we know it's finished? Wait for it to tell us it's done: receive on the quit channel.
func main() {
	quit := make(chan string)

	c := boring("Joe", quit)

	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye!"
	fmt.Printf("Joe says:%q\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s,%d", msg, i):
			case quitMsg := <-quit:
				//	Think about we need to do sth. before return,
				//	like remove temp files, some cleaning stuff,
				//	we have to make sure all those things been done before return.
				//	Here's the solution.
				cleanup()
				quit <- quitMsg + " See you!"
				return
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func cleanup() {
	fmt.Println("clean ops...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done.")
}
