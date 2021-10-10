package main

import (
	"fmt"
)

var (
	isTrue  bool
	isFalse bool
)

func main() {
	if isTrue && isFalse {
		fmt.Println("Both conditions needs to be true")
	}

	if isTrue || isFalse {
		fmt.Println("Only one condition needs to be true")
	}
}
