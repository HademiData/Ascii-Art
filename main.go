package main

import (
	"fmt"
	"os"
	 "strings"
)

func main() {
	data, err := os.ReadFile("standard.txt")

	if err != nil {
		panic(err)
	}

	text := string(data)

	visible := strings.ReplaceAll(text, " ", ".")

	fmt.Println(visible)
}
