package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type text struct {
		Name        string
		words		int
	}

	type book struct {
		text
		Name        string
		Published   bool
		Author      string
		releaseDate string
	}

	pacman := book{
		text: text{Name: "pac-man", words: 1012831683761236712},
		Published:   true,
		Author:      "Shigeo Funaki Shigeichi Ishimura",
		releaseDate: "July 1980",
	}


	our, err := json.MarshalIndent(pacman, "", "\t")
	if err != nil {
		return 
	}
	fmt.Printf("Pacman: %s\n", string(our))
}