package main

import "fmt"

func main() {
	// go func() {
	// 	fmt.Println("hello")
	// }()
	// time.Sleep(100 * time.Millisecond) // this puts the main goroutine to sleep for a little while, which gives the other goroutine time to execute
	// fmt.Println("xd")

	c := make(chan int)
	go prod(1, 2, c)
	go prod(3, 4, c)

	a := <-c
	b := <-c
	fmt.Println(a * b)
}

// prod takes two integers and a channel and passes the product of the two integers into the channel
func prod(x, y int, c chan int) {
	c <- x * y
}
