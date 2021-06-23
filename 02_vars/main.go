package main

import (
	"fmt"
)

func main() {
	stringVar := "string example"
	intVar := 99
	floatVar := 7.992
	boolVar := false
	sliceVar := []string{"apple", "bear"}
	charSliceVar1 := []byte{'a', 'p', 'p', 'l', 'e'}
	charSliceVar2 := []rune{'b', 'e', 'a', 'r'}
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v", stringVar, intVar, boolVar, floatVar, sliceVar, charSliceVar1, charSliceVar2)
}
