package main

import (
	"bufio"
	"fmt"
	"os"
)

func reader() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		if len(text) != 0 {
			return text
		} else {
			// exit if user entered an empty string
			break
		}

	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
		return ""
	}
	return ""
}

func main() {
	var board [3][3]string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = reader()
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("[", board[i][j], "]")
		}
		fmt.Println()
	}
}