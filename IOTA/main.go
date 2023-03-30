package main

func main() { 
	const (
		monday = 0
		tuesday = 1
		wednesday = 2
		thuesday = 3
		friday = 4
		saturday = 5
		sunday = 6
	)

	const (
		monday_IOTA = iota + 1
		tuesday_IOTA
		wednesday_IOTA
		thuesday_IOTA
		friday_IOTA
		saturday_IOTA
		sunday_IOTA
	)

	println("without IOTA: ",monday, tuesday, wednesday, thuesday, friday, saturday, sunday)
	println("with IOTA: ",monday_IOTA, tuesday_IOTA, wednesday_IOTA, thuesday_IOTA, friday_IOTA, saturday_IOTA, sunday_IOTA)

}