package main

import "fmt"

func main() {
	speed := 10 // speed is int
	time := 2.5 // time is float64
	aceleration := speed * int(time)
	fmt.Println("A aceleração é: ", aceleration)
}