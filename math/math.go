package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Happy", math.Pi, "Day")
	fmt.Println(add(42, 13))
}
