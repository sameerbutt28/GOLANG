package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the pizza rating: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of this rating is %T \n ", input)

}
