package main

import (
	"fmt"
)

// Declaration of float
var (
	f1 float32
	f2 float64
)

func main() {
	var maxFloat32 float32
	maxFloat32 = 16777216
	fmt.Println(maxFloat32 == maxFloat32+10)

	fmt.Println(f1, f2)
	fmt.Println(maxFloat32 + 2000)
	fmt.Println(maxFloat32 + 10000)
}
