package main

import "fmt"

func main() {
	var books [4]string

	books[0] = "Design patterns"
	books[1] = "Clean code"
	books[2] = "The Witcher"
	books[3] = "Stay golden"

	fmt.Printf("Books : %t\n", books)
	fmt.Println("Books : \n", books)
	fmt.Printf("Books : %q\n", books)
	fmt.Printf("Books : %#v\n", books)

	fmt.Println("-------------")

	// Literal array
	var books_literal = [4]string {
		"Design patterns",
		"Clean code",
		"The Witcher",
		"Stay golden",
	}

	fmt.Printf("books_literal : %t\n", books_literal)
	fmt.Println("books_literal : \n", books_literal)
	fmt.Printf("books_literal : %q\n", books_literal)
	fmt.Printf("books_literal : %#v\n", books_literal)

	fmt.Println("-------------")

	// Literal array
	var books_literal_dynamic = [...]string {
		"Design patterns",
		"Clean code",
		"The Witcher",
		"Stay golden",
	}

	fmt.Printf("books_literal_dynamic : %t\n", books_literal_dynamic)
	fmt.Println("books_literal_dynamic : \n", books_literal_dynamic)
	fmt.Printf("books_literal_dynamic : %q\n", books_literal_dynamic)
	fmt.Printf("books_literal_dynamic : %#v\n", books_literal_dynamic)
}