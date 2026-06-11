package comparable

import (
	"cmp"
	"slices"
)

func FindDuplicates[T comparable](xs []T) []T {
	count := make(map[T]int)
	var result []T
	for _, x := range xs {
		count[x]++
		if count[x] == 2 {
			result = append(result, x)
		}
	}
	return result
}
func CountOccurrences[T comparable](xs []T) map[T]int {
	var count map[T]int
	count = make(map[T]int)
	for _, x := range xs {
		count[x]++
	}
	return count
}

type Book struct {
	Title  string
	Rating float64
}

func TopBooks(books []Book, n int) []Book {
	if n <= 0 {
		return nil
	}
	clone := slices.Clone(books)
	slices.SortFunc(clone, func(a, b Book) int {
		return cmp.Compare(b.Rating, a.Rating)
	})
	if n > len(clone) {
		n = len(clone)
	}
	return clone[:n]
}
