package main

import "fmt"

func main() {
	type Game struct {
		Name        string
		Published   bool
		Author      string
		releaseDate string
	}

	pacman := Game{
		Name:        "pac-man",
		Published:   true,
		Author:      "Shigeo Funaki Shigeichi Ishimura",
		releaseDate: "July 1980",
	}

	fmt.Printf("Pacman: %+v\n", pacman)
}