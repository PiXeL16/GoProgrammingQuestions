package main

import (
	"fmt"
)

func main() {

	sortedArrayWithDuplicates := []int{1, 2, 2, 3, 4, 5, 6, 6, 7}

	fmt.Println(removeDuplicates(sortedArrayWithDuplicates))

}

func removeDuplicates(nums []int) []int {

	var uniqueSortedArray []int
	y := 0

	if len(uniqueSortedArray) == 0 {
		uniqueSortedArray = append(uniqueSortedArray, nums[0])
	}

	for x := 1; x < len(nums); x++ {

		if uniqueSortedArray[y] != nums[x] {
			uniqueSortedArray = append(uniqueSortedArray, nums[x])
			y++
		}
	}

	return uniqueSortedArray

}
