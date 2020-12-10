package go_koans

import "fmt"

func divideFourBy(i int) int {
	return 4 / i
}

const DIVISOR = 0

func aboutPanics() {
	var n int

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from divideFourBy")
			n = 2
		}
	}()

	n = divideFourBy(DIVISOR)

	assert(n == 2) // panics are exceptional errors at runtime
}