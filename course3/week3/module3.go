package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	// if you want you can use these numbers: -8 8 3 2 17 9 18 27 19 0 10 12 21 888 12 313
	fmt.Println("Type a list of numbers you want sorted here, all on one line :")

	numbers := getUserInput()
	if !isValid(numbers) {
		fmt.Println("provided amount of numbers should be in the table of 4")
		return
	}

	var q []int
	var wg sync.WaitGroup
	wg.Add(4)
	steps := len(numbers) / 4
	for i := 0; i < len(numbers); i += steps {
		q = numbers[i : i+steps]
		go Sort(q, &wg)
	}
	wg.Wait()
	sort.Ints(numbers)
	fmt.Println(numbers)
}

func getUserInput() (unsorted []int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	splittedInput := strings.Split(scanner.Text(), " ")
	var numbers []int

	for _, v := range splittedInput {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		numbers = append(numbers, num)
	}

	return numbers
}

func Sort(quarter []int, wg *sync.WaitGroup) {
	fmt.Println("unsorted quarter: ", quarter)
	for j := 0; j < len(quarter); j++ {
		for i := 0; i < len(quarter)-1; i++ {
			if quarter[i] > quarter[i+1] {
				swap(quarter, i)
			}
		}
	}
	fmt.Println("sorted quarter: ", quarter)
	fmt.Println()
	wg.Done()
}

func swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func isValid(sli []int) bool {
	return len(sli)%4 == 0
}
