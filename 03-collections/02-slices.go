package main

import "fmt"

func main() {
	// declaration
	// var nos []int

	// initalization
	var nos []int = []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// iterating using indexer
	fmt.Println("iterating using indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	newNos := nos // a copy of the REFERENCE is created
	newNos[0] = 9999
	fmt.Printf("newNos[0] = %d, nos[0] = %d\n", newNos[0], nos[0])
	nos[0] = 3

	fmt.Printf("Before sorting, nos = %v\n", nos)
	sort(nos)
	fmt.Printf("After sorting, nos = %v\n", nos)

	// dynanically add items
	nos = append(nos, 7)
	fmt.Println(nos)

	nos = append(nos, 30, 20, 50, 10)
	fmt.Println(nos)

	hundreds := []int{400, 100, 200}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	// slicing
	fmt.Println("nos[3:7] =", nos[3:7])
	fmt.Println("nos[: 7] =", nos[:7])
	fmt.Println("nos[3: ] =", nos[3:])

}

func sort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
