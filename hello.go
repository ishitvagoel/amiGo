package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello World !", rand.Intn(10))
	fmt.Println(add(2,3))
	fmt.Println(swap("1", "2"))
	fmt.Println(split(10))
	fmt.Println(flags())
	fmt.Printf("%T, %.5f", z, getComplexAbs(z))
	fmt.Println()
	fmt.Println(Sqrt(10000))
	testDefer()
}
