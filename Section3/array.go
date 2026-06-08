package main

import "fmt"

func main() {
	// Array
	var numbers [2]int
	fmt.Printf("%+v\n", numbers)

	numbers[0] = 1
	numbers[1] = 2

	fmt.Printf("%+v\n", numbers)

	primes := [4]int{2, 3, 5}
	fmt.Printf("%v+\n", primes)
}
