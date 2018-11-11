package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	type name struct {
		fname string
		lname string
	}

	var filename string

	fmt.Println("Please enter the filename: ")
	_, err := fmt.Scan(&filename)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	names := make([]*name, 0, 100)
	var line string
	var splitLine []string

	for scanner.Scan() {
		line = scanner.Text()
		splitLine = strings.Split(line, " ")
		names = append(names, &name{fname: splitLine[0], lname: splitLine[1]})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, name := range names {
		fmt.Printf("Firstname: %s, Lastname: %s\n", name.fname, name.lname)
	}
}
