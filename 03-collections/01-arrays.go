package main

import "fmt"

func main() {
	// declaration
	// var nos [5]int

	// initialization
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// using :=
	nos := [5]int{3, 1, 4, 2, 5}
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

	newNos := nos // copy is created coz Arrays are VALUES
	newNos[0] = 9999
	fmt.Printf("newNos[0] = %d, nos[0] = %d\n", newNos[0], nos[0])

	x := [5]int{3, 1, 4, 2, 5}
	y := [5]int{3, 1, 4, 2, 5}
	fmt.Println("x == y ? :", x == y)

	fmt.Println("Before sorting x = ", x)
	sortNos(&x)
	fmt.Println("After sorting x = ", x)

}

func sortNos(list *[5]int) {
	fmt.Println((*list)[0], list[0])

	fmt.Println("[sortNos] list = ", list)
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	fmt.Println("[sortNos] list = ", list)
}
