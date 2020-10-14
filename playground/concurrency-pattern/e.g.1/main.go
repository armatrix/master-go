package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("runtime.NumCPU(): ", runtime.NumCPU())
	jobsCount := 10
	group := sync.WaitGroup{}
	for i := 0; i < jobsCount; i++ {
		group.Add(1)
		go func(i int) {
			time.Sleep(time.Second * 3) //模拟任务执行的时间
			fmt.Printf("hello %d!\n", i)
			group.Done()
		}(i)
		fmt.Printf("index: %d,goroutine Num: %d.\n", i, runtime.NumGoroutine())
	}
	group.Wait()
	fmt.Println("done!")
}
