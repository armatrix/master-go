package main

import (
	"fmt"
	"runtime"
	"time"
)

func printStats(point string, mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println(point+" mem.Alloc:", mem.Alloc)
	fmt.Println(point+" mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println(point+" mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println(point+" mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

// GODEBUG=gctrace=1 go run stat.go
func main() {
	var mem runtime.MemStats
	printStats("start", mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	printStats("after first make slice", mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
		time.Sleep(5 * time.Second)
	}
	printStats("2nd make slice", mem)
}
