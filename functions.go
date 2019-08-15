package main

import (
	"fmt"
	"math/cmplx"
)

var (
	flag, is_false, is_true bool = true, false, true
	z complex128 = -1 + 2i
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 41 / 99
	y = sum - x
	return 
}

func flags() (bool, bool, bool, bool) {
	is_no := false
	return flag, is_false, is_true, is_no
}

func getComplexAbs(z complex128) (x float64) {
	x = cmplx.Abs(z)
	return x
}

func Sqrt(x float64) float64 {
	// Newton's Method
	z := x/2
	var prev float64
	for z - prev != 0 {
		prev = z
		z -= (z*z - x) / (2*z)
	}
	fmt.Println(z)
	return z
}

func testDefer() {
	defer Sqrt(100)
	defer Sqrt(64)
	defer Sqrt(25)
	Sqrt(9)
}

