package main

import (
	"fmt"
)

func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Decalre and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[:])   // make a copy of fruitSlice
	fmt.Println(fruitSlice[1:3]) // make a copy of fruitSlice from index 1 to 3(excluded)

	fruitSlice = append(fruitSlice, "Pear")
	fmt.Println(fruitSlice)
}
