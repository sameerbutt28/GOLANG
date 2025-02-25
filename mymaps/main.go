package main

import (
	"fmt"
)

func main() {
	fmt.Println("Mapps in GOLANG.")      //HASH TABLES, KEY VALUE PAIRS
	languages := make(map[string]string) //INNER VALUE OR INNER STRING  IS KEY AND THE STRING OUTSIDE OF THE SQUARE BRACKETS IS VALUE
	languages["JS"] = "JAVASCRIPT"
	languages["RB"] = "RUBY"
	languages["SF"] = "SWIFT"
	fmt.Println("Mapping is done in a way: ", languages)
	fmt.Println("JS short form is of ", languages["JS"])
	// delete(languages, "RB")
	fmt.Println("Mapping is done in a way: ", languages)

	//LOOPS IN GOLANG
	for key, value := range languages { //YOU CAN ALSO PLACE THE UNDERSCORE SIGN INSTEAD OF THE KEY IN THE FOR LOOP THROUGH WHIHC U CAN EASILY IGNORE  THE KEY WITH ANY ERRORS.
		fmt.Printf("For key %v, Value is %v\n", key, value)
	}

}
