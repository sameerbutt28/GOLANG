package main

import "fmt"

func main() {
	fmt.Println("Started learning struct in GOLANG.")
	//NO INHERTIANCE INN GOLANG, NO PARENT CONCEPT
	sameer := User{"Sameer", "sameer@go.dev", 25, true}
	fmt.Println("Data about the user is:", sameer)
	fmt.Printf("Data about the user is: %+v \n", sameer)
	fmt.Printf("Name of user is %v and the email of user is %v \n", sameer.Name, sameer.Email)
}

type User struct {
	Name    string
	Email   string
	Age     int
	Verfied bool
}
