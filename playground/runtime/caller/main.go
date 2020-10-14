package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	f2()
}

func f1() {
	pc, file, line, ok := runtime.Caller(2)
	//pc, file, line, ok := runtime.Caller(0)
	//pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("failed")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	s, i := runtime.FuncForPC(pc).FileLine(pc)
	fmt.Printf("s: %s; i: %d\n", s, i)
	fmt.Printf("func name: %s\n", funcName)
	fmt.Printf("pc: %T\n", pc)
	fmt.Printf("file: %v\n", file)
	fmt.Printf("base path: %v\n", path.Base(file))
	fmt.Printf("line: %d\n", line)
}
func f2() {
	f1()
}
