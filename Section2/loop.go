package main

import "fmt"

func main() {
	// for - only way to loop
	// It like C programming loop
	for i := 1; i <= 10; i++ {
		//fmt.Println(i)
	}

	// while loop
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("----------------------")

	// if else
	tmp := 22
	if tmp > 30 {
		fmt.Println("Greater than 30")
	} else {
		fmt.Println("Number is less than 30")
	}

	fmt.Println("------------------")

	// switch statement

	day := "Monday"
	fmt.Println("Today is ", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend no work")
	case "Monday", "Tuesday":
		fmt.Println("Work days")
	default:
		fmt.Println("Midweek")
	}

}
