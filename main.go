package main

import (
	"errors"
	"fmt"
)

func main() {
	//
}

func getUserInput(prompt string) (string, error) {
	var value string
	var valuePointer *string = &value

	fmt.Print(prompt)

	fmt.Scan(valuePointer)

	if *valuePointer == "" {
		return "", errors.New("Invalid input")
	}

	return value, nil
}