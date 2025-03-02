package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in the GOLANG.")
	fmt.Println("WE")

	var skills [4]string
	skills[0] = "FRONTEND"
	skills[2] = "BACKEND"
	skills[3] = "DEVOPS"
	fmt.Println("The list of skills is: ", skills)
	fmt.Println("The lenght of the skills array is: ", len(skills)) //IT ALSWAYS GOING TO SHOW THE LEMTH SAME AS U INIT AT THE START WHILE INIT THE ARRAY NO MATTER WHAT HOW MANY ITEMS U HAVE PLACED WITHIN THE ARRAY
}
