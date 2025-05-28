package main

import "fmt"

func main() {

	/*
		// declaration
		var productRanks map[string]int

		// initialization
		productRanks = make(map[string]int)

		// adding items
		productRanks["pen"] = 5
		productRanks["pencil"] = 1
		productRanks["marker"] = 3
	*/

	/*
		var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	*/

	var productRanks map[string]int = map[string]int{
		"pen":    5,
		"pencil": 1,
		"marker": 3,
	}

	fmt.Println("productRanks :", productRanks)

	// Accessing values by key
	fmt.Printf("Rank of pencil : %d\n", productRanks["pencil"])

	fmt.Printf("len(productRanks) : %d\n", len(productRanks))

	fmt.Println("Iteration using range")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existence of a key
	// keyToCheck := "stapler"
	keyToCheck := "marker"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q = %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Key [%q] does not exist!\n", keyToCheck)
	}

	// removing an item
	// keyToRemove := "marker"
	keyToRemove := "stapler"
	delete(productRanks, keyToRemove)
	fmt.Printf("After removing key [%q], productRanks = %v\n", keyToRemove, productRanks)

}
