// 这里我们要模拟的不是让每个任务单独的在goroutine中运行，而是通过控制goroutine的数量，创建m个任务，在n个goroutine中执行
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	startTime := time.Now()

	// 计算第n个fib数
	fibIndex := 45
	jobs := make(chan int, fibIndex)
	results := make(chan int, fibIndex)

	// worker 数量
	workerTolly := 1

	for i := 0; i < workerTolly; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < fibIndex; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < fibIndex; j++ {
		fmt.Printf("index: %d,goroutine Num: %d\n", j, runtime.NumGoroutine())
		fmt.Println(<-results)
	}
	elapsed := time.Since(startTime)
	fmt.Println(elapsed)
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
