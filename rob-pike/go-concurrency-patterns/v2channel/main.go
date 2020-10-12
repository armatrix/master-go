// 该例子中在新开的一个goroutine中执行一个任务，通过一个无缓冲的channel来在main函数和新开的go程之间通信
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go boring("I am boring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q.\n", <-c)
	}
	fmt.Println("You're boring, I'm leaving!")
}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Second)
	}
}
