package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	createdAnimals := make(map[string]Animal)
	var input, command, animalName, animalType, information string
	var splittedInput []string

	fmt.Println("Create new animals and query information about them!")
	fmt.Println("Create a new animal using the following command: newanimal {animal_name} {animal_type}")
	fmt.Println("Available animal types: cow, bird and snake")
	fmt.Println("Query information of a created animal by using the following command: query {animal_name} {information_name}")
	fmt.Println("Available animal names: the ones you created earlier, Available information names: eat, move and speak")
	fmt.Println("You can exit by typing 'exit'")
	fmt.Print("> ")

	for scanner.Scan() {
		input = scanner.Text()
		if input == "exit" {
			break
		}
		splittedInput = strings.Split(input, " ")
		if len(splittedInput) != 3 {
			fmt.Println("Enter 3 values, seperated by a space!")
			fmt.Print("> ")
			continue
		}

		command = cleanInput(splittedInput[0])
		animalName = cleanInput(splittedInput[1])

		switch command {
		case "newanimal":
			animalType = cleanInput(splittedInput[2])
			animal, err := createAnimal(animalType)
			if err != nil {
				fmt.Println(err)
			}
			createdAnimals[animalName] = animal
			fmt.Println("Created animal!")
		case "query":
			information = cleanInput(splittedInput[2])
			if animal, ok := createdAnimals[animalName]; ok {
				executeQuery(animal, information)
			} else {
				fmt.Println("Requested animal doesn't exist yet.")
				fmt.Print("> ")
				continue
			}
		default:
			fmt.Println("Unknown command, use newanimal or query")
		}

		fmt.Print("> ")
	}
}

// createAnimal creates and returns the animal corresponding to the provided animal type
func createAnimal(animalType string) (Animal, error) {
	var animal Animal

	switch animalType {
	case "cow":
		animal = Cow{"grass", "walk", "moo"}
	case "bird":
		animal = Bird{"worms", "fly", "chirp"}
	case "snake":
		animal = Snake{"mice", "slither", "hsss"}
	default:
		return nil, errors.New("Unavailable animal type")
	}

	return animal, nil
}

func executeQuery(animal Animal, information string) {
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

// cleanInput returns the input back without trailing whitespaces in all lower case
func cleanInput(input string) (cleanedInput string) {
	cleanedInput = strings.ToLower(input)
	cleanedInput = strings.Trim(cleanedInput, " ")
	return
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

// ===== Cow =====
type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

// ===== Snake =====
type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

// ===== Bird =====
type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}
