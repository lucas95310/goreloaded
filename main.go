package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	text := string(content)

	text = ProcessText(text)

	err = os.WriteFile(os.Args[2], []byte(text), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
