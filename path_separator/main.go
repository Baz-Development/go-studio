package main

import ("fmt"; "path")

func main() {
	fmt.Println("path separator")
	// short declaration
	dir, file := path.Split("src/assets/image.png")
	fmt.Println("path separator dir: ", dir)
	fmt.Println("path separator file: ", file)
}