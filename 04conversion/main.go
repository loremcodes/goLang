package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//TODO: Take user input from 1 to 5 and 1 to final input

	// Get input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input your rating between 1 to 5")
	input, _ := reader.ReadString('\n')

	// adding 1 to the initial input -> convert string to int (strconv.ParseInt)
	addToInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(addToInput+1, " is your final rating Thanks..!")
	}
}

/* 
New things learnt: 
1. strconv
	-> ParseFloat
2. strings
	-> strings.TrimSpace
3. if else in go
4. nil
5. err ok syntax
*/
