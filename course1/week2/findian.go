package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string

	fmt.Println("Enter a string, type 'exit' to quit: ")
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			break
		}

		input = strings.ToLower(scanner.Text())

		if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not found! Try again: ")
			continue
		}
	}
}
