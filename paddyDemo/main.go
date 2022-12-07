//TODO: Take user input and display it.

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your fname")
    input, _ := reader.ReadString('\n')
    fmt.Println(input)
}