package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input your rating between 1 to 5")
	input, _ := reader.ReadString('\n')
	// fmt.Print(input)

	// adding 1 to the initial input
	addToInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(addToInput + 1, " is your final rating Thanks..!")
	}
}
