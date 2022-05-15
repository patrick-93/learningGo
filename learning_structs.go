package main

import "fmt"

func Struct_test() {
	// Create a new var of type "Custom" from structs.go
	myVar := Custom{"custom", 17}
	myVar.name = "bob" // Edit the name field
	fmt.Printf("Name: %s, Val: %d\n", myVar.name, myVar.val)
}