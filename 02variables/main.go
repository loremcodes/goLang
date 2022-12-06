package main

import "fmt"

const UserId int = 1967971242007 //  Public variable because 'U' is captial

func main() {
	// create variable -> string
	var username string = "prasad"
	// Print the variable
	fmt.Println(username)
	// find the type of variable
	fmt.Printf("The username variable is of type: %T \n", username)

	// variable -> bool
	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("The isTrue variable is of type: %T \n", isTrue)

	// default values for string
	var emptyInt int
	fmt.Println(emptyInt)

	// Implicit type
	var variableWithoutType = 34
	fmt.Println(variableWithoutType)

	// no var style -> only avaibale in main function
	variableWithoutVarKey := 4307
	fmt.Println(variableWithoutVarKey)

	// Calling public variable
	fmt.Println(UserId)
	fmt.Printf("UserId is type of: %T \n ", UserId)
}
