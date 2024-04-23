package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	// Map each element 'nums' in slice 'nums' to its square and store the result in a new slice
	// Sort the new slice in non-decreasing order
	// Return the sorted slice
	squared := make([]int, len(nums))
	for i, num := range nums {
		squared[i] = num * num
	}
	sort.Ints(squared)
	return squared
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	sortedSquaresArray := sortedSquares(nums)
	fmt.Println(sortedSquaresArray) // Output: [0 1 9 16 100]
}
