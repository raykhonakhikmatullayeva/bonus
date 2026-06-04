package main

import (
	comparable2 "bonus/internal/comparable"
	"fmt"
)

func main() {
	a := comparable2.FindDuplicates([]int{2, 3, 3, 3, 2, 4})
	b := comparable2.CountOccurrences([]string{"еда", "еда", "такси"})

	fmt.Println("Duplicates: ", a)
	fmt.Println("Counts: ", b)

}
