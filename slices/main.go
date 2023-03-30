package main

import "fmt"

func main() {
	var games []string
	games = append(games, "The witcher 3")

	fmt.Printf("slice: %T\n", games)
	fmt.Printf("slice: %d\n", len(games))

	// limit capacity

	games = append(games, "The witcher 2", "The witcher 1", "The witcher 4")
	fmt.Printf("slice: %T\n", games)
	fmt.Printf("slice: %d\n", len(games))

	rpg := games[:2:2]
	fmt.Printf("slice limit capacity: %T\n", rpg)
	fmt.Printf("slice limit capacity: %d\n", len(rpg))

	// Allocate a initial size to array

	slice_preallocated := make([]int, 3)
	fmt.Printf("slice_preallocated: %T\n", slice_preallocated)
	fmt.Printf("slice_preallocated: %d\n", len(slice_preallocated))

	// copy
	var copy_slice []string
	copy(copy_slice, games)
	fmt.Printf("slice copy: %T\n", copy_slice)
	fmt.Printf("slice copy: %d\n", len(copy_slice))

}