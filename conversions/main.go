package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to out PIZZA APP")
	fmt.Println("Please rate our pizza between 1-5.")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Add 1 to your rating ", numRating+1)
	}

}
