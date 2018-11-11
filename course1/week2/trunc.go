package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	fmt.Println("Enter 'exit' to quit.")
	fmt.Println("Please enter a floating point number you want to be truncated: ")
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			break
		}
		input = scanner.Text()

		floatInput, _ := strconv.ParseFloat(input, 64)
		fmt.Printf("Truncated from %s to %d\n", input, int(floatInput))
	}
}
