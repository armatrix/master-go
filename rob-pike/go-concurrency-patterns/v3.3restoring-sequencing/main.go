// Restoring sequencing

// Send a channel on a channel, making goroutine wait its return.
// Receive all messages, then enable them again by sending on a private channel.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Message contains a channel for the reply.
type Message struct {
	data string
	wait chan bool
}

func main() {
	waitForIt := make(chan bool) // Shared between all messages.
	c := fanIn(boring(Message{"aaron", waitForIt}), boring(Message{"gina", waitForIt}))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.data)

		msg2 := <-c
		fmt.Println(msg2.data)

		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring, I'm leaving!")
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
			//time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg Message) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool) //Shared between all messages.
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s,%d", msg.data, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) + time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}
