package main

func main() {
	// assign a slice of integers to a variable
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// iterate over the slice of integers
	for i, number := range numbers {
		// check if the number is even or odd

		if number%2 == 0 { // if the number is even
			println(i, "is even")
		} else { // if the number is odd
			println(i, "is odd")
		}
	}
}
