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
	books := []comparable2.Book{
		{"Go in Action", 4.2},
		{"The Go Programming Language", 4.8},
		{"Learning Go", 4.5},
	}
	c := comparable2.TopBooks(books, 2)
	fmt.Printf("Top Books: %v\n", c)
}
