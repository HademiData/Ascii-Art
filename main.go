package main

import (
	"Ascii-Art/helperFunctions"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	inputText := os.Args[1]
	//fmt.Println(inputText)

	inputText = strings.ReplaceAll(inputText, "\\n", "\n")
	//fmt.Println(inputText)

	if inputText == "" {
		return
	} else {
		buffer, err := os.ReadFile("standard.txt")

		if err != nil {
			log.Fatal("Error Reading the banner file..")
		}

		banner := strings.Split(string(buffer), "\n")
		//fmt.Println(banner)

		lines := strings.Split(inputText, "\n")
		//fmt.Println(len(lines))

		for _, line := range lines {
			if line == "" {
				fmt.Println()
				continue
			}
			helperFunctions.PrintAsciiArt(line, banner)
		}
	}
}
