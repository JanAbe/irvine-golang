package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input int64
	myslice := make([]int64, 0, 3)

	fmt.Println("Enter a number, type 'X' to quit: ")
	for scanner.Scan() {
		if scanner.Text() == "X" {
			break
		}

		input, _ = strconv.ParseInt(scanner.Text(), 10, 64)
		myslice = append(myslice, input)

		sort.Slice(myslice, func(i, j int) bool {
			return myslice[i] < myslice[j]
		})

		fmt.Println(myslice)
	}
}
