// 示例展示了goroutine的基本用法，通过go关键字新开一个goroutine，你可以尝试分别注释掉两段sleep函数，观察现象
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("Aaron")
	fmt.Println("I am listening...")
	time.Sleep(2 * time.Second) // 当注释掉这里的时候，我们会发现boring(msg string)会出现没来得及执行main函数直接退出的情况
	fmt.Println("You're boring, I'm leaving!")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i, time.Now().Format("15:04:05"), "is boring!") // 当注释掉这里的时候，我们会发现在14行输出"You're boring, I'm leaving!"后仍继续输出两boring函数中的内容
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
