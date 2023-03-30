package main

func main() {
	for i := 0; i < 10; i++ {
		println("ola mundo ", i)
	}

	for i := 0; i < 10; i++ {
		println("ola mundo ", i)
		break
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			println("par")
			continue
		}
		println("impar")
	}
}