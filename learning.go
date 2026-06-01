package main

import "fmt"

// constants
const HOST string = "localhost"

// enums
const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	//Simple descriptions of string and number
	fmt.Printf("Hello World \n")
	fmt.Println(1 + 1)
	fmt.Println(true)
	fmt.Println("What is it")

	//Declaring variables
	//Normally we declare from left to right but in go it is opposite

	var greeting string //zero-value is an empty string
	greeting = "Hello World"
	fmt.Println(greeting)

	var count int
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstname, lastname string
	firstname = "john"
	lastname = "doe"
	fmt.Println(firstname, lastname)

	//Shorter version now
	//if you want to declare and initialize a variable this is the prefered option :=
	email := "test@com"
	fmt.Println(email)

	//Constatnts
	fmt.Println(HOST)

	//Enums
	fmt.Println(Sunday)
}
