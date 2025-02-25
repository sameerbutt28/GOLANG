package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Lecture about the slices concept.")
	var alpha = []string{"a", "b", "c"} //THIS AUTOMATICALLY EXPANDS THE MEMORY FOR US, ADD AS MANY DATA AS U CAN WITH NO LIMIT AND THE MEMORY WILL INCREASE IN THAT WAY.
	fmt.Printf("The data is this  is of type: %T \n", alpha)
	alpha = append(alpha, "d", "e")
	fmt.Println(alpha)
	alpha = append(alpha[1:])
	fmt.Println(alpha)

	highScores := make([]int, 4) //default allocaion of the memory
	highScores[0] = 28
	highScores[1] = 11
	highScores[2] = 3
	highScores[3] = 10
	// highScores[4] = 14
	fmt.Println(highScores)
	highScores = append(highScores, 65, 12)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //TP CHECK IF THE INTERGER DATA IS SORTED OR NOT

	//REMOVE A VALUE FROM THE SLICE ON THE BASIS OF INDEX
	courses := []string{"react", "go", "python", "swift", "ruby", "mongo"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
