package comparable

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
