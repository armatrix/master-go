package main

import (
	"reflect"
	"sync"
	"time"
)

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	par(5, arr, job)
}

func job(i int) {
	time.Sleep(time.Second * 2)
	println(i)
}

func par(concurrency int, arr interface{}, f interface{}) {
	throttle := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	varr := reflect.ValueOf(arr)
	l := varr.Len()

	rf := reflect.ValueOf(f)

	wg.Add(l)
	for i := 0; i < l; i++ {
		throttle <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() {
				<-throttle
			}()
			rf.Call([]reflect.Value{varr.Index(i)})
		}(i)
	}

	wg.Wait()
}
