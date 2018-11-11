package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dict := make(map[string]string)
	var name, address string

	fmt.Println("Enter a name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Enter an address: ")
	_, err = fmt.Scan(&address)
	if err != nil {
		fmt.Println(err)
	}

	dict["name"] = name
	dict["address"] = address

	marshalledDict, err := json.Marshal(dict)
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(marshalledDict)
	fmt.Println()
}
