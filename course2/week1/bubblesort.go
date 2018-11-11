package main

import (
	"fmt"
	"strconv"
)

func main() {
	unsorted := GetUserInput(10)
	BubbleSort(unsorted)
	fmt.Println(unsorted)
}

// GetUserInput asks the user n times to add a number and returns a slice containing these numbers
func GetUserInput(n int) (unsorted []int) {
	var input string

	fmt.Printf("Enter %d numbers and press the return key after each number\n Or type 'exit' to quit sooner: \n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		if input == "exit" {
			break
		}

		convInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("error: ", err)
		}

		unsorted = append(unsorted, convInput)
	}

	return unsorted
}

// BubbleSort sorts an unsorted slice using the bubble sort algorithm
func BubbleSort(unsorted []int) {
	for j := 0; j < len(unsorted); j++ {
		for i := 0; i < len(unsorted)-j-1; i++ {
			if unsorted[i] > unsorted[i+1] {
				Swap(unsorted, i)
			}
		}
	}
}

// Swap swaps two elements from the provided slice at the provided index
func Swap(sli []int, indx int) {
	temp := sli[indx]
	sli[indx] = sli[indx+1]
	sli[indx+1] = temp
}
