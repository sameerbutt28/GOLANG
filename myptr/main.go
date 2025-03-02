package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers..")
	var myptr *int
	fmt.Println("Value of the pointer is:  ", myptr)

	myNumber := 28
	var ptr = &myNumber
	fmt.Println("Memory location of the pointer is: ", ptr)
	fmt.Println("Value of the pointer is: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", *ptr)

}
