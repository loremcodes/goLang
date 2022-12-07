package main

import "fmt"

func main() {
	/*
		    1. In go array has a fixed length
		    2. inferred -> complier will decide the length of array

			syntax with var keyword
		    var array_name = [array_size]data_type{values} -> array length is defined

		    or

		    var array_name = [...]data_type{values} -> array length is inferred

		    syntax with :=
		    array_name := [array_size]data_type{values} -> array length is defined

		    or

		    array_name := [...]data_types{values} -> array length is inferred
	*/

    // initalize array
    unInitialized := [5] int {}
    fmt.Println(unInitialized, " -> uninitialized array")

    // Partially initialized array
    partiallyInitialized := [5] int {1, 2}
    fmt.Println(partiallyInitialized, " -> partially initialized array")

    // fully initialized array
    fullyInitialized := [5] int {1, 2, 3, 4, 5}
    fmt.Println(fullyInitialized, " -> fully initialized array")

	// declare an array with var keyword
	var arr = [...]int{1, 2, 3}
	fmt.Println(arr)

	// declare an array without var keyword
	arrAnother := [...]int{4, 5, 6}
	fmt.Print(arrAnother, "\n")

	// access element of of arr
	fmt.Print(arr[0], "\n")

	// chnaging element of arr
	arr[0] = 0
    fmt.Print(arr[0] , " element value has changed \n")

    // find length of an array
    fmt.Print(len(arr))
}
