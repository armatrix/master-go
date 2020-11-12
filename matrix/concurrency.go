package matrix

import (
	"reflect"
	"sync"
)

// Concurrent reduce item in slice
func Concurrent(concurrency int, slice interface{}, reduceFunc interface{}) {
	throttle := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	sliValue := reflect.ValueOf(slice)
	length := sliValue.Len()

	rf := reflect.ValueOf(reduceFunc)

	wg.Add(length)
	for i := 0; i < length; i++ {
		throttle <- struct{}{}

		go func(i int) {
			defer wg.Done()
			defer func() {
				<-throttle
			}()
			rf.Call([]reflect.Value{sliValue.Index(i)})
		}(i)
	}

	wg.Wait()
}
