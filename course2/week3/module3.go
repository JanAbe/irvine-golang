package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input, animalName, information string
	var splittedInput []string

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"} // Do birds peep? why not chirp? :)
	snake := Animal{"mice", "slither", "hsss"}

	fmt.Println("Request information about an animal in the following way: {animal_name} {information_type}")
	fmt.Println("Available animal names: cow, bird and snake")
	fmt.Println("Available information types: eat, move and speak")
	fmt.Println("You can exit by type 'exit'")
	fmt.Print("> ")

	for scanner.Scan() {
		input = scanner.Text()
		if input == "exit" {
			break
		}

		splittedInput = strings.Split(input, " ")
		if len(splittedInput) != 2 {
			fmt.Println("Enter 2 values, seperated by a space!")
			continue
		}

		animalName = strings.ToLower(strings.Trim(splittedInput[0], " "))
		information = strings.ToLower(strings.Trim(splittedInput[1], " "))

		switch animalName {
		case "cow":
			execMethod(cow, information)
		case "bird":
			execMethod(bird, information)
		case "snake":
			execMethod(snake, information)
		default:
			fmt.Println("You requested an unavailable animal")
		}
		fmt.Print("> ")
	}
}

func execMethod(animal Animal, information string) {
	switch information {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("You requested unavailable information")
	}
}

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}
