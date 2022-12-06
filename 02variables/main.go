package main

import "fmt"

func main() {
	// create variable -> string
	var username string = "prasad"
	// Print the variable
	fmt.Println(username)
	// find the type of variable
	fmt.Printf("The username variable is of type: %T \n", username)

	//    variable -> bool
	var isTrue bool = true
    fmt.Println(isTrue)
    fmt.Printf("The isTrue variable is of type: %T", isTrue)
}
