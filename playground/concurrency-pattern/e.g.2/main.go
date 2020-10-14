package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	jobsCount := 10
	group := sync.WaitGroup{}
	jobsChan := make(chan int, 3)
	// a) 生成指定数目的goroutine，每个goroutine消费jobsChan中的数据
	poolCount := 3
	for i := 0; i < poolCount; i++ {
		go func() {
			for j := range jobsChan {
				time.Sleep(time.Second * 3)
				fmt.Printf("hello %d\n", j)
				group.Done()
			}
		}()
	}
	// b) 把job依次推送到jobsChan供goroutine消费
	for i := 0; i < jobsCount; i++ {
		jobsChan <- i
		group.Add(1)
		fmt.Printf("index: %d,goroutine Num: %d\n", i, runtime.NumGoroutine())
	}
	group.Wait()
	fmt.Println("done!")
}
