package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func readLine(prompt string) (string, error) {
	if prompt != "" {
		fmt.Println(prompt)
	}
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", errors.New("failed to read input")
	}
	return scanner.Text(), nil
}

func main() {
	userMap := make(map[string]string)
	in, err := readLine("Please enter a name:")
	if err != nil {
		panic(err)
	}
	userMap["name"] = in
	in, err = readLine("Please enter an address:")
	if err != nil {
		panic(err)
	}
	userMap["address"] = in
	uj, err := json.Marshal(userMap)
	if err != nil {
		fmt.Printf("Error occured while marshalling to JSON:\n%v", err.Error())
	}
	fmt.Printf("JSON object:\n%s\n", uj)
}
