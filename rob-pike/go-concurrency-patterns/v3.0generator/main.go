// Generator
// 在go中 function和channel都可以作为一类值（first class value），该例中通过boring函数来返回一个channel，在boring函数中新开goroutine来做业务的处理，
// 注意其中关闭channel的时间，这里可以作为一个基本的范式，将任务丢在goroutine中执行，返回channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("boring!") // Function returning a channel.
	for i := 0; i < 50; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring, I'm leaving!")
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings.
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s,%d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) + time.Millisecond)
		}
		close(c)
	}()
	return c // Return the channel to the caller.
}
