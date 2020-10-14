package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%b\n", 52) // 110100
	fmt.Printf("%b\n", math.Float32bits(0.1))
}
