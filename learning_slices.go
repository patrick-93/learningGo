package main

import "fmt"

func Slice_test() {
	// Create a new slice using make() to build an array under the hood
	myslice := make([]int, 0)
	// Print length/capacity of the initial array
	print_slice(myslice)

	// Create a new array of ints and loop over it using a range for loop
	fmt.Println("Starting looping over the array")
	myarr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

	for _, v := range myarr {
		// fmt.Printf("Cur index: %d, Val: %d\n", i, v)

		// Append to the slice
		myslice = append(myslice, v)
	}
	// Show length/capacity of the slice after adding to it
	print_slice(myslice)
}

func print_slice(s []int) {
	fmt.Printf("Current slice size Length: %d, Capacity: %d\n", len(s), cap(s))
}