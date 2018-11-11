package main

import "fmt"

func main() {
	// The race condition occurs because it is not known whether x will have the value 0, 1, or 2.
	// It depends on the order in which the goroutines are executed, and this order is not known.
	var x int
	go func() {
		x++
	}()

	go func() {
		x++
	}()

	fmt.Println(x)
}
