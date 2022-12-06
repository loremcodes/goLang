package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	// Taking input form user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")

	// comma ok or err ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Print("You entered ", input)
	fmt.Printf("input is of type: %T \n", input)
}