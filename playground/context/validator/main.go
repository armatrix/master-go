package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 4 * time.Second

func main() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()
	startAt := time.Now()
	if validateOps1(ctx) || validateOps2(ctx) || validateOps3(ctx) {
		fmt.Println("some condition has been occured")
	}
	elapsed := time.Since(startAt)
	fmt.Printf("|| ops cost: %v\n", elapsed)

}

func validateOps1(ctx context.Context) bool {
	time.Sleep(3 * time.Second)
	fmt.Println("exec validateOps1")
	return true
}

func validateOps2(ctx context.Context) bool {
	time.Sleep(3 * time.Second)
	fmt.Println("exec validateOps2")
	return true
}

func validateOps3(ctx context.Context) bool {
	time.Sleep(3 * time.Second)
	fmt.Println("exec validateOps3")
	return false
}
