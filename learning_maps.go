package main

import "fmt"

func Map_test() {
	// make an empty map of <string, int>
	mymap := make(map[string]int)

	// Add some values
	mymap["sometext"] = 17
	mymap["another"] = 27
	mymap["another"] = 4 // overwite a cur val

	// print the whole thing
	fmt.Println(mymap)

	// Loop over and print keys/vals
	print_kv(mymap)

	// Delete a value
	fmt.Println("Deleting value from mymap")
	delete(mymap, "another")
	print_kv(mymap)

	// Check if a key exists in the map
	elem, ok := mymap["somtext"] // non existing element
	if ok == true {
		fmt.Printf("Element exists, %d", elem)
	} else {
		fmt.Println("element does not exist")
	}

}

func print_kv(m map[string] int) {
	fmt.Println("Looping over map...")
	for k, v := range m {
		fmt.Printf("\tKey: %s, Val: %d\n", k, v)
	}
}
