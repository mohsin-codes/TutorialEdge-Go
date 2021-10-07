// different types of integers

// 1. uint8
// 2. int8
// 3. uint16
// 4. int16
// 5. uint32
// 6. int32
// 7. uint64
// 8. int64

// Overflow program

package main

func main() {
	print("Hello, World!")

	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	print(myint)
}
