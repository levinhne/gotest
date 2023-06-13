package main

import (
	"fmt"
	"reflect"
)

// Main function
func main() {

	map_1 := map[int]string{
		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}
	map_2 := map[int]string{
		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}
	// DeepEqual is used to check
	// two interfaces are equal or not
	res1 := reflect.DeepEqual(map_1, map_2)
	fmt.Println("Is Map 1 is equal to Map 2: ", res1)
}
